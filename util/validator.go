package util

import (
	"github.com/go-playground/validator/v10"
)

// Error Map
type ErrMap map[string]string

type CustomValidator struct {
	Validator *validator.Validate
}

// Validate ...
func (cv *CustomValidator) Validate(i interface{}) error {
	err := cv.Validator.Struct(i)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return err
		}
		// errmsg := ErrMap{}
		// for _, err := range err.(validator.ValidationErrors) {
		// 	key := strings.ToLower(err.Field())
		// 	errmsg[key] = fmt.Sprintf("There is an error with '%s' field", err.Field())

		// }
		// return errmsg
	}
	return nil
}
