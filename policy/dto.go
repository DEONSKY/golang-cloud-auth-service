package policy

import "time"

type CreatePolicyPayload struct {
	Name           string `json:"name" validate:"required,max=255"`
	OrganizationId string `json:"organizationId" validate:"required,uuid"`
}

type PolicyResponse struct {
	Id             string `json:"id"`
	Name           string `json:"name"`
	OrganizationId string `json:"organizationId"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type UpdatePolicyPayload struct {
	Name string `json:"name" validate:"omitempty,max=255"`
}

func (ent PolicyEntity) mapEntity() PolicyResponse {
	return PolicyResponse{
		Id:             ent.Id,
		Name:           ent.Name,
		OrganizationId: ent.OrganizationId,
		CreatedAt:      ent.CreatedAt,
		UpdatedAt:      ent.UpdatedAt,
	}
}
