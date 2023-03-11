package genericrepo

import (
	"errors"
	"fmt"
	"strings"

	"github.com/forfam/authentication-service/constants"
	"github.com/forfam/authentication-service/customerror"
	"github.com/forfam/authentication-service/log"
	"github.com/forfam/authentication-service/postgres"
	"gorm.io/gorm"
)

func Take[T Entity](item *T, targetName string, logger log.Logger) error {

	result := postgres.AuthenticationDb.Take(item)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		description := fmt.Sprintf(constants.ErrorNotFoundLogMsg, targetName)
		logger.Error(fmt.Sprintf(description+constants.ReqItem, (*item).String()))
		return customerror.NewNotFoundError(description, &result.Error, nil)
	}

	if result.Error != nil {
		description := fmt.Sprintf(constants.ErrorDuringFindLogMsg, targetName)
		logger.Error(fmt.Sprintf(description+constants.ReqItem+constants.Error, (*item).String(), result.Error))
		return customerror.NewInternalServerError(description, &result.Error, nil)
	}

	return nil
}

func IsRelationNotExists[T Entity](item *T, targetNames []string, logger log.Logger) error {

	result := postgres.AuthenticationDb.Take(&item)
	if result.RowsAffected != 0 {
		description := fmt.Sprintf(constants.ConflictExistedIn, strings.Join(targetNames, " - "))
		logger.Error(description)
		return customerror.NewConflictError(description, &result.Error, nil)
	}
	return nil

}

func Create[T Entity](item *T, targetName string, logger log.Logger) error {

	result := postgres.AuthenticationDb.Create(&item)
	if result.Error != nil {
		description := fmt.Sprintf(constants.ErrorDuringCreateLogMsg, targetName)
		logger.Error(fmt.Sprintf(description+constants.Error+constants.ReqItem, result.Error, item))
		return customerror.NewInternalServerError(description, &result.Error, nil)
	}
	return nil

}

func Update[T Entity](item *T, targetName string, logger log.Logger) error {

	result := postgres.AuthenticationDb.Save(&item)
	if result.Error != nil {
		description := fmt.Sprintf(constants.ErrorDuringUpdateLogMsg, targetName)
		logger.Error(fmt.Sprintf(description+constants.Error+constants.ReqItem, result.Error, item))
		return customerror.NewInternalServerError(description, &result.Error, nil)
	}
	return nil

}

func Delete[T Entity](item *T, targetName string, logger log.Logger) error {

	result := postgres.AuthenticationDb.Delete(&item)
	if result.Error != nil {
		description := fmt.Sprintf(constants.ErrorDuringDeleteLogMsg, targetName)
		logger.Error(fmt.Sprintf(description+constants.Error+constants.Id, result.Error, (*item).String()))
		return customerror.NewInternalServerError(description, &result.Error, nil)
	}
	return nil

}

type Entity interface {
	String() string
}
