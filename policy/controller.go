package policy

import (
	"github.com/gofiber/fiber/v2"

	"github.com/forfam/authentication-service/server"
	fiberutil "github.com/forfam/authentication-service/utils/fiber"
	"github.com/forfam/authentication-service/utils/pagination"
)

func createPolicyHandler(ctx *fiber.Ctx) error {
	ctx.Accepts("application/json")

	createPolicyPayload, err := fiberutil.ParseBodyAndValidate[CreatePolicyPayload](ctx)
	if err != nil {
		return err
	}

	item, err := CreatePolicy(createPolicyPayload)

	if err != nil {
		return err
	}

	return ctx.JSON(item.mapEntity())
}

func getPaginatedPoliciesList(ctx *fiber.Ctx) error {
	paginationPayload, err := fiberutil.ParseQueryAndValidate[pagination.PaginationOptions](ctx)
	if err != nil {
		return err
	}

	res, err := GetPoliciesPaginated(ctx.Params("organizationId"), paginationPayload)
	if err != nil {
		return err
	}

	return ctx.JSON(res)
}

func updatePolicyHandler(ctx *fiber.Ctx) error {
	ctx.Accepts("application/json")

	updatePolicyPayload, err := fiberutil.ParseBodyAndValidate[UpdatePolicyPayload](ctx)
	if err != nil {
		return err
	}

	item, err := UpdatePolicy(ctx.Params("id"), updatePolicyPayload)
	if err != nil {
		return err
	}

	if item == nil {
		return ctx.Status(fiber.StatusNotModified).Send(nil)
	}

	return ctx.Status(fiber.StatusOK).JSON(item.mapEntity())
}

func deletePolicyHandle(ctx *fiber.Ctx) error {
	ctx.Accepts("application/json")

	item, err := deletePolicy(ctx.Params("id"))
	if err != nil {
		return err
	}

	if item == nil {
		return ctx.Status(fiber.StatusNotModified).Send(nil)
	}

	return ctx.Status(fiber.StatusOK).JSON(item.mapEntity())
}

func init() {
	policyGroup := server.Api.Group("/policies")
	policyGroup.Post("/", createPolicyHandler)
	policyGroup.Patch("/:id", updatePolicyHandler)
	policyGroup.Get("/:organizationId", getPaginatedPoliciesList)
	policyGroup.Delete("/:id", deletePolicyHandle)
}
