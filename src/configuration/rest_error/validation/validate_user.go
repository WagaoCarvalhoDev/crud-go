package validation

import (
	resterror "crudgo/src/configuration/rest_error"
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
)

var (
	validate = validator.New()
	transl   ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		unt := ut.New(en, en)
		transl, _ = unt.GetTranslator("en")
		en_translation.RegisterDefaultTranslations(val, transl)
	}
}

func ValidateUserError(validation_error error) *resterror.RestError {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validation_error, &jsonErr) {
		return resterror.NewBadRequestError("Invalid field type")
	} else if errors.As(validation_error, &jsonValidationError) {
		errorCauses := []resterror.Causes{}
		for _, e := range validation_error.(validator.ValidationErrors) {
			cause := resterror.Causes{
				Message: e.Translate(transl),
				Field:   e.Field(),
			}
			errorCauses = append(errorCauses, cause)
		}
		return resterror.NewBadRequestValidationError("Some fields are invalid", errorCauses)
	} else {
		return resterror.NewBadRequestError("Error trying to fields")
	}
}
