package model

import (
	"time"

	"gorm.io/gorm"
)

//Book struct represents books table in database
type Subject struct {
	ID           uint64         `gorm:"primary_key:auto_increment" json:"id"`
	Title        string         `gorm:"type:varchar(255)" json:"title"`
	Description  string         `gorm:"type:text" json:"description"`
	RepoID       string         `gorm:"type:text" json:"repoId"`
	ProjectID    uint64         `gorm:"not null" json:"-"`
	TeamLeaderID uint64         `gorm:"not null" json:"-"`
	CreatedAt    time.Time      `json:"createdAt"`
	UpdatedAt    time.Time      `json:"updatedAt"`
	DeletedAt    gorm.DeletedAt `json:"-"`
}