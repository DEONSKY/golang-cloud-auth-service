package group

import (
	"github.com/gofiber/fiber/v2"

	"github.com/forfam/authentication-service/server"
	fiberutil "github.com/forfam/authentication-service/utils/fiber"
	"github.com/forfam/authentication-service/utils/pagination"
)

func createGroupHandler(ctx *fiber.Ctx) error {
	ctx.Accepts("application/json")

	createGroupPayload, err := fiberutil.ParseBodyAndValidate[CreateGroupPayload](ctx)
	if err != nil {
		return err
	}

	item, err := CreateGroup(createGroupPayload)

	if err != nil {
		return err
	}

	return ctx.JSON(item.mapEntity())
}

func getPaginatedGroupsList(ctx *fiber.Ctx) error {
	paginationPayload, err := fiberutil.ParseQueryAndValidate[pagination.PaginationOptions](ctx)
	if err != nil {
		return err
	}

	res, err := GetGroupsPaginated(ctx.Params("organizationId"), paginationPayload)
	if err != nil {
		return err
	}

	return ctx.JSON(res)
}

func updateGroupHandler(ctx *fiber.Ctx) error {
	ctx.Accepts("application/json")

	updateGroupPayload, err := fiberutil.ParseBodyAndValidate[UpdateGroupPayload](ctx)
	if err != nil {
		return err
	}

	item, err := UpdateGroup(ctx.Params("id"), updateGroupPayload)
	if err != nil {
		return err
	}

	if item == nil {
		return ctx.Status(fiber.StatusNotModified).Send(nil)
	}

	return ctx.Status(fiber.StatusOK).JSON(item.mapEntity())
}

func deleteGroupHandle(ctx *fiber.Ctx) error {
	ctx.Accepts("application/json")

	item, err := deleteGroup(ctx.Params("id"))
	if err != nil {
		return err
	}

	if item == nil {
		return ctx.Status(fiber.StatusNotModified).Send(nil)
	}

	return ctx.Status(fiber.StatusOK).JSON(item.mapEntity())
}

func init() {
	policyGroup := server.Api.Group("/groups")
	policyGroup.Post("/", createGroupHandler)
	policyGroup.Patch("/:id", updateGroupHandler)
	policyGroup.Get("/:organizationId", getPaginatedGroupsList)
	policyGroup.Delete("/:id", deleteGroupHandle)
}
