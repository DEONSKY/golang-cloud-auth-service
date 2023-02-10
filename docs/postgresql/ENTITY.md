```go

//User represents users table in database
type User struct {
	ID                string         `gorm:"type:uuid;default:gen_random_uuid();primaryKey`
	Name              string         `gorm:"type:varchar(255)"`
	Email             string         `gorm:"uniqueIndex;type:varchar(255)"`
	ProfilePictureURL string         `gorm:"type:text"`
	Password          string         `gorm:"->;<-;not null"`
	Token             string         `gorm:"-"`
	Subjects          *[]Subject     `gorm:"many2many:SubjectUser;constraint:onUpdate:CASCADE,onDelete:CASCADE"`
	CreatedAt         time.Time      
	UpdatedAt         time.Time      
	DeletedAt         gorm.DeletedAt 
}

```

```go

type Issue struct {
	ID              string         `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Title           string         `gorm:"type:varchar(255)" `
	Description     string         `gorm:"type:text"`
	IssueForeignId  string         `gorm:"type:text"`
	TargetTime      uint32         
	SpendingTime    uint32         
	Progress        uint8          
	SubjectID       string         `gorm:"not null"`
	Subject         Subject        `gorm:"foreignkey:SubjectID;"`
	ParentIssueID   *string        
	StatusID        uint8          `gorm:"not null;default:1"`
	ChildIssues     []Issue        `gorm:"foreignkey:ParentIssueID;"`
	DependentIssues []Issue        `gorm:"many2many:DependentIssues;"`
	Comments        []IssueComment `gorm:"foreignkey:ID;constraint:onUpdate:CASCADE,onDelete:CASCADE"`
	ReporterID      string         `gorm:"not null"`
	Reporter        User           `gorm:"foreignkey:ReporterID;"`
	AssignieID      *string        
	Assignie        User           `gorm:"foreignkey:AssignieID;"`
	CreatedAt       time.Time      
	UpdatedAt       time.Time      
	DeletedAt       gorm.DeletedAt 
}
```

```go
type Subject struct {
	ID           string         `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Title        string         `gorm:"type:varchar(255)"`
	Description  string         `gorm:"type:text"`
	RepoID       string         `gorm:"type:text"`
	ProjectID    string         `gorm:"not null"`
	Project      Project        `gorm:"foreignkey:ProjectID;constraint:onUpdate:CASCADE,onDelete:CASCADE"`
	Issues       []Issue        
	Stages       []Stage        `gorm:"foreignkey:id;constraint:onUpdate:CASCADE,onDelete:CASCADE"`
	User         []User         `gorm:"many2many:SubjectUser;constraint:onUpdate:CASCADE,onDelete:CASCADE"`
	TeamLeaderID string         `gorm:"not null"`
	TeamLeader   User           `gorm:"foreignkey:TeamLeaderID;constraint:onUpdate:CASCADE,onDelete:CASCADE"`
	CreatedAt    time.Time      
	UpdatedAt    time.Time      
	DeletedAt    gorm.DeletedAt 
}
```