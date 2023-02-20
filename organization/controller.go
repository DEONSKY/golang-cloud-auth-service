package organization

import (
	"github.com/gofiber/fiber/v2"

	"github.com/forfam/authentication-service/server"
	"github.com/forfam/authentication-service/utils/pagination"
)

func getPaginatedOrganizationList(ctx *fiber.Ctx) error {
	query := new(pagination.PaginationOptions)

	if err := ctx.QueryParser(query); err != nil {
		return err
	}

	res, err := GetOrganizationsPaginated(query)
	if err != nil {
		return err
	}

	return ctx.JSON(res)
}

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

func updateOrganizationHandler(ctx *fiber.Ctx) error {
	ctx.Accepts("application/json")

	body := new(UpdateOrganizationPayload)

	if err := ctx.BodyParser(body); err != nil {
		return err
	}

	if err := server.ValidateStruct[UpdateOrganizationPayload](*body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}

	item, err := UpdateOrganization(ctx.Params("id"), body)

	if err != nil {
		return err
	}

	if item == nil {
		return ctx.Status(fiber.StatusNotModified).Send(nil)
	}

	return ctx.Status(fiber.StatusOK).JSON(mapEntity(item))
}

func init() {
	organizationsGroup := server.Api.Group("/organizations")
	organizationsGroup.Post("/", createOrganizationHandler)
	organizationsGroup.Patch("/:id", updateOrganizationHandler)
	organizationsGroup.Get("/", getPaginatedOrganizationList)
}
