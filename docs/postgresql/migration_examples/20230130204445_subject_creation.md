```go
package migrations

import (
	"github.com/forfam/authentication-service/src/model"
	"gorm.io/gorm"
)

func init() {
	migrator.AddMigration(&Migration{
		Version: "20230130204445",
		Up:      mig_20230130204445_subject_creation_up,
		Down:    mig_20230130204445_subject_creation_down,
	})
}

func mig_20230130204445_subject_creation_up(db *gorm.DB) error {
	db.Migrator().CreateTable(&model.Subject{})
	return nil
}

func mig_20230130204445_subject_creation_down(db *gorm.DB) error {
	db.Migrator().DropTable(&model.Subject{})
	return nil
}
```