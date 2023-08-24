package helper

import (
	"errors"
	"github.com/go-playground/validator/v10"
)

func Validate(err error) any {
	out := make(map[string]interface{})

	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		for _, v := range ve {
			out[v.Field()] = v.Tag()
		}
		return out
	}

	return err.Error()
}
