package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo"
	"github.com/likejehu/crudapi/db"
	"github.com/likejehu/crudapi/models"
	"gopkg.in/go-playground/validator.v10"
)

//Bookdatabase is mine interface for any (real and mock) datastorage for books
type Bookdatabase interface {
	Delete(key string) error
	Get(key string) (*models.Book, error)
	Post(key string, book *models.Book) models.Book
	Update(key string, book *models.Book) (*models.Book, error)
	GetAll() map[string]*models.Book
}

// fieldError is for  custom validation message
type fieldError struct {
	err validator.FieldError
}

func (q fieldError) String() string {
	var sb strings.Builder

	sb.WriteString("validation failed on field '" + q.err.Field() + "'")
	sb.WriteString(", condition: " + q.err.ActualTag())

	// Print condition parameters, e.g. oneof=red blue -> { red blue }
	if q.err.Param() != "" {
		sb.WriteString(" { " + q.err.Param() + " }")
	}

	if q.err.Value() != nil && q.err.Value() != "" {
		sb.WriteString(fmt.Sprintf(", actual: %v", q.err.Value()))
	}

	return sb.String()
}

// Handler is empty struct for handlers
type Handler struct {
	Bookmap Bookdatabase
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
		for _, fieldErr := range err.(validator.ValidationErrors) {
			return c.JSON(http.StatusBadRequest, fieldError{fieldErr}.String())
		}
	}
	_, e := h.Bookmap.Update(id, b)
	if e == db.ErrorNotFound {
		return c.JSON(http.StatusNotFound, "book not found")
	}
	return c.JSON(http.StatusOK, b)
}

// CreateBook is for  creating new book
func (h *Handler) CreateBook(c echo.Context) (err error) {
	b := new(models.Book)
	if err := c.Bind(b); err != nil {
		return err
	}
	id := uuid.New().String()
	if err := c.Validate(b); err != nil {
		for _, fieldErr := range err.(validator.ValidationErrors) {
			return c.JSON(http.StatusBadRequest, fieldError{fieldErr}.String())
		}
	}
	h.Bookmap.Post(id, b)
	return c.JSON(http.StatusCreated, b)
}

// GetBook is for  returning book by id
func (h *Handler) GetBook(c echo.Context) (err error) {
	id := c.Param("id")
	book, err := h.Bookmap.Get(id)
	if err == db.ErrorNotFound {
		return c.JSON(http.StatusNotFound, "book not found")
	}
	return c.JSON(http.StatusOK, book)
}

// DeleteBook is for  deleting book by id
func (h *Handler) DeleteBook(c echo.Context) (err error) {
	id := c.Param("id")
	err = h.Bookmap.Delete(id)
	if err == db.ErrorNotFound {
		return c.JSON(http.StatusNotFound, "book not found")
	}
	return c.NoContent(http.StatusNoContent)
}

// GetLibrary is for returning all the books
func (h *Handler) GetLibrary(c echo.Context) (err error) {
	return c.JSON(http.StatusOK, h.Bookmap.GetAll())
}
