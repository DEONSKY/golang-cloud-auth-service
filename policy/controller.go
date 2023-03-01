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
	if createPolicyPayload == nil || err != nil {
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
	if paginationPayload == nil || err != nil {
		return err
	}

	res, err := GetPoliciesPaginated(ctx.Params("organizationId"), paginationPayload)
	if err != nil {
		return err
	}

	return ctx.JSON(res)
}

func init() {
	policyGroup := server.Api.Group("/policies")
	policyGroup.Post("/", createPolicyHandler)
	policyGroup.Get("/:organizationId", getPaginatedPoliciesList)
}
