package policy

import (
	"fmt"
	"time"

	"github.com/forfam/authentication-service/organization"
	"gorm.io/gorm"
)

type PolicyEntity struct {
	Id             string                          `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name           string                          `gorm:"type:varchar(255);not null"`
	OrganizationId string                          `gorm:"type:uuid"`
	Organization   organization.OrganizationEntity `gorm:"foreignkey:OrganizationId;"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt
}

func (entity *PolicyEntity) TableName() string {
	return "policies"
}

func (entity PolicyEntity) String() string {
	return fmt.Sprintf(`Id: %s, Name: %s, OrganizationId: %s, CreatedAt: %s, UpdatedAt %s, DeletedAt %s`,
		entity.Id, entity.Name, entity.OrganizationId, entity.CreatedAt, entity.UpdatedAt, entity.DeletedAt.Time)
}
