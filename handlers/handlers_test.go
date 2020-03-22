package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/likejehu/crudapi/handlers/mocks"
	"github.com/likejehu/crudapi/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gopkg.in/go-playground/validator.v10"
)

var bookJSON = `{"title":"SUper kniga","author":"Igor","publisher":"Superizdatel","date":"2020-02-02","rating":3,"status":"CheckedIn"}
`
var badbookJSON = `{"title":"","author":"Igor","publisher":"Superizdatel","date":"2020-02-02","rating":3,"status":"CheckedIn"}
`
var testBook = &models.Book{Title: "SUper kniga", Author: "Igor", Publisher: "Superizdatel", PublishDate: "2020-02-02", Rating: 0x3, Status: "CheckedIn"}
var testBadBook = &models.Book{Title: "", Author: "Igor", Publisher: "Superizdatel", PublishDate: "2020-02-02", Rating: 0x3, Status: "CheckedIn"}

type customValidator struct {
	Validator *validator.Validate
}

func (cv *customValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}

func TestCreateBook(t *testing.T) {

	t.Run("succes case", func(t *testing.T) {
		//setup
		e := echo.New()
		validator := validator.New()
		e.Validator = &customValidator{validator}
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(bookJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		mockBookdatabase := &mocks.Bookdatabase{}
		h := &Handler{mockBookdatabase}
		mockBookdatabase.On("Post", mock.Anything, testBook).Return(*testBook)
		h.CreateBook(c)
		mockBookdatabase.AssertExpectations(t)
		// Assertions
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, bookJSON, rec.Body.String())
	})

	t.Run("validation error", func(t *testing.T) {
		//setup
		e := echo.New()
		validator := validator.New()
		e.Validator = &customValidator{validator}
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(badbookJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		mockBookdatabase := &mocks.Bookdatabase{}
		h := &Handler{mockBookdatabase}
		h.CreateBook(c)
		// Assertions
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})
}
func TestGetBookwithMockery(t *testing.T) {
	//setup
	e := echo.New()
	e.MaxParam(5)
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	mockBookdatabase := &mocks.Bookdatabase{}
	h := &Handler{mockBookdatabase}
	mockBookdatabase.On("Get", mock.Anything).Return(testBook)
	h.GetBook(c)
	mockBookdatabase.AssertExpectations(t)
	// Assertions
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, bookJSON, rec.Body.String())
}

func TestDeleteBookwithMockery(t *testing.T) {
	//setup
	e := echo.New()
	e.MaxParam(5)
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	mockBookdatabase := &mocks.Bookdatabase{}
	h := &Handler{mockBookdatabase}
	c.SetPath("/books/:id")
	c.SetParamNames("id")
	c.SetParamValues("80369cc7-41af-48a0-be9e-a71377bcb337")
	exp := ""
	mockBookdatabase.On("Delete", mock.Anything).Return(exp)
	h.DeleteBook(c)
	mockBookdatabase.AssertExpectations(t)
	// Assertions
	assert.Equal(t, http.StatusNoContent, rec.Code)
	assert.Equal(t, exp, rec.Body.String())

}

func TestGetLibrarywithMockery(t *testing.T) {
	//setup
	e := echo.New()
	e.MaxParam(5)
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	mockBookdatabase := &mocks.Bookdatabase{}
	h := &Handler{mockBookdatabase}
	mockBookdatabase.On("GetAll").Return(map[string]*models.Book{})
	h.GetLibrary(c)
	mockBookdatabase.AssertExpectations(t)
	// Assertions
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestUpdateBookwithMockery(t *testing.T) {
	//setup
	e := echo.New()
	e.MaxParam(5)
	validator := validator.New()
	e.Validator = &customValidator{validator}
	req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(bookJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	mockBookdatabase := &mocks.Bookdatabase{}
	h := &Handler{mockBookdatabase}
	mockBookdatabase.On("Update", mock.Anything, testBook).Return(testBook)
	h.UpdateBook(c)
	mockBookdatabase.AssertExpectations(t)
	// Assertions
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, bookJSON, rec.Body.String())
}
