package customerror

import (
	"fmt"
	"time"

	"github.com/forfam/authentication-service/i18n"
)

type CoreError struct {
	Err         error
	Description string
	CreatedAt   time.Time
	CoreError   error
	HttpCode    int
	I18n        *i18n.Translatable
}

func (err CoreError) Error() string {
	return fmt.Sprintf("[%s] %s : %s / Caused by: %s ", err.CreatedAt, err.Err, err.Description, err.CoreError)
}

func (err CoreError) Expose() string {
	return fmt.Sprintf("[%s] %s : %s / Caused by: %s ", err.CreatedAt, err.Err, err.Description, err.CoreError)
}

func new(Err error, description string, coreError *error, httpCode int, i18n *i18n.Translatable) *CoreError {
	var decidedError = *coreError
	if coreError == nil {
		decidedError = fmt.Errorf(description)
	}
	return &CoreError{
		Err:         Err,
		Description: description,
		CreatedAt:   time.Now(),
		CoreError:   decidedError,
		HttpCode:    httpCode,
		I18n:        i18n,
	}
}

var (
	InternalServerError = fmt.Errorf("Internal Server Error")
	NotFoundError       = fmt.Errorf("Not Found")
	BadRequestError     = fmt.Errorf("Bad Request")
	UnauthorizedError   = fmt.Errorf("Unauthorized")
	ConflictError       = fmt.Errorf("Conflict")
	ValidationError     = fmt.Errorf("Validation Error")
)

func NewInternalServerError(description string, coreError *error, i18n *i18n.Translatable) *CoreError {
	return new(InternalServerError, description, coreError, 500, i18n)
}

func NewNotFoundError(description string, coreError *error, i18n *i18n.Translatable) *CoreError {
	return new(NotFoundError, description, coreError, 404, i18n)
}

func NewBadRequestError(description string, coreError *error, i18n *i18n.Translatable) *CoreError {
	return new(BadRequestError, description, coreError, 400, i18n)
}

func NewUnauthorizedError(description string, coreError *error, i18n *i18n.Translatable) *CoreError {
	return new(UnauthorizedError, description, coreError, 401, i18n)
}

func NewConflictError(description string, coreError *error, i18n *i18n.Translatable) *CoreError {
	return new(ConflictError, description, coreError, 409, i18n)
}

func NewValidationError(description string, coreError *error, i18n *i18n.Translatable) *CoreError {
	return new(ValidationError, description, coreError, 400, i18n)
}
