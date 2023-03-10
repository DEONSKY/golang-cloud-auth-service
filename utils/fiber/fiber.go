package fiber

import (
	"github.com/forfam/authentication-service/customerror"
	"github.com/forfam/authentication-service/server"
	"github.com/gofiber/fiber/v2"
)

func ParseBodyAndValidate[T any](ctx *fiber.Ctx) (*T, error) {
	body := new(T)

	if err := ctx.BodyParser(body); err != nil {
		return nil, customerror.NewBadRequestError("Something went wrong when body parsing", &err, nil)
	}

	if validationErrs := server.ValidateStruct(*body); validationErrs != nil {
		return nil, customerror.NewValidationError("Body not fitting validation rules", &validationErrs, nil)
	}

	return body, nil
}

func ParseQueryAndValidate[T any](ctx *fiber.Ctx) (*T, error) {
	query := new(T)

	if err := ctx.QueryParser(query); err != nil {
		return nil, customerror.NewBadRequestError("Something went wrong when query parsing", &err, nil)
	}

	if validationErrs := server.ValidateStruct(*query); validationErrs != nil {
		return nil, customerror.NewValidationError("Query not fitting validation rules", &validationErrs, nil)
	}

	return query, nil
}
