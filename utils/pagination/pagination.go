package pagination

import (
	"fmt"

	"gorm.io/gorm"
)

type PaginationOptions struct {
	Page           int    `query:"page"`
	Limit          int    `query:"limit"`
	OrderDirection string `query:"orderDirection"`
	OrderBy        string `query:"orderBy"`
	// Filters T
}

func (opt *PaginationOptions) GetLimit() int {
	if opt.Limit < 1 {
		return 10
	}

	return opt.Limit
}

func (opt *PaginationOptions) GetPage() int {
	if opt.Page < 1 {
		return 1
	}

	return opt.Page
}

func (opt *PaginationOptions) GetOffset() int {
	return (opt.GetPage() - 1) * opt.GetLimit()
}

func (opt *PaginationOptions) GetSort() string {
	direction := opt.OrderDirection
	field := opt.OrderBy

	if direction == "" {
		direction = "desc"
	}

	if field == "" {
		field = "Id"
	}

	return fmt.Sprintf("%s %s", field, direction)
}

type PaginationResult [T]struct {
	Items      []*T  `json:"items"`
	TotalCount int64 `json:"totalCount"`
	Page       int   `json:"page"`
	Limit      int   `json:"limit"`
}

func startupPagination[T any](db *gorm.DB, model T, opt *PaginationOptions) func(db *gorm.DB) *gorm.DB {
	var totalCount int64
	db.Model(model).Count(&totalCount)

	pagination.TotalCount = totalCount

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort())
	}
}

func Paginate[T any](db *gorm.DB, model T, opt *PaginationOptions) (*PaginationResult[T], error) {
	var items []*T

	pagination := &PaginationResult[T]{
		Items: items,
		Page:  opt.GetPage(),
		Limit: opt.GetLimit(),
	}

	db.Scopes(startupPagination(db, model, opt)).Find(&items)
	pagination.Items = items
	pagination.TotalCount = opt.TotalCount

	return &pagination, nil
}
