package group

import (
	"time"

	"github.com/forfam/authentication-service/organization"
	"gorm.io/gorm"
)

type GroupEntity struct {
	Id             string                          `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name           string                          `gorm:"type:varchar(255);not null"`
	Description    string                          `gorm:"text"`
	OrganizationId string                          `gorm:"type:uuid"`
	Organization   organization.OrganizationEntity `gorm:"foreignkey:OrganizationId;"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt
}

func (entity *GroupEntity) TableName() string {
	return "groups"
}
