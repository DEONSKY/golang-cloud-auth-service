package grouppolicy

import (
	"github.com/gofiber/fiber/v2"

	"github.com/forfam/authentication-service/server"
	fiberutil "github.com/forfam/authentication-service/utils/fiber"
)

func createPolicyHandler(ctx *fiber.Ctx) error {
	ctx.Accepts("application/json")

	creategroupPolicyPayload, err := fiberutil.ParseBodyAndValidate[AddPolicyToGroupPayload](ctx)
	if err != nil {
		return err
	}

	item, err := CreatePolicy(creategroupPolicyPayload)
	if err != nil {
		return err
	}

	return ctx.JSON(item.mapEntity())
}

/*
func getPaginatedPoliciesOfGroup(ctx *fiber.Ctx) error {
	paginationPayload, err := fiberutil.ParseQueryAndValidate[pagination.PaginationOptions](ctx)
	if paginationPayload == nil || err != nil {
		return err
	}

	res, err := GetGroupsPaginated(ctx.Params("organizationId"), paginationPayload)
	if err != nil {
		return err
	}

	return ctx.JSON(res)
}*/

func deleteGroupPolicyByIdHandle(ctx *fiber.Ctx) error {
	ctx.Accepts("application/json")

	item, err := deleteGroupPolicyById(ctx.Params("id"))
	if err != nil {
		return err
	}

	if item == nil {
		return ctx.Status(fiber.StatusNotModified).Send(nil)
	}

	return ctx.Status(fiber.StatusOK).JSON(item.mapEntity())
}

func init() {
	policyGroup := server.Api.Group("/group-policies")
	policyGroup.Post("/", createPolicyHandler)
	policyGroup.Delete("/:id", deleteGroupPolicyByIdHandle)
}
