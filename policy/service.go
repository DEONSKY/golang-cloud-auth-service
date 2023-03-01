package policy

import (
	"fmt"

	"github.com/forfam/authentication-service/organization"
	"github.com/forfam/authentication-service/postgres"
	"github.com/forfam/authentication-service/utils/pagination"
)

func CreatePolicy(data *CreatePolicyPayload) (*PolicyEntity, error) {
	item := PolicyEntity{
		Name:           data.Name,
		OrganizationId: data.OrganizationId,
	}

	if err := postgres.AuthenticationDb.First(&organization.OrganizationEntity{Id: item.OrganizationId}).Error; err != nil {
		logger.Error(fmt.Sprintf(`Organization not found: %s`, err))
		return nil, err
	}

	if err := postgres.AuthenticationDb.Create(&item).Error; err != nil {
		logger.Error(fmt.Sprintf(`Something went wrong during creation of "Policy": %s - Error: %s`, data, err))
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

	result := postgres.AuthenticationDb.First(&item)

	if result.Error != nil {
		logger.Error(fmt.Sprintf(`Something went wrong during find "Policy"! Id: %d - Error: %s`, id, result.Error))
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		logger.Warning(fmt.Sprintf(`Policy not found for update! Id: %d`, id))
		return nil, nil
	}

	if len(data.Name) > 0 {
		item.Name = data.Name
	}

	result = postgres.AuthenticationDb.Save(&item)

	if result.Error != nil {
		logger.Error(fmt.Sprintf(`Something went wrong during update "Policy"! Id: %d - Error: %s`, id, result.Error))
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		logger.Warning(fmt.Sprintf(`Policy not updated! Id: %d - Data: %s`, id, data))
		return nil, nil
	}

	return &item, nil
}
