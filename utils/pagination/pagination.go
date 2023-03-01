package pagination

import (
	"fmt"

	"gorm.io/gorm"
)

type PaginationOptions struct {
	Page           int    `query:"page" validate:"required,number"`
	Limit          int    `query:"limit"`
	OrderDirection string `query:"orderDirection" validate:"omitempty,oneof=asc desc"`
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

type PaginationResult[T any] struct {
	Items      []*T  `json:"items"`
	TotalCount int64 `json:"totalCount"`
	Page       int   `json:"page"`
	Limit      int   `json:"limit"`
}

func startupPagination[DTO any](db *gorm.DB, opt *PaginationOptions, pagination *PaginationResult[DTO]) func(db *gorm.DB) *gorm.DB {
	var totalCount int64
	db.Count(&totalCount)

	pagination.TotalCount = totalCount

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(opt.GetOffset()).Limit(opt.GetLimit()).Order(opt.GetSort())
	}
}

func Paginate[DTO any](db *gorm.DB, opt *PaginationOptions) (*PaginationResult[DTO], error) {
	var items []*DTO

	pagination := &PaginationResult[DTO]{
		Items: items,
		Page:  opt.GetPage(),
		Limit: opt.GetLimit(),
	}

	db.Scopes(startupPagination(db, opt, pagination)).Find(&items)

	pagination.Items = items

	return pagination, nil
}
