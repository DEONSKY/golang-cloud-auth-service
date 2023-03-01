package organization

import (
	"fmt"

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

	if err := postgres.AuthenticationDb.Create(&item).Error; err != nil {
		logger.Error(fmt.Sprintf(`Something went wrong during creation of "Organization": %s - Error: %s`, data, err))
		return nil, err
	}

	return &item, nil
}

func UpdateOrganization(id string, data *UpdateOrganizationPayload) (*OrganizationEntity, error) {
	item := OrganizationEntity{
		Id: id,
	}

	result := postgres.AuthenticationDb.First(&item)

	if result.Error != nil {
		logger.Error(fmt.Sprintf(`Something went wrong during find "Organization": %d - Error: %s`, id, result.Error))
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		logger.Warning(fmt.Sprintf(`Organization not found for update! OrganizationId: %d`, id))
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
		logger.Error(fmt.Sprintf(`Something went wrong during update "Organization"! OrganizationId: %d - Error: %s`, id, result.Error))
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		logger.Warning(fmt.Sprintf(`Organization not updated! OrganizationId: %d - Data: %s`, id, data))
		return nil, nil
	}

	return &item, nil
}
