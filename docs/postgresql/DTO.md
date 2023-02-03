
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

```go
type SubjectNavTreeResponse struct {
	ID          uint64 `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ProjectID   uint64 `json:"project_id"`
}
```

```go
type IssueGetQuery struct {
	SubjectID      *uint64 `query:"subjectID"`
	ProjectID      *uint64 `query:"projectID"`
	ReporterID     *uint64 `query:"reporterID"`
	AssignieID     *uint64 `query:"assignieID"`
	Status         *uint8  `query:"status"`
	ParentIssueID  *uint64 `query:"parentIssueID"`
	GetOnlyOrphans *bool   `query:"getOnlyOrphans"`
}
```