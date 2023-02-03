```go

//User represents users table in database
type User struct {
	ID                uint64         `gorm:"primary_key:auto_increment" json:"id"`
	Name              string         `gorm:"type:varchar(255)" json:"name"`
	Email             string         `gorm:"uniqueIndex;type:varchar(255)" json:"email"`
	ProfilePictureURL string         `gorm:"type:text" json:"profilePictureURL"`
	Password          string         `gorm:"->;<-;not null" json:"-"`
	Token             string         `gorm:"-" json:"token,omitempty"`
	Subjects          *[]Subject     `gorm:"many2many:SubjectUser;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"subjects"`
	CreatedAt         time.Time      `json:"createdAt"`
	UpdatedAt         time.Time      `json:"updatedAt"`
	DeletedAt         gorm.DeletedAt `json:"-"`
}

```

```go

type Issue struct {
	ID              uint64         `gorm:"primary_key:auto_increment" json:"id"`
	Title           string         `gorm:"type:varchar(255)" json:"title"`
	Description     string         `gorm:"type:text" json:"description"`
	IssueForeignId  string         `gorm:"type:text" json:"issueForeignId"`
	TargetTime      uint32         `json:"targetTime"`
	SpendingTime    uint32         `json:"spendingTime"`
	Progress        uint8          `json:"progress"`
	SubjectID       uint64         `gorm:"not null" json:"-"`
	Subject         Subject        `gorm:"foreignkey:SubjectID;" json:"-"`
	ParentIssueID   *uint64        `json:"parentIssueID"`
	StatusID        uint8          `gorm:"not null;default:1" json:"status"`
	ChildIssues     []Issue        `gorm:"foreignkey:ParentIssueID;" json:"-"`
	DependentIssues []Issue        `gorm:"many2many:DependentIssues;" json:"-"`
	Comments        []IssueComment `gorm:"foreignkey:ID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"-"`
	ReporterID      uint64         `gorm:"not null" json:"-"`
	Reporter        User           `gorm:"foreignkey:ReporterID;" json:"-"`
	AssignieID      *uint64        `json:"-"`
	Assignie        User           `gorm:"foreignkey:AssignieID;" json:"-"`
	CreatedAt       time.Time      `json:"createdAt"`
	UpdatedAt       time.Time      `json:"updatedAt"`
	DeletedAt       gorm.DeletedAt `json:"-"`
}
```

```go
type Subject struct {
	ID           uint64         `gorm:"primary_key:auto_increment" json:"id"`
	Title        string         `gorm:"type:varchar(255)" json:"title"`
	Description  string         `gorm:"type:text" json:"description"`
	RepoID       string         `gorm:"type:text" json:"repoId"`
	ProjectID    uint64         `gorm:"not null" json:"-"`
	Project      Project        `gorm:"foreignkey:ProjectID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"-"`
	Issues       []Issue        `json:"-"`
	Stages       []Stage        `gorm:"foreignkey:id;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"-"`
	User         []User         `gorm:"many2many:SubjectUser;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"-"`
	TeamLeaderID uint64         `gorm:"not null" json:"-"`
	TeamLeader   User           `gorm:"foreignkey:TeamLeaderID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"-"`
	CreatedAt    time.Time      `json:"createdAt"`
	UpdatedAt    time.Time      `json:"updatedAt"`
	DeletedAt    gorm.DeletedAt `json:"-"`
}
```