package config

import (
	"errors"

	"github.com/forfam/authentication-service/customerror"
	"github.com/gofiber/fiber/v2"
)

func NewErrorHandlerConfig() func(ctx *fiber.Ctx, err error) error {
	return func(ctx *fiber.Ctx, err error) error {
		// Status code defaults to 500
		code := fiber.StatusInternalServerError

		// Retrieve the custom status code if it's a *fiber.Error
		var e *fiber.Error
		if errors.As(err, &e) {
			code = e.Code
		}

		if re, ok := err.(*customerror.CoreError); ok {

			if errors.Is(re.Err, customerror.ValidationError) {
				return ctx.Status(re.HttpCode).JSON(re.MapToValidationErrorResponse())
			}

			return ctx.Status(re.HttpCode).JSON(re.MapToCoreErrorResponse())
		}

		// Set Content-Type: text/plain; charset=utf-8
		ctx.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)

		// Return status code with error message
		return ctx.Status(code).SendString(err.Error())
	}

}
