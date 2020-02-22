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
	if err := c.Bind(b); err != nil {
		return err
	}
	id := c.Param("id")
	db.BooksDB[id].Title = b.Title
	return c.JSON(http.StatusOK, db.BooksDB[id])
}

// CreateBook is for  creating new book
func CreateBook(c echo.Context) (err error) {

	b := new(models.Book)
	if err := c.Bind(b); err != nil {
		return err
	}
	id := uuid.New().String()
	b.ID = id

	db.Library.Put(id, b)
	return c.JSON(http.StatusCreated, b)
}

// GetBook is for  returning book by id
func GetBook(c echo.Context) (err error) {
	id := c.Param("id")
	return c.JSON(http.StatusOK, db.BooksDB[id])
}

// DeleteBook is for  deleting book by id
func DeleteBook(c echo.Context) (err error) {
	id := c.Param("id")
	delete(db.BooksDB, id)
	return c.NoContent(http.StatusNoContent)
}
