package handlers

import (
	"net/http"

	"github.com/go-playground/validator"

	"github.com/google/uuid"
	"github.com/labstack/echo"
	"github.com/likejehu/crudapi/db"
	"github.com/likejehu/crudapi/models"
)

// CustomValidator is mine validator
type CustomValidator struct {
	validator *validator.Validate
}

// Validate is for  validating
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

//----------
// Handlers
//----------

// UpdateBook is  for  updating books properties
func UpdateBook(c echo.Context) (err error) {
	b := new(models.Book)
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	if err := c.Bind(b); err != nil {
		return err
	}
	id := c.Param("id")
	if err = c.Validate(b); err != nil {
		return
	}

	return c.JSON(http.StatusOK, db.Library.Update(id, b))
}

// CreateBook is for  creating new book
func CreateBook(c echo.Context) (err error) {

	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	b := new(models.Book)
	if err := c.Bind(b); err != nil {
		return err
	}
	id := uuid.New().String()
	b.ID = id
	if err = c.Validate(b); err != nil {
		return
	}

	db.Library.Put(id, b)
	return c.JSON(http.StatusCreated, b)
}

// GetBook is for  returning book by id
func GetBook(c echo.Context) (err error) {
	id := c.Param("id")
	return c.JSON(http.StatusOK, db.Library.Get(id))
}

// DeleteBook is for  deleting book by id
func DeleteBook(c echo.Context) (err error) {
	id := c.Param("id")
	db.Library.Delete(id)
	return c.NoContent(http.StatusNoContent)
}
