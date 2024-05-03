package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/kenboo0426/franky_assessment/migration/migrations"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
	"github.com/uptrace/bun/migrate"
	"github.com/urfave/cli/v2"
)

type CreateProductRequest struct {
	Name   string   `json:"name"`
	Brand  string   `json:"brand"`
	Images []string `json:"images"`
}

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		panic(err)
	}
	sqlDB, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	))
	if err != nil {
		panic(err)
	}

	db := bun.NewDB(sqlDB, mysqldialect.New())
	defer db.Close()

	app := &cli.App{
		Name: "bun",

		Commands: []*cli.Command{
			newDBCommand(migrate.NewMigrator(db, migrations.Migrations)),
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func newDBCommand(migrator *migrate.Migrator) *cli.Command {
	return &cli.Command{
		Name:  "db",
		Usage: "database migrations",
		Subcommands: []*cli.Command{
			{
				Name:  "init",
				Usage: "create migration tables",
				Action: func(c *cli.Context) error {
					return migrator.Init(c.Context)
				},
			},
			{
				Name:  "migrate",
				Usage: "migrate database",
				Action: func(c *cli.Context) error {
					if err := migrator.Lock(c.Context); err != nil {
						return err
					}
					defer migrator.Unlock(c.Context) //nolint:errcheck

					group, err := migrator.Migrate(c.Context)
					if err != nil {
						return err
					}
					if group.IsZero() {
						fmt.Printf("there are no new migrations to run (database is up to date)\n")
						return nil
					}
					fmt.Printf("migrated to %s\n", group)
					return nil
				},
			},
			{
				Name:  "rollback",
				Usage: "rollback the last migration group",
				Action: func(c *cli.Context) error {
					if err := migrator.Lock(c.Context); err != nil {
						return err
					}
					defer migrator.Unlock(c.Context) //nolint:errcheck

					group, err := migrator.Rollback(c.Context)
					if err != nil {
						return err
					}
					if group.IsZero() {
						fmt.Printf("there are no groups to roll back\n")
						return nil
					}
					fmt.Printf("rolled back %s\n", group)
					return nil
				},
			},
			{
				Name:  "lock",
				Usage: "lock migrations",
				Action: func(c *cli.Context) error {
					return migrator.Lock(c.Context)
				},
			},
			{
				Name:  "unlock",
				Usage: "unlock migrations",
				Action: func(c *cli.Context) error {
					return migrator.Unlock(c.Context)
				},
			},
			{
				Name:  "create_go",
				Usage: "create Go migration",
				Action: func(c *cli.Context) error {
					name := strings.Join(c.Args().Slice(), "_")
					mf, err := migrator.CreateGoMigration(c.Context, name)
					if err != nil {
						return err
					}
					fmt.Printf("created migration %s (%s)\n", mf.Name, mf.Path)
					return nil
				},
			},
			{
				Name:  "create_sql",
				Usage: "create up and down SQL migrations",
				Action: func(c *cli.Context) error {
					name := strings.Join(c.Args().Slice(), "_")
					files, err := migrator.CreateSQLMigrations(c.Context, name)
					if err != nil {
						return err
					}

					for _, mf := range files {
						fmt.Printf("created migration %s (%s)\n", mf.Name, mf.Path)
					}

					return nil
				},
			},
			{
				Name:  "create_tx_sql",
				Usage: "create up and down transactional SQL migrations",
				Action: func(c *cli.Context) error {
					name := strings.Join(c.Args().Slice(), "_")
					files, err := migrator.CreateTxSQLMigrations(c.Context, name)
					if err != nil {
						return err
					}

					for _, mf := range files {
						fmt.Printf("created transaction migration %s (%s)\n", mf.Name, mf.Path)
					}

					return nil
				},
			},
			{
				Name:  "status",
				Usage: "print migrations status",
				Action: func(c *cli.Context) error {
					ms, err := migrator.MigrationsWithStatus(c.Context)
					if err != nil {
						return err
					}
					fmt.Printf("migrations: %s\n", ms)
					fmt.Printf("unapplied migrations: %s\n", ms.Unapplied())
					fmt.Printf("last migration group: %s\n", ms.LastGroup())
					return nil
				},
			},
			{
				Name:  "mark_applied",
				Usage: "mark migrations as applied without actually running them",
				Action: func(c *cli.Context) error {
					group, err := migrator.Migrate(c.Context, migrate.WithNopMigration())
					if err != nil {
						return err
					}
					if group.IsZero() {
						fmt.Printf("there are no new migrations to mark as applied\n")
						return nil
					}
					fmt.Printf("marked as applied %s\n", group)
					return nil
				},
			},
			{
				Name:  "init_data",
				Usage: "insert product data",
				Action: func(c *cli.Context) error {
					requestData := []CreateProductRequest{
						{
							Name:  "KENZO 'TIGER CREST' POLO SHIRT",
							Brand: "KENZO",
							Images: []string{
								"https://stok.store/cdn/shop/files/20220304040105603_E52---kenzo---FA65PO0014PU01B_4_M1.jpg", "https://stok.store/cdn/shop/files/20220304040124567_E52---kenzo---FA65PO0014PU01B_5_M1.jpg",
								"https://stok.store/cdn/shop/files/20220304040128900_E52---kenzo---FA65PO0014PU01B_7_M1.jpg",
							},
						},
						{
							Name:  "A.P.C. 'AURELIA' DENIM DRESS",
							Brand: "A.P.C.",
							Images: []string{
								"https://stok.store/cdn/shop/files/5T953DF0503F33F0002_01_M_2024-02-22T08-12-47.316Z.jpg",
								"https://stok.store/cdn/shop/files/20211218140947600_E52---apc---COETKF05822IAL_2_M1.jpg",
							},
						},
						{
							Name:  "MIU MIU VINTAGE LEATHER ANKLE BOOTS",
							Brand: "MIU MIU",
							Images: []string{
								"https://stok.store/cdn/shop/files/5T953DF0503F33F0002_01_M_2024-02-22T08-12-47.316Z.jpg",
								"https://stok.store/cdn/shop/files/5T953DF0503F33F0002_04_M_2024-02-22T08-12-47.566Z.jpg",
							},
						},
						{
							Name:  "Alexander McQUEEN 'SLIM TREAD' ANKLE BOOTS",
							Brand: "Alexander McQUEEN",
							Images: []string{
								"https://stok.store/cdn/shop/files/20220125140133136_E52---alexander_20mcqueen---690812W4SQ11053_1_M1.jpg",
							},
						},
						{
							Name:  "STIVALETTO",
							Brand: "Alexander McQUEEN",
							Images: []string{
								"https://stok.store/cdn/shop/files/757487WIDU11000_2023-07-07T07-27-52.221Z.jpg",
								"https://stok.store/cdn/shop/files/757487WIDU11000_5_P_2023-07-07T07-27-52.533Z.jpg",
								"https://stok.store/cdn/shop/files/757487WIDU11000_3_P_2023-07-07T07-27-52.392Z.jpg",
							},
						},
						{
							Name:  "Alexander McQUEEN Black leather zipped card holder with logo",
							Brand: "Alexander McQUEEN",
							Images: []string{
								"https://stok.store/cdn/shop/files/6831171AAMJ_O_ALEXQ-1070.a.jpg",
							},
						},
						{
							Name:  "Alexander McQUEEN White and clay Oversize Sneaker",
							Brand: "Alexander McQUEEN",
							Images: []string{
								"https://stok.store/cdn/shop/files/718139WIEE5_O_ALEXQ-8742.a.jpg",
							},
						},
						{
							Name:  "Alexander McQUEEN Black camera bag with leather details",
							Brand: "Alexander McQUEEN",
							Images: []string{
								"https://stok.store/cdn/shop/files/7262921AAQ0_O_ALEXQ-1000.a.jpg",
							},
						},
						{
							Name:  "A.P.C. 'GRACE SMALL' CROSSBODY BAG",
							Brand: "A.P.C.",
							Images: []string{
								"https://stok.store/cdn/shop/files/20230505000427217_A55---apc---COGFAF61413LZZ_1_M1.jpg",
								"https://stok.store/cdn/shop/files/20230505000427378_A55---apc---COGFAF61413LZZ_2_M1.jpg",
								"https://stok.store/cdn/shop/files/20230505000427514_A55---apc---COGFAF61413LZZ_3_M1.jpg",
								"https://stok.store/cdn/shop/files/20230505000432125_A55---apc---COGFAF61413LZZ_4_M1.jpg",
							},
						},
						{
							Name:  "A.P.C. JAMIE' CROSSBODY BAG",
							Brand: "A.P.C.",
							Images: []string{
								"https://stok.store/cdn/shop/files/20220221181331577_E52---apc---PXBMWF63412LZZBLACK_3_M1.jpg",
								"https://stok.store/cdn/shop/files/20220221181331639_E52---apc---PXBMWF63412LZZBLACK_4_M1.jpg",
							},
						},
					}

					for _, d := range requestData {
						jsonData, err := json.Marshal(d)
						if err != nil {
							return err
						}

						_, err = http.Post("http://localhost:3000/api/product", "application/json", bytes.NewBuffer(jsonData))
						if err != nil {
							return err
						}
					}

					fmt.Println("success insert data")
					return nil
				},
			},
		},
	}
}
