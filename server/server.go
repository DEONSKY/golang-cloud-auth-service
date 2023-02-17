package server

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"github.com/forfam/authentication-service/config"
)

var App *fiber.App
var Api fiber.Router

func init() {
	App = fiber.New()
	Api = App.Group("/api")
}

func Listen() {
	App.Listen(fmt.Sprintf(":%d", config.HTTP_PORT))
}
