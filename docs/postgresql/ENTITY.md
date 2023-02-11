Gorm automaticly maps all fields inside Entities

## Example Models

for more info: https://gorm.io/docs/models.html

```go

//User represents users table in database
type User struct {
	Id                string         `gorm:"type:uuid;default:gen_random_uuid();primaryKey`
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
	Id              string         `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Title           string         `gorm:"type:varchar(255)" `
	Description     string         `gorm:"type:text"`
	IssueForeignId  string         `gorm:"type:text"`
	TargetTime      uint32         
	SpendingTime    uint32         
	Progress        uint8          
	SubjectId       string         `gorm:"not null"`
	Subject         Subject        `gorm:"foreignkey:SubjectId;"`
	ParentIssueId   *string        
	StatusId        uint8          `gorm:"not null;default:1"`
	ChildIssues     []Issue        `gorm:"foreignkey:ParentIssueId;"`
	DependentIssues []Issue        `gorm:"many2many:DependentIssues;"`
	Comments        []IssueComment `gorm:"foreignkey:Id;constraint:onUpdate:CASCADE,onDelete:CASCADE"`
	ReporterId      string         `gorm:"not null"`
	Reporter        User           `gorm:"foreignkey:ReporterId;"`
	AssignieId      *string        
	Assignie        User           `gorm:"foreignkey:AssignieId;"`
	CreatedAt       time.Time      
	UpdatedAt       time.Time      
	DeletedAt       gorm.DeletedAt 
}
```

```go
type Subject struct {
	Id           string         `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Title        string         `gorm:"type:varchar(255)"`
	Description  string         `gorm:"type:text"`
	RepoId       string         `gorm:"type:text"`
	ProjectId    string         `gorm:"not null"`
	Project      Project        `gorm:"foreignkey:ProjectId;constraint:onUpdate:CASCADE,onDelete:CASCADE"`
	Issues       []Issue        
	Stages       []Stage        `gorm:"foreignkey:id;constraint:onUpdate:CASCADE,onDelete:CASCADE"`
	User         []User         `gorm:"many2many:SubjectUser;constraint:onUpdate:CASCADE,onDelete:CASCADE"`
	TeamLeaderId string         `gorm:"not null"`
	TeamLeader   User           `gorm:"foreignkey:TeamLeaderId;constraint:onUpdate:CASCADE,onDelete:CASCADE"`
	CreatedAt    time.Time      
	UpdatedAt    time.Time      
	DeletedAt    gorm.DeletedAt 
}
```