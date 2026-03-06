package main

import (
	"log"

	"product-api/internal/config"
	"product-api/internal/database"
	"product-api/internal/domain/product"
	"product-api/internal/handler"
	"product-api/internal/middleware"
	"product-api/internal/repository"
	"product-api/internal/service"

	"github.com/gofiber/fiber/v2"
)

func main() {

	cfg := config.Load()

	db, err := database.Connect(cfg.DB)
	if err != nil {
		log.Fatalf("database connection failed: %v", err)
	}

	if err := db.AutoMigrate(&product.Product{}); err != nil {
		log.Fatalf("auto-migrate failed: %v", err)
	}

	repo := repository.NewProductRepository(db)
	svc := service.NewProductService(repo)
	handler := handler.NewProductHandler(svc)

	app := fiber.New(fiber.Config{
		AppName:      cfg.App.Name,
		ErrorHandler: middleware.ErrorHandler,
	})

	middleware.Register(app, cfg)
	handler.RegisterRoutes(app)

	log.Printf("🚀  %s listening on :%s", cfg.App.Name, cfg.App.Port)
	log.Fatal(app.Listen(":" + cfg.App.Port))
}
