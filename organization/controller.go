package organization

import (
	"github.com/gofiber/fiber/v2"

	"github.com/forfam/authentication-service/server"
	fiberutil "github.com/forfam/authentication-service/utils/fiber"
	"github.com/forfam/authentication-service/utils/pagination"
)

func getPaginatedOrganizationList(ctx *fiber.Ctx) error {

	query, err := fiberutil.ParseBodyAndValidate[pagination.PaginationOptions](ctx)
	if err != nil {
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

	body, err := fiberutil.ParseBodyAndValidate[CreateOrganizationPayload](ctx)
	if err != nil {
		return err
	}

	item, err := CreateOrganization(body)

	if err != nil {
		return err
	}

	return ctx.JSON(mapEntity(item))
}

func updateOrganizationHandler(ctx *fiber.Ctx) error {
	ctx.Accepts("application/json")

	body, err := fiberutil.ParseBodyAndValidate[UpdateOrganizationPayload](ctx)
	if err != nil {
		return err
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
