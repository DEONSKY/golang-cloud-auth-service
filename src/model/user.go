package model

import (
	"time"

	"gorm.io/gorm"
)

//User represents users table in database
type User struct {
	ID                uint64         `gorm:"primary_key:auto_increment" json:"id"`
	Name              string         `gorm:"type:varchar(255)" json:"name"`
	Email             string         `gorm:"uniqueIndex;type:varchar(255)" json:"email"`
	ProfilePictureURL string         `gorm:"type:text" json:"profilePictureURL"`
	Password          string         `gorm:"->;<-;not null" json:"-"`
	Token             string         `gorm:"-" json:"token,omitempty"`
	CreatedAt         time.Time      `json:"createdAt"`
	UpdatedAt         time.Time      `json:"updatedAt"`
	DeletedAt         gorm.DeletedAt `json:"-"`
}