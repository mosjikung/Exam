package middleware

import (
	"product-api/internal/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func Register(app *fiber.App, cfg *config.Config) {
	app.Use(recover.New())
	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${status} ${latency} ${method} ${path} | ${locals:requestid}\n",
	}))
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,DELETE,OPTIONS",
		AllowHeaders: "Content-Type,X-Request-ID",
	}))

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
			"app":    cfg.App.Name,
			"env":    cfg.App.Env,
		})
	})
}

func ErrorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	msg := "internal server error"

	var fe *fiber.Error
	if e, ok := err.(*fiber.Error); ok {
		fe = e
		code = fe.Code
		msg = fe.Message
	}

	return c.Status(code).JSON(fiber.Map{
		"error":      msg,
		"statusCode": code,
		"path":       c.Path(),
	})
}
