package utilities

import (
	"fmt"

	"github.com/asaskevich/govalidator"
)

// Validator function Represent validate struct data
func Validator(data interface{}) error {
	_, err := govalidator.ValidateStruct(data)

	if err != nil {
		return fmt.Errorf(err.Error())
	}

	return nil
}
