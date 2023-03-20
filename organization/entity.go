package organization

import (
	"fmt"
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

func (entity OrganizationEntity) String() string {
	return fmt.Sprintf(`Id: %s, Name: %s, Description: %s, CreatedAt: %s, UpdatedAt %s, DeletedAt %s`,
		entity.Id, entity.Name, entity.Description, entity.CreatedAt, entity.UpdatedAt, entity.DeletedAt.Time)
}
