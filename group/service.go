package group

import (
	"fmt"

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

	if err := postgres.AuthenticationDb.First(&organization.OrganizationEntity{Id: item.OrganizationId}).Error; err != nil {
		logger.Error(fmt.Sprintf(`Organization not found: %s`, err))
		return nil, err
	}

	if err := postgres.AuthenticationDb.Create(&item).Error; err != nil {
		logger.Error(fmt.Sprintf(`Something went wrong during creation of "Group": %s - Error: %s`, data, err))
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

	result := postgres.AuthenticationDb.First(&item)

	if result.Error != nil {
		logger.Error(fmt.Sprintf(`Something went wrong during find "Group"! Id: %d - Error: %s`, id, result.Error))
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		logger.Warning(fmt.Sprintf(`Group not found for update! Id: %d`, id))
		return nil, nil
	}

	if len(data.Name) > 0 {
		item.Name = data.Name
	}

	if len(data.Description) > 0 {
		item.Description = data.Description
	}

	result = postgres.AuthenticationDb.Save(&item)

	if result.Error != nil {
		logger.Error(fmt.Sprintf(`Something went wrong during update "Group"! Id: %d - Error: %s`, id, result.Error))
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		logger.Warning(fmt.Sprintf(`Group not updated! Id: %d - Data: %s`, id, data))
		return nil, nil
	}

	return &item, nil
}

func deleteGroup(id string) (*GroupEntity, error) {
	item := GroupEntity{
		Id: id,
	}

	result := postgres.AuthenticationDb.First(&item)

	if result.Error != nil {
		logger.Error(fmt.Sprintf(`Something went wrong during find "Group"! Id: %d - Error: %s`, id, result.Error))
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		logger.Warning(fmt.Sprintf(`Group not found for delete! Id: %d`, id))
		return nil, nil
	}

	result = postgres.AuthenticationDb.Delete(&item)

	if result.Error != nil {
		logger.Error(fmt.Sprintf(`Something went wrong during delete "Group"! Id: %d - Error: %s`, id, result.Error))
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		logger.Warning(fmt.Sprintf(`Group not deleted! Id: %d`, id))
		return nil, nil
	}

	return &item, nil
}
