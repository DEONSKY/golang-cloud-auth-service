package fiber

import (
	"github.com/forfam/authentication-service/server"
	"github.com/gofiber/fiber/v2"
)

func ParseBodyAndValidate[T any](ctx *fiber.Ctx) (*T, error) {
	body := new(T)

	if err := ctx.BodyParser(body); err != nil {
		return nil, ctx.Status(fiber.StatusBadRequest).JSON(err)
	}

	if validationErrs := server.ValidateStruct(*body); validationErrs != nil {
		return nil, ctx.Status(fiber.StatusBadRequest).JSON(validationErrs)
	}

	return body, nil
}

func ParseQueryAndValidate[T any](ctx *fiber.Ctx) (*T, error) {
	query := new(T)

	if err := ctx.QueryParser(query); err != nil {
		return nil, ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	if validationErrs := server.ValidateStruct(*query); validationErrs != nil {
		return nil, ctx.Status(fiber.StatusBadRequest).JSON(validationErrs)
	}

	return query, nil
}
