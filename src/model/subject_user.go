package model

// Book struct represents books table in database
type SubjectUser struct {
	ID        uint64  `gorm:"primary_key:auto_increment" json:"id"`
	UserID    uint64  `gorm:"not null" json:"-"`
	User      User    `gorm:"foreignkey:UserID;"`
	SubjectID uint64  `gorm:"not null" json:"-"`
	Subject   Subject `gorm:"foreignkey:SubjectID;"`
}
