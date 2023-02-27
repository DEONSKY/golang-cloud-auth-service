package policy

import "time"

type CreatePolicyPayload struct {
	Name           string `json:"name" validate:"required,max=255"`
	OrganizationId string `json:"organizationId" validate:"required,uuid"`
}

type PolicyResponse struct {
	Id             string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name           string `gorm:"type:varchar(255);not null"`
	OrganizationId string `gorm:"type:uuid"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
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
