package organization

import (
	"github.com/gofiber/fiber/v2"

	"github.com/forfam/authentication-service/server"
)

func createOrganizationHandler(ctx *fiber.Ctx) error {
	ctx.Accepts("application/json")

	body := new(CreateOrganizationPayload)

	if err := ctx.BodyParser(body); err != nil {
		return err
	}

	if validationErrs := server.ValidateStruct[CreateOrganizationPayload](*body); validationErrs != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(validationErrs)
	}

	item, err := CreateOrganization(body)

	if err != nil {
		return err
	}

	return ctx.JSON(mapEntity(item))
}

func init() {
	organizationsGroup := server.Api.Group("/organizations")
	organizationsGroup.Post("/", createOrganizationHandler)
}
