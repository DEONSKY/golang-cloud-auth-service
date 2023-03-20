package customerror

import (
	"time"

	"github.com/forfam/authentication-service/i18n"
)

type CoreErrorResponse struct {
	Err         string    `json:"error"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	HttpCode    int       `json:"httpCode"`
	I18nMessage *string   `json:"message"`
}

type ValidationErrorResponse struct {
	Err         string    `json:"error"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	CoreError   error     `json:"cause"`
	HttpCode    int       `json:"httpCode"`
	I18nMessage *string   `json:"message"`
}

func (err CoreError) MapToCoreErrorResponse() CoreErrorResponse {
	return CoreErrorResponse{
		Err:         err.Err.Error(),
		Description: err.Description,
		CreatedAt:   err.CreatedAt,
		HttpCode:    err.HttpCode,
		I18nMessage: i18n.Translate(err.I18n),
	}
}

func (err CoreError) MapToValidationErrorResponse() ValidationErrorResponse {
	return ValidationErrorResponse{
		Err:         err.Err.Error(),
		Description: err.Description,
		CoreError:   err.CauseError,
		CreatedAt:   err.CreatedAt,
		HttpCode:    err.HttpCode,
		I18nMessage: i18n.Translate(err.I18n),
	}
}
