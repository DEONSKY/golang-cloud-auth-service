package models

import (
	"time"

	"gorm.io/gorm"
)

type OrganizationEntity struct {
	Id          string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name        string `gorm:"type:varchar(255);not null"`
	Description string `gorm:"text"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}

func (entity *OrganizationEntity) TableName() string {
	return "organizations"
}

type CreateOrganizationPayload struct {
	Name        string `json:"name" validate:"required,max=255"`
	Description string `json:"description" validate:"required"`
}

type UpdateOrganizationPayload struct {
	Name        string `json:"name" validate:"omitempty,max=255"`
	Description string `json:"description" validate:"omitempty"`
}

type OrganizationResponse struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func MapOrganizationEntity(entity *OrganizationEntity) OrganizationResponse {
	return OrganizationResponse{
		Id:          entity.Id,
		Name:        entity.Name,
		Description: entity.Description,
		CreatedAt:   entity.CreatedAt,
		UpdatedAt:   entity.UpdatedAt,
	}
}
