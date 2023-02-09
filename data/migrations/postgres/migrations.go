package migrations

import (
	"sort"

	"gorm.io/gorm"
)

type PostgresMigration struct {
	Name string
	Up   func(*gorm.DB) error
	Down func(*gorm.DB) error
}

var Migrations []PostgresMigration

type ByName []PostgresMigration

func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].Name < a[j].Name }

func Sort() {
	sort.Sort(ByName(Migrations))
}
