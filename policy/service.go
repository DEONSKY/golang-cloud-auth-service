package policy

import (
	"github.com/forfam/authentication-service/organization"
	"github.com/forfam/authentication-service/postgres"
	"github.com/forfam/authentication-service/utils/pagination"
	reslog "github.com/forfam/authentication-service/utils/resultlogger"
)

func CreatePolicy(data *CreatePolicyPayload) (*PolicyEntity, error) {
	item := PolicyEntity{
		Name:           data.Name,
		OrganizationId: data.OrganizationId,
	}

	result := postgres.AuthenticationDb.First(&organization.OrganizationEntity{Id: item.OrganizationId})

	result, err := reslog.LogGormResult(result, logger, item.OrganizationId, reslog.ErrorNotFoundLogMsg, "", "Organization")
	if result == nil {
		return nil, err
	}

	result = postgres.AuthenticationDb.Create(&item)
	result, err = reslog.LogGormResult(result, logger, "undefined", reslog.ErrorDuringCreateLogMsg, "", "Policy")
	if result == nil {
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

	result, err := reslog.LogGormResult(result, logger, id, reslog.ErrorDuringFindLogMsg, reslog.NotFoundForUpdateLogMsg, "Policy")
	if result == nil {
		return nil, err
	}

	if len(data.Name) > 0 {
		item.Name = data.Name
	}

	result = postgres.AuthenticationDb.Save(&item)

	result, err = reslog.LogGormResult(result, logger, id, reslog.ErrorDuringUpdateLogMsg, reslog.NotUpdatedLogMsg, "Policy")
	if result == nil {
		return nil, err
	}

	return &item, nil
}

func deletePolicy(id string) (*PolicyEntity, error) {
	item := PolicyEntity{
		Id: id,
	}

	result := postgres.AuthenticationDb.First(&item)

	result, err := reslog.LogGormResult(result, logger, id, reslog.ErrorDuringFindLogMsg, reslog.NotFoundForDeleteLogMsg, "Policy")
	if result == nil {
		return nil, err
	}

	result = postgres.AuthenticationDb.Delete(&item)

	result, err = reslog.LogGormResult(result, logger, id, reslog.ErrorDuringDeleteLogMsg, reslog.NotDeletedLogMsg, "Policy")
	if result == nil {
		return nil, err
	}

	return &item, nil
}
