package validator

import (
	"fmt"
	"reflect"
	"strings"
)

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

type ValidationErrors []ValidationError

func (e ValidationErrors) Error() string {
	msgs := make([]string, len(e))
	for i, ve := range e {
		msgs[i] = ve.Error()
	}
	return strings.Join(msgs, "; ")
}

// Validate performs basic required-field validation using struct tags.
// For production use, replace with github.com/go-playground/validator.
func Validate(v any) error {
	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	var errs ValidationErrors
	t := val.Type()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := val.Field(i)
		tag := field.Tag.Get("validate")

		if strings.Contains(tag, "required") {
			if value.IsZero() {
				jsonTag := field.Tag.Get("json")
				name := strings.Split(jsonTag, ",")[0]
				if name == "" {
					name = field.Name
				}
				errs = append(errs, ValidationError{Field: name, Message: "is required"})
			}
		}
	}

	if len(errs) > 0 {
		return errs
	}
	return nil
}
