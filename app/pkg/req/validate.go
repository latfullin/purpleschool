package req

import (
	"sync"

	"github.com/go-playground/validator/v10"
)

var (
	validate *validator.Validate
	only     sync.Once
)

func IsValid[T any](payload T) error {

	only.Do(func() {
		validate = validator.New()
	})

	if err := validate.Struct(payload); err != nil {
		return err
	}

	return nil
}
