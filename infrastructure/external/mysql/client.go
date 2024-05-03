package mysql

import "database/sql"

type MysqlClient interface {
}

type mysqlClient struct {
	db *sql.DB
}

func NewMysqlClient(db *sql.DB) MysqlClient {
	return &mysqlClient{
		db: db,
	}
}
