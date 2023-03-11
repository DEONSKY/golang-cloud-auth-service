package i18n

import "fmt"

type Translatable struct {
	MessageCode   string
	MessageValues *map[string]interface{}
}

// Dummy translate
func Translate(translatable *Translatable) *string {
	if translatable == nil {
		return nil
	}
	message := "(i18n not implemented yet) Message: " + translatable.MessageCode
	if translatable.MessageValues == nil {
		return &message
	}
	var values map[string]interface{}
	values = *translatable.MessageValues
	for val := range *translatable.MessageValues {
		message += fmt.Sprintf(" | Key: %v : Value: %v", val, values[val])
	}
	return &message
}

func NewTranslatable(code string, values *map[string]interface{}) *Translatable {
	return &Translatable{
		MessageCode:   code,
		MessageValues: values,
	}
}
