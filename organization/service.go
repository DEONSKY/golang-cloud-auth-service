package organization

import (
	"fmt"

	"github.com/forfam/authentication-service/postgres"
)

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
