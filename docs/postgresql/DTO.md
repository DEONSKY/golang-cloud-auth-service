Also we may need to make changes on the for using associations with dto

```go
type IssueKanbanResponse struct {
	Status	StatusResponse
	Issues	[]IssueResponse
}

type IssueResponse struct {
	Id		string		
	StatusId 	uint32
	Status 		StatusResponse 		`gorm:"-"`
	Reporter 	UserLabelResponse
	Assignie 	*UserLabelResponse
	ChildIssues 	[]*LeafIssueResponse 	`gorm:"foreignkey:ParentIssueId;"`
	DependentIssues []*LeafIssueResponse 	`gorm:"many2many:DependentIssues;foreignkey:Id;joinForeignKey:issueId;References:Id;joinReferences:dependentIssueId"`
}

type LeafIssueResponse struct {
	Id 		string
	Reporter 	UserLabelResponse
	AssignieId 	*string
	Assignie 	*UserLabelResponse
}

// TableName overrides the table name for smart select

func (LeafIssueResponse) TableName() string {
	return "issues"
}

type IssueCommentResponse struct {
	Context 	string
	IssueId 	string
	CreatorId 	string
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
	Id 			string
	Name 			string
	ProfilePictureURL 	string
}
```

```go
type SubjectNavTreeResponse struct {
	Id          string
	Title       string
	Description string
	ProjectId   string
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
