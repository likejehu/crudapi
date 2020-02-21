package handlers

import (
	"github.com/go-playground/validator"
)

// CustomValidator is mine validator
type CustomValidator struct {
	validator *validator.Validate
}

// Validate is for  validating
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {

}
