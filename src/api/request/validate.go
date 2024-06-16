package request

import (
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	rest_err "github.com/marcelofabianov/go-mvc/src/errors"
)

var Validate = validator.New()
var transl ut.Translator

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		uni := ut.New(en, en)
		transl, _ = uni.GetTranslator("en")
		en_translations.RegisterDefaultTranslations(val, transl)
	}
}

func ValidateRequest(validation_err error) *rest_err.RestErr {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationErr validator.ValidationErrors

	if errors.As(validation_err, &jsonErr) {
		println("Invalid field type")
		return rest_err.NewBadRequestError("Invalid field type")
	}

	if errors.As(validation_err, &jsonValidationErr) {
		println("Invalid request body")
		errorsCauses := []rest_err.Causes{}

		for _, e := range validation_err.(validator.ValidationErrors) {
			errorsCauses = append(errorsCauses, rest_err.Causes{
				Field:   e.Field(),
				Message: e.Translate(transl),
			})
		}

		return rest_err.NewUnprocessableEntityError("Some fields are invalid", errorsCauses)
	}

	return rest_err.NewBadRequestError("Invalid request body")
}
