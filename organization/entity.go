package organization

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
