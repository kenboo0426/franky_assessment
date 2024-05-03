package interfaces

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kenboo0426/franky_assessment/application"
	"github.com/kenboo0426/franky_assessment/infrastructure/dao"
	"github.com/kenboo0426/franky_assessment/infrastructure/external/mysql"
	"github.com/kenboo0426/franky_assessment/infrastructure/service"
)

func InitializeHTTPServer() {
	ctx := context.Background()

	db := mysql.NewDBConnection()

	// Repository
	productRepository := dao.NewProductRepository(db)
	productImageRepository := dao.NewProductImageRepository(db)
	brandRepository := dao.NewBrandRepository(db)

	// Service
	productService := service.NewProductService(db)

	// Usecase
	productUsecase := application.NewProductUsecase(brandRepository, productRepository, productImageRepository, productService)

	// Handler
	productHandler := NewProductHandler(productUsecase)

	g := gin.Default()
	g.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{
			"HEAD",
			"GET",
			"POST",
			"PUT",
			"PATCH",
			"DELETE",
			"OPTIONS",
		},
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Cookie",
			"Accept-Encoding",
			"Authorization",
		},
		AllowCredentials: true,
	}))

	api := g.Group("/api")
	product := api.Group("/product")
	{
		product.GET("/", productHandler.GetAll)
		product.GET("/search", productHandler.Search)
		product.POST("/", productHandler.Create)
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", os.Getenv("BACKEND_PORT")),
		Handler: g,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
}
