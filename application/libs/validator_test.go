package libs

import (
	"encoding/json"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func convertValidationErrorsToSlice(errs validator.ValidationErrors) []error {
	var errors []error
	for _, err := range errs {
		errors = append(errors, err)
	}

	return errors
}

func TestParseFieldError(t *testing.T) {
	cv := NewCustomValidator()
	validate := validator.New()
	err := validate.Var("", "required")
	validationErrors := err.(validator.ValidationErrors)
	parsedErrors := cv.ParseError(convertValidationErrorsToSlice(validationErrors)...)

	expected := []ErrorMessage{
		{
			Field:   "required",
			Message: "Required field is required",
		},
	}

	assert.Equal(t, expected, parsedErrors)
}

func TestParseMarshallingError(t *testing.T) {
	cv := NewCustomValidator()
	var invalidJSON = []byte(`{"name":123}`)
	var data map[string]string
	err := json.Unmarshal(invalidJSON, &data)
	parsedErrors := cv.ParseError(err)

	expected := []ErrorMessage{
		{
			Field:   "name",
			Message: "The field Name must be a string",
		},
	}

	assert.Equal(t, expected, parsedErrors)
}

func TestParseErrorWithMultipleErrors(t *testing.T) {
	cv := NewCustomValidator()
	validate := validator.New()
	type User struct {
		Email string `validate:"required,email"`
	}
	user := User{}
	err := validate.Struct(user)
	validationErrors := err.(validator.ValidationErrors)
	parsedErrors := cv.ParseError(convertValidationErrorsToSlice(validationErrors)...)

	expected := []ErrorMessage{
		{
			Field:   "email",
			Message: "Email field is required",
		},
		{
			Field:   "email",
			Message: "Email field is not valid",
		},
	}

	assert.Equal(t, expected, parsedErrors)
}

func TestParseErrorWithDefaultError(t *testing.T) {
	cv := NewCustomValidator()
	err := &json.UnmarshalTypeError{
		Value: "number",
		Type:  nil,
		Field: "age",
	}
	parsedErrors := cv.ParseError(err)

	expected := []ErrorMessage{
		{
			Field:   "age",
			Message: "The field Age must be a ",
		},
	}

	assert.Equal(t, expected, parsedErrors)
}
