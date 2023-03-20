package group

import "time"

type CreateGroupPayload struct {
	Name           string `json:"name" validate:"required,max=255"`
	Description    string `json:"description" validate:"required"`
	OrganizationId string `json:"organizationId" validate:"required,uuid"`
}

type GroupResponse struct {
	Id             string `json:"id"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	OrganizationId string `json:"organizationId"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type UpdateGroupPayload struct {
	Name        string `json:"name" validate:"omitempty,max=255"`
	Description string `json:"description" validate:"required"`
}

func (ent GroupEntity) mapEntity() GroupResponse {
	return GroupResponse{
		Id:             ent.Id,
		Name:           ent.Name,
		Description:    ent.Description,
		OrganizationId: ent.OrganizationId,
		CreatedAt:      ent.CreatedAt,
		UpdatedAt:      ent.UpdatedAt,
	}
}
