package validator

import (
	v "github.com/go-playground/validator/v10"
	"reflect"
	"strings"
	"sync"
)

var (
	validatorInstance *v.Validate

	once = &sync.Once{}
)

func Get() *v.Validate {
	return validatorInstance
}

func Load() (err error) {
	once.Do(func() {
		validatorInstance = v.New()

		validatorInstance.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

			if name == "-" {
				return ""
			}

			return name
		})

	})

	return err
}
