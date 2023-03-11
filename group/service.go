package group

import (
	"github.com/forfam/authentication-service/genericrepo"
	"github.com/forfam/authentication-service/organization"
	"github.com/forfam/authentication-service/postgres"
	"github.com/forfam/authentication-service/utils/pagination"
)

func CreateGroup(data *CreateGroupPayload) (*GroupEntity, error) {
	item := GroupEntity{
		Name:           data.Name,
		Description:    data.Description,
		OrganizationId: data.OrganizationId,
	}

	if err := genericrepo.Take(&organization.OrganizationEntity{Id: item.OrganizationId}, "Organization", *logger); err != nil {
		return nil, err
	}

	if err := genericrepo.Create(&item, "Policy", *logger); err != nil {
		return nil, err
	}

	return &item, nil
}

func GetGroupsPaginated(organizationId string, opt *pagination.PaginationOptions) (*pagination.PaginationResult[GroupResponse], error) {

	tx := postgres.AuthenticationDb.Model(GroupEntity{}).Where("organization_id = ?", organizationId)

	return pagination.Paginate[GroupResponse](tx, opt)
}

func UpdateGroup(id string, data *UpdateGroupPayload) (*GroupEntity, error) {
	item := GroupEntity{
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

	if err := genericrepo.Update(&item, "Group", *logger); err != nil {
		return nil, err
	}

	return &item, nil
}

func deleteGroup(id string) (*GroupEntity, error) {
	item := GroupEntity{
		Id: id,
	}

	if err := genericrepo.Take(&item, "Group", *logger); err != nil {
		return nil, err
	}

	if err := genericrepo.Delete(&item, "Group", *logger); err != nil {
		return nil, err
	}

	return &item, nil
}
