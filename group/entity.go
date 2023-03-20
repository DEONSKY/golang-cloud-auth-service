package group

import (
	"fmt"
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

func (entity GroupEntity) String() string {
	return fmt.Sprintf(`Id: %s, Name: %s, Description: %s, OrganizationId: %s, CreatedAt: %s, UpdatedAt %s, DeletedAt %s`,
		entity.Id, entity.Name, entity.Description, entity.OrganizationId, entity.CreatedAt, entity.UpdatedAt, entity.DeletedAt.Time)
}
