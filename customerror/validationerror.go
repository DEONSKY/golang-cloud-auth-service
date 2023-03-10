package customerror

import "fmt"

type ValidationErrorData struct {
	FailedField string
	Tag         string
	Value       string
}

type ValidationErrors struct {
	Errors []ValidationErrorData
}

func (errRes *ValidationErrorData) Error() string {
	return fmt.Sprintf("Failed Field : %s, Tag : %s, Value: %s", errRes.FailedField, errRes.Tag, errRes.Value)
}

func (errRes *ValidationErrors) Error() string {
	var message string
	for _, err := range errRes.Errors {
		message += err.Error()
	}
	return message
}
