# Migration Commands

## Example empty migration creation
```sh
go run ./cmd psql:migration-create -n user_create
```

## Migrate specific migration
```sh
go run ./cmd psql:migrate -n 20230210014853_user_create  
```

## Migrate all migrations with order
```sh
go run ./cmd psql:migrate 
```

## Undo specific migration
```sh
go run ./cmd psql:migrate-undo -n 20230210014853_user_create
```

# Migration Examples

In this section we will create user-subject many to many relational tables with migrations.

Fist we will create user table:

## 20230210153114_user_create.go Migration

In this migration we are creating our first entity(user) in database

We defined an user struct with uniqeu name for this migration. This structs should be unique inside package. For this reason we
are adding version date after struct name.  But gorm naming tables with struct name automaticly. For this reason we are defining our
table name with TableName() function.

```go
package migrations

import (
	"gorm.io/gorm"
)

type users_20230210153114 struct {
	Id   string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name string `gorm:"type:varchar"`
}

func (table *users_20230210153114) TableName() string {
	return "users"
}

func mig_20230210153114_user_create_up(transaction *gorm.DB) error {
	if err := transaction.Migrator().CreateTable(&users_20230210153114{}); err != nil {
		return err
	}
	return nil
}

func mig_20230210153114_user_create_down(transaction *gorm.DB) error {
	if err := transaction.Migrator().DropTable(&users_20230210153114{}); err != nil {
		return err
	}
	return nil
}

func init() {
	Migrations = append(Migrations, PostgresMigration{
		Name: "20230210153114_user_create",
		Up:   mig_20230210153114_user_create_up,
		Down: mig_20230210153114_user_create_down,
	})
}

```

After that we are creating subject migration

## 20230210155706_subject_create.go Migration

We are creating this migration like user migration

```go
package migrations

import (
	"gorm.io/gorm"
)

type subject_20230210155706 struct {
	Id   string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name string `gorm:"type:varchar"`
}

func (table *subject_20230210155706) TableName() string {
	return "subjects"
}

func mig_20230210155706_subject_create_up(transaction *gorm.DB) error {
	if err := transaction.Migrator().CreateTable(&subject_20230210155706{}); err != nil {
		return err
	}

	return nil
}

func mig_20230210155706_subject_create_down(transaction *gorm.DB) error {
	if err := transaction.Migrator().DropTable(&subject_20230210155706{}); err != nil {
		return err
	}
	return nil
}

func init() {
	Migrations = append(Migrations, PostgresMigration{
		Name: "20230210155706_subject_create",
		Up:   mig_20230210155706_subject_create_up,
		Down: mig_20230210155706_subject_create_down,
	})
}
```

## 20230210160630_subject_user_create.go Migration

Now we are creating migration for association table. But we need to create empty structs for precreated table.  

```go
package migrations

import (
	"gorm.io/gorm"
)

type subject_user_20230210160630 struct {
	Id     string               `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserId string               `gorm:"type:uuid"`
	User   users_20230210153114 `gorm:"foreignkey:UserId;"`
}

func (table *subject_user_20230210160630) TableName() string {
	return "subject_users"
}

func mig_20230210160630_subject_user_create_up(transaction *gorm.DB) error {
	if err := transaction.Migrator().CreateTable(&subject_user_20230210160630{}); err != nil {
		return err
	}
	return nil
}

func mig_20230210160630_subject_user_create_down(transaction *gorm.DB) error {
	if err := transaction.Migrator().DropTable(&subject_user_20230210160630{}); err != nil {
		return err
	}

	return nil
}

func init() {
	Migrations = append(Migrations, PostgresMigration{
		Name: "20230210160630_subject_user_create",
		Up:   mig_20230210160630_subject_user_create_up,
		Down: mig_20230210160630_subject_user_create_down,
	})
}
```

## 20230210222930_subject_user_add_subject_fk.go Migration

If we want to create add some fields and foreign keys after cretion we should define structs with requeired fields. After that we can use
this structs for creating migrations

```go
package migrations

import (
	"gorm.io/gorm"
)

//Note: redefination of fields is required

type subject_user_20230210222930 struct {
	Id        string                 `gorm:"type:uuid"`
	SubjectId string                 `gorm:"type:uuid"`
	Subject   subject_20230210222930 `gorm:"foreignkey:SubjectId;"`
}

func (table *subject_user_20230210222930) TableName() string {
	return "subject_users"
}

type subject_20230210222930 struct {
	Id string `gorm:"type:uuid"`
}

func (table *subject_20230210222930) TableName() string {
	return "subjects"
}

func mig_20230210222930_subject_user_add_subject_fk_up(transaction *gorm.DB) error {
	if err := transaction.Migrator().AddColumn(&subject_user_20230210222930{}, "SubjectId"); err != nil {
		return err
	}
	if err := transaction.Migrator().CreateConstraint(&subject_user_20230210222930{}, "fk_subject_users_subject"); err != nil {
		return err
	}
	return nil
}

func mig_20230210222930_subject_user_add_subject_fk_down(transaction *gorm.DB) error {

	if err := transaction.Migrator().DropConstraint(&subject_user_20230210222930{}, "fk_subject_users_subject"); err != nil {
		return err
	}
	if err := transaction.Migrator().DropColumn(&subject_user_20230210222930{}, "SubjectID"); err != nil {
		return err
	}

	return nil
}

func init() {
	Migrations = append(Migrations, PostgresMigration{
		Name: "20230210222930_subject_user_add_subject_fk",
		Up:   mig_20230210222930_subject_user_add_subject_fk_up,
		Down: mig_20230210222930_subject_user_add_subject_fk_down,
	})
}
```
