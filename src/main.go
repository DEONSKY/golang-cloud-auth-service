package main

import (
	"strconv"

	"github.com/forfam/authentication-service/src/config"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":" + strconv.Itoa(config.HTTP_PORT))
}
