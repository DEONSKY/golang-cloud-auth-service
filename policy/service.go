package policy

import (
	"github.com/forfam/authentication-service/genericrepo"
	"github.com/forfam/authentication-service/organization"
	"github.com/forfam/authentication-service/postgres"
	"github.com/forfam/authentication-service/utils/pagination"
)

func CreatePolicy(data *CreatePolicyPayload) (*PolicyEntity, error) {
	item := PolicyEntity{
		Name:           data.Name,
		OrganizationId: data.OrganizationId,
	}

	if err := genericrepo.Take(&organization.OrganizationEntity{Id: item.OrganizationId}, "Organization", *logger); err != nil {
		return nil, err
	}

	if err := genericrepo.IsRelationNotExists(&item, []string{"Organization", "Policy"}, *logger); err != nil {
		return nil, err
	}

	if err := genericrepo.Create(&item, "Policy", *logger); err != nil {
		return nil, err
	}

	return &item, nil
}

func GetPoliciesPaginated(organizationId string, opt *pagination.PaginationOptions) (*pagination.PaginationResult[PolicyResponse], error) {

	tx := postgres.AuthenticationDb.Model(PolicyEntity{}).Where("organization_id = ?", organizationId)

	return pagination.Paginate[PolicyResponse](tx, opt)
}

func UpdatePolicy(id string, data *UpdatePolicyPayload) (*PolicyEntity, error) {
	item := PolicyEntity{
		Id: id,
	}

	if err := genericrepo.Take(&item, "Policy", *logger); err != nil {
		return nil, err
	}

	if len(data.Name) > 0 {
		item.Name = data.Name
	}

	if err := genericrepo.Update(&item, "Policy", *logger); err != nil {
		return nil, err
	}

	return &item, nil
}

func deletePolicy(id string) (*PolicyEntity, error) {
	item := PolicyEntity{
		Id: id,
	}

	if err := genericrepo.Take(&item, "Policy", *logger); err != nil {
		return nil, err
	}

	if err := genericrepo.Delete(&item, "Policy", *logger); err != nil {
		return nil, err
	}

	return &item, nil
}
