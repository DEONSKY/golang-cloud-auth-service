package main

import (
	"strconv"

	"github.com/gofiber/fiber/v2"

	"github.com/forfam/authentication-service/src/config"
	"github.com/forfam/authentication-service/src/files"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Post("/files", files.UploadFileEndpoint)

	app.Listen(":" + strconv.Itoa(config.HTTP_PORT))
}
