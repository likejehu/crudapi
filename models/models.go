package models

import (
	"time"

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

// Book  is mine book representation
type Book struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	Publisher   string    `json:"publisher"`
	PublishDate time.Time `json:"date"`
	Rating      uint8     `json:"rating"`
	Status      string    `json:"status"`
}

// Books  is mine books storage representation
var Books = map[int]*Book{}

func main() {

}
