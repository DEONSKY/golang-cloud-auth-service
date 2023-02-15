
## Example Save Method Without Transaction

```go
type SubjectCreateRequest struct {
	Title        string `form:"title" validate:"required,max=32"`
	Description  string `form:"description" validate:"required,max=255"`
	ProjectId    uint64 `form:"projectId" binding:"required"`
	TeamLeaderId uint64 `form:"teamLeaderId" binding:"required"`
}
```

```go
type Subject struct {
	Id           uint64         `gorm:"primary_key:auto_increment"`
	Title        string         `gorm:"type:varchar(255)"`
	Description  string         `gorm:"type:text"`
	RepoId       string         `gorm:"type:text"`
	ProjectId    uint64         `gorm:"not null"`
	Project      Project        `gorm:"foreignkey:ProjectId;constraint:onUpdate:CASCADE,onDelete:CASCADE"`
	Issues       []Issue
	Stages       []Stage        `gorm:"foreignkey:id;constraint:onUpdate:CASCADE,onDelete:CASCADE"`
	User         []User         `gorm:"many2many:SubjectUser;constraint:onUpdate:CASCADE,onDelete:CASCADE"`
	TeamLeaderId uint64         `gorm:"not null"`
	TeamLeader   User           `gorm:"foreignkey:TeamLeaderId;constraint:onUpdate:CASCADE,onDelete:CASCADE"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt
}
```

We are mapping SubjectCreateRequest dto to Subject model inside service

```go
	subjectToCreate := model.Subject{}
	err := smapping.FillStruct(&subjectToCreate, smapping.MapFields(&subjectCreateDTO))
```

After that we are giving our repository for create process

```go
func InsertSubject(subject model.Subject) (*model.Subject, error) {

	if result := config.DB.Save(&subject); result.Error != nil {
	    return nil, result.Error
    }
	return &subject, nil
}
```

## Basic Get Method

```go
type SubjectNavTreeResponse struct {
	Id          string
	Title       string
	Description string
	ProjectId   string
}
```

```go
func GetSubjectsByUserId(userId string) ([]response.SubjectNavTreeResponse, error) {

	var subjectNavTreeResponse []response.SubjectNavTreeResponse

	if result := config.DB.Model(&model.Subject{}).
		Joins("INNER JOIN subject_users su on su.subject_id = id").
		Where("su.user_id", userId).Order("Id").Find(&subjectNavTreeResponse); result.Error != nil {
		return nil, result.Error
	}
	return subjectNavTreeResponse, nil

}
```

## Dynamic query creation with output complex nested dto by chaining

If we are using query DTO with optional fields, repository query needs to be dynamic too
```go
type IssueGetQuery struct {
	SubjectId      *string `query:"subjectId"`
	ProjectId      *string `query:"projectId"`
	ReporterId     *string `query:"reporterId"`
	AssignieId     *string `query:"assignieId"`
	Status         *uint8  `query:"status"`
	ParentIssueId  *string `query:"parentIssueId"`
	GetOnlyOrphans *bool   `query:"getOnlyOrphans"`
}
```

```go
func GetIssues(issueGetQuery *request.IssueGetQuery, userId string) ([]response.IssueResponse, error) {

	var issues []response.IssueResponse
	chain := config.DB.Model(&model.Issue{}).
		Preload("ChildIssues").//Loads child entity
		Preload("DependentIssues").
		Preload("Assignie").
		Preload("Reporter").
		Joins("INNER JOIN subjects s on subject_id = s.id").
		Joins("INNER JOIN subject_users su on su.subject_id = s.id").
		Where("user_id", userId)
	
	//Optional Find Clause
	if issueGetQuery.ReporterId != nil {
		chain = chain.Where("reporter_id", issueGetQuery.ReporterId)
	}

	if result := chain.Find(&issues); result.Error != nil {
		return nil, result.Error
	}

	return issues, nil
}
```

## Insert Association

In this example We are claiming required model from our another repositories

```go
	issue, err := service.issueRepository.FindIssueByAccess(issueId, userId)
```

```go
	dependentIssue, err := service.issueRepository.FindIssueByAccess(dependentIssueId, userId)
```

```go
func InsertDependentIssueAssociation(issue model.Issue, dependentIssue model.Issue) (*model.Issue, error) {
	if err := config.DB.Model(&issue).Omit("DependentIssues.*").Association("DependentIssues").Append(&dependentIssue); err != nil {
		return nil, err
	}
	return &issue, nil
}
```

Example user update method
```go
func UpdateUser(user model.User) model.User {
	if user.Password != "" {
		user.Password = hashAndSalt([]byte(user.Password))
	} else {
		var tempUser model.User
		config.DB.Find(&tempUser, user.Id)
		user.Password = tempUser.Password
	}
	config.DB.Save(&user)
	return user
}
```

## Using transactions with querys

With gorm create, update, delete  methods has transactions as default.
But we can create transactions like this:
```go
db.Transaction(func(tx *gorm.DB) error {  
	// do some database operations in the transaction (use 'tx' from this point, not 'db')  
 	if err := tx.Create(&Animal{Name: "Giraffe"}).Error; err != nil {  
 		// return any error will rollback  
 		return err  
 	}  
  
 	if err := tx.Create(&Animal{Name: "Lion"}).Error; err != nil {  
 		return err  
 	}  
  
 	// return nil will commit the whole transaction  
 	return nil  
})
```

For more info:
https://gorm.io/docs/transactions.html