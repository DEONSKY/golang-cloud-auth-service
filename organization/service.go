package organization

import (
	"github.com/forfam/authentication-service/genericrepo"
	"github.com/forfam/authentication-service/postgres"
	"github.com/forfam/authentication-service/utils/pagination"
)

func GetOrganizationsPaginated(opt *pagination.PaginationOptions) (*pagination.PaginationResult[OrganizationResponse], error) {

	tx := postgres.AuthenticationDb.Model(OrganizationEntity{})
	return pagination.Paginate[OrganizationResponse](tx, opt)

}

func CreateOrganization(data *CreateOrganizationPayload) (*OrganizationEntity, error) {
	item := OrganizationEntity{
		Name:        data.Name,
		Description: data.Description,
	}

	if err := genericrepo.Create(&item, "Organization", *logger); err != nil {
		return nil, err
	}

	return &item, nil
}

func UpdateOrganization(id string, data *UpdateOrganizationPayload) (*OrganizationEntity, error) {
	item := OrganizationEntity{
		Id: id,
	}

	if err := genericrepo.Take(&item, "Organization", *logger); err != nil {
		return nil, err
	}

	if len(data.Name) > 0 {
		item.Name = data.Name
	}

	if len(data.Description) > 0 {
		item.Description = data.Description
	}

	if err := genericrepo.Update(&item, "Organization", *logger); err != nil {
		return nil, err
	}

	return &item, nil
}
