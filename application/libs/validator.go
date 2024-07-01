package libs

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/polluxdev/trx-core-svc/common/helper"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type CustomValidator struct{}

type ErrorMessage struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func NewCustomValidator() *CustomValidator {
	return &CustomValidator{}
}

func (cv *CustomValidator) ParseError(errs ...error) []ErrorMessage {
	var out []ErrorMessage
	for _, err := range errs {
		switch typedError := any(err).(type) {
		case validator.ValidationErrors:
			for _, e := range typedError {
				out = append(out, cv.parseFieldError(e))
			}
		case *json.UnmarshalTypeError:
			out = append(out, cv.parseMarshallingError(*typedError))
		default:
			out = append(out, ErrorMessage{Field: err.Error(), Message: err.Error()})
		}
	}

	return out
}

func (cv *CustomValidator) parseFieldError(e validator.FieldError) ErrorMessage {
	field := helper.ToSnakeCase(e.Field())
	errorMessage := new(ErrorMessage)
	errorMessage.Field = field
	tag := strings.Split(e.Tag(), "|")[0]
	words := strings.Split(field, "_")
	caser := cases.Title(language.English)
	for i, word := range words {
		words[i] = caser.String(word)
	}
	field = strings.Join(words, " ")

	switch tag {
	case "required":
		errorMessage.Message = fmt.Sprintf("%s field is required", field)
		return *errorMessage
	case "email":
		errorMessage.Message = fmt.Sprintf("%s field is not valid", field)
		return *errorMessage
	case "min":
		errorMessage.Message = fmt.Sprintf("%s field must be at least %s characters", field, e.Param())
		return *errorMessage
	case "max":
		errorMessage.Message = fmt.Sprintf("%s field must be at most %s characters", field, e.Param())
		return *errorMessage
	default:
		english := en.New()
		translator := ut.New(english, english)
		if translatorInstance, found := translator.GetTranslator("en"); found {
			errorMessage.Message = e.Translate(translatorInstance)
		} else {
			errorMessage.Message = fmt.Errorf("%v", e).Error()
		}
		return *errorMessage
	}
}

func (cv *CustomValidator) parseMarshallingError(e json.UnmarshalTypeError) ErrorMessage {
	words := strings.Split(e.Field, "_")
	caser := cases.Title(language.English)
	for i, word := range words {
		words[i] = caser.String(word)
	}
	field := strings.Join(words, " ")

	return ErrorMessage{
		Field:   e.Field,
		Message: fmt.Sprintf("The field %s must be a %s", field, e.Type.String()),
	}
}
