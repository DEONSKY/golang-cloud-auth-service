package grouppolicy

import (
	"time"

	"github.com/forfam/authentication-service/group"
	"github.com/forfam/authentication-service/policy"
	"gorm.io/gorm"
)

type GroupPolicyEntity struct {
	Id        string              `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	PolicyId  string              `gorm:"type:uuid"`
	Policy    policy.PolicyEntity `gorm:"foreignkey:PolicyId;"`
	GroupId   string              `gorm:"type:uuid"`
	Group     group.GroupEntity   `gorm:"foreignkey:GroupId;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (entity *GroupPolicyEntity) TableName() string {
	return "group_policies"
}
