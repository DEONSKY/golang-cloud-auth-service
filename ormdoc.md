
Example Save Method Without Transaction
```go
func InsertSubject(subject model.Subject) (*model.Subject, error) {

	if result := config.DB.Save(&subject); result.Error != nil {

	return nil, result.Error
}
	return &subject, nil
}
```
Basic Get Method
```go
func GetSubjectsByUserId(userID uint64) ([]response.SubjectNavTreeResponse, error) {

	var subjectNavTreeResponse []response.SubjectNavTreeResponse

	if result := config.DB.Model(&model.Subject{}).
		Joins("INNER JOIN subject_users su on su.subject_id = id").
		Where("su.user_id", userID).Order("ID").Find(&subjectNavTreeResponse); result.Error != nil {
		return nil, result.Error
	}
	return subjectNavTreeResponse, nil

}
```

Dynamic query creation with output complex nested dto by chaining
```go
func (db *issueConnection) GetIssues(issueGetQuery *request.IssueGetQuery, userID uint64) ([]response.IssueResponse, error) {

	var issues []response.IssueResponse
	chain := config.DB.Model(&model.Issue{}).
	Preload("ChildIssues").//Loads child endtity
	Preload("DependentIssues").
	Preload("Assignie").
	Preload("Reporter").
	Joins("INNER JOIN subjects s on subject_id = s.id").
	Joins("INNER JOIN subject_users su on su.subject_id = s.id").
	Where("user_id", userID)
	
	//Optional Find Clause
	if issueGetQuery.ReporterID != nil {
		chain = chain.Where("reporter_id", issueGetQuery.ReporterID)
	}


	if result := chain.Find(&issues); result.Error != nil {
		return nil, result.Error
	}
	return issues, nil
}
```
Also we may need to make changes on the for using associations with dto
```go
type IssueKanbanResponse struct {
	Status StatusResponse `json:"status"`
	Issues []IssueResponse `json:"issues"`
}

type IssueResponse struct {
	ID uint64 `json:"id"`
	StatusID uint32 `json:"statusID"`
	Status StatusResponse `gorm:"-" json:"status"`
	Reporter UserLabelResponse `json:"reporter"`
	Assignie *UserLabelResponse `json:"assignie"`
	ChildIssues []*LeafIssueResponse `gorm:"foreignkey:ParentIssueID;" json:"issues"`
	DependentIssues []*LeafIssueResponse `gorm:"many2many:DependentIssues;foreignkey:ID;joinForeignKey:issueID;References:ID;joinReferences:dependentIssueID" json:"dependentIssues"`
}

type LeafIssueResponse struct {
	ID uint64 `json:"id"`
	Reporter UserLabelResponse `json:"reporter"`
	AssignieID *uint64 `json:"assignieID"`
	Assignie *UserLabelResponse `json:"assignie"`
}

// TableName overrides the table name for smart select

func (LeafIssueResponse) TableName() string {
	return "issues"
}

type IssueCommentResponse struct {
	Context string `json:"context"`
	IssueID uint64 `json:"issueID"`
	CreatorID uint64 `json:"-"`
}

func (IssueCommentResponse) TableName() string {
	return "issue_comments"
}

type StatusResponse struct {
	ID uint32
	Title string
	HexCode string
}

func (UserLabelResponse) TableName() string {
	return "users"
}

type UserLabelResponse struct {
	ID uint64 `json:"id"`
	Name string `json:"name"`
	ProfilePictureURL string `json:"profilePictureURL"`
}
```
Insert Association

```go
func (db *issueConnection) 		InsertDependentIssueAssociation(issue model.Issue, dependentIssue model.Issue) (*model.Issue, error) {
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
		config.DB.Find(&tempUser, user.ID)
		user.Password = tempUser.Password
	}
	config.DB.Save(&user)
	return user
}
```

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