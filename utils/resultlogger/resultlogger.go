package resultlogger

import (
	"fmt"

	"gorm.io/gorm"

	"github.com/forfam/authentication-service/log"
)

//type LogMessage string
//Defined as string instead of struct due to performance reasons

const (
	ErrorDuringFindLogMsg   string = `Something went wrong during find "%s"! Id: %s - Error: %s`
	ErrorNotFoundLogMsg     string = `%s not found! Id: %s - Error: %s`
	NotFoundForDeleteLogMsg string = `%s not found for delete! Id: %s`
	NotFoundForUpdateLogMsg string = `%s not found for update! Id: %s`
	ErrorDuringDeleteLogMsg string = `Something went wrong during delete "%s"! Id: %s - Error: %s`
	NotDeletedLogMsg        string = `%s not deleted! Id: %s`
	ErrorDuringUpdateLogMsg string = `Something went wrong during update "%s"! Id: %s - Error: %s`
	NotUpdatedLogMsg        string = `%s not updated! Id: %s - Data: %s`
	ErrorDuringCreateLogMsg string = `Something went wrong during creation of "%s"! Id: %s - Error: %s`
)

func LogGormResult(result *gorm.DB, logger *log.Logger, id string, errlogMsg string, rowsAffectedLogMsg string, entName string) (*gorm.DB, error) {
	if result.Error != nil {
		logger.Error(fmt.Sprintf(errlogMsg, entName, id, result.Error))
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		logger.Warning(fmt.Sprintf(rowsAffectedLogMsg, entName, id))
		return nil, nil
	}
	return result, nil
}
