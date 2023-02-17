package organization

import "time"

type CreateOrganizationPayload struct {
	Name        string `json:"name" validate:"required,alpha,max=255"`
	Description string `json:"description" validate:"required,alpha"`
}

type OrganizationResponse struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func mapCreateOrganizationPayload(payload CreateOrganizationPayload) OrganizationEntity {
	return OrganizationEntity{
		Name:        payload.Name,
		Description: payload.Description,
	}
}

func mapEntity(entity *OrganizationEntity) OrganizationResponse {
	return OrganizationResponse{
		Id:          entity.Id,
		Name:        entity.Name,
		Description: entity.Description,
		CreatedAt:   entity.CreatedAt,
		UpdatedAt:   entity.UpdatedAt,
	}
}
