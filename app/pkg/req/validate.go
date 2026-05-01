package req

import "github.com/go-playground/validator/v10"

var validate = validator.New()

func IsValid[T any](payload T) error {

	if err := validate.Struct(payload); err != nil {
		return err
	}

	return nil
}
