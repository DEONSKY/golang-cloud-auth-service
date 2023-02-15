package main

import (
	"strconv"

	"github.com/gofiber/fiber/v2"

	"github.com/forfam/authentication-service/config"
	"github.com/forfam/authentication-service/files"
	"github.com/forfam/authentication-service/postgres"
)

func main() {
	postgres.InitAuthenticationDb()
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Post("/files", files.UploadFileEndpoint)

	app.Listen(":" + strconv.Itoa(config.HTTP_PORT))
}
