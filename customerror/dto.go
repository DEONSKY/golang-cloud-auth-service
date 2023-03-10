package customerror

import "time"

type CoreErrorResponse struct {
	Err         string    `json:"error"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	HttpCode    int       `json:"httpCode"`
}

type ValidationErrorResponse struct {
	Err         string    `json:"error"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	CoreError   error     `json:"coreError"`
	HttpCode    int       `json:"httpCode"`
}

func (err CoreError) MapToCoreErrorResponse() CoreErrorResponse {
	return CoreErrorResponse{
		Err:         err.Err.Error(),
		Description: err.Description,
		CreatedAt:   err.CreatedAt,
		HttpCode:    err.HttpCode,
	}
}

func (err CoreError) MapToValidationErrorResponse() ValidationErrorResponse {
	return ValidationErrorResponse{
		Err:         err.Err.Error(),
		Description: err.Description,
		CoreError:   err.CoreError,
		CreatedAt:   err.CreatedAt,
		HttpCode:    err.HttpCode,
	}
}
