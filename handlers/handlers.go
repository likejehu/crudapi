package handlers

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo"
	"github.com/likejehu/crudapi/db"
	"github.com/likejehu/crudapi/models"
)

// Handler is empty struct for handlers
type Handler struct {
	db map[string]*models.Book
}

//----------
// Handlers
//----------

// UpdateBook is  for  updating books properties
func (h *Handler) UpdateBook(c echo.Context) (err error) {
	b := new(models.Book)

	if err := c.Bind(b); err != nil {
		return err
	}
	id := c.Param("id")
	if err := c.Validate(b); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, db.Library.Update(id, b))
}

// CreateBook is for  creating new book
func (h *Handler) CreateBook(c echo.Context) (err error) {

	b := new(models.Book)
	if err := c.Bind(b); err != nil {
		return err
	}
	id := uuid.New().String()
	if err := c.Validate(b); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, db.Library.Post(id, b))
}

// GetBook is for  returning book by id
func (h *Handler) GetBook(c echo.Context) (err error) {
	id := c.Param("id")
	return c.JSON(http.StatusOK, db.Library.Get(id))
}

// DeleteBook is for  deleting book by id
func (h *Handler) DeleteBook(c echo.Context) (err error) {
	id := c.Param("id")
	db.Library.Delete(id)
	return c.NoContent(http.StatusNoContent)
}

// GetLibrary is for returning all the books
func (h *Handler) GetLibrary(c echo.Context) (err error) {
	return c.JSON(http.StatusOK, db.Library.GetAll())
}
