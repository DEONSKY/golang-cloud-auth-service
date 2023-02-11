Also we may need to make changes on the for using associations with dto

```go
type IssueKanbanResponse struct {
	Status	StatusResponse	`json:"status"`
	Issues	[]IssueResponse	`json:"issues"`
}

type IssueResponse struct {
	Id		string 			`json:"id"`
	StatusId 	uint32 			`json:"statusId"`
	Status 		StatusResponse 		`gorm:"-" json:"status"`
	Reporter 	UserLabelResponse 	`json:"reporter"`
	Assignie 	*UserLabelResponse 	`json:"assignie"`
	ChildIssues 	[]*LeafIssueResponse 	`gorm:"foreignkey:ParentIssueId;" json:"issues"`
	DependentIssues []*LeafIssueResponse 	`gorm:"many2many:DependentIssues;foreignkey:Id;joinForeignKey:issueId;References:Id;joinReferences:dependentIssueId" json:"dependentIssues"`
}

type LeafIssueResponse struct {
	Id 		string 			`json:"id"`
	Reporter 	UserLabelResponse 	`json:"reporter"`
	AssignieId 	*string 		`json:"assignieId"`
	Assignie 	*UserLabelResponse 	`json:"assignie"`
}

// TableName overrides the table name for smart select

func (LeafIssueResponse) TableName() string {
	return "issues"
}

type IssueCommentResponse struct {
	Context 	string	`json:"context"`
	IssueId 	string 	`json:"issueId"`
	CreatorId 	string 	`json:"-"`
}

func (IssueCommentResponse) TableName() string {
	return "issue_comments"
}

type StatusResponse struct {
	Id 	string
	Title 	string
	HexCode string
}

func (UserLabelResponse) TableName() string {
	return "users"
}

type UserLabelResponse struct {
	Id 			string 	`json:"id"`
	Name 			string 	`json:"name"`
	ProfilePictureURL 	string 	`json:"profilePictureURL"`
}
```

```go
type SubjectNavTreeResponse struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ProjectId   string `json:"project_id"`
}
```

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
