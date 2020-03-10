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

type mockingDB struct {
	B map[string]*models.Book
}

// Delete is delete func
func (d *mockingDB) Delete(key string) {
	delete(d.B, key)
	return
}

// Get is read func
func (d *mockingDB) Get(key string) *models.Book {
	book := d.B[key]
	return book
}

// Post is create func
func (d *mockingDB) Post(key string, book *models.Book) models.Book {
	d.B[key] = book
	return *d.B[key]
}

// Update is update func
func (d *mockingDB) Update(key string, book *models.Book) (v *models.Book) {
	d.B[key] = book
	v = d.B[key]
	return
}

// GetAll is for returning all the books
func (d *mockingDB) GetAll() map[string]*models.Book {
	books := d.B
	return books
}
func newmockDB() *mockingDB {
	var m mockingDB
	m.B = make(map[string]*models.Book)
	return &m
}

var mockDB = newmockDB()
var bookJSON = `{"title":"SUper kniga","author":"Igor","publisher":"Superizdatel","date":"2020-02-02","rating":3,"status":"CheckedIn"}
`

type customValidator struct {
	Validator *validator.Validate
}

func (cv *customValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}

func TestCreateBookwithMockery(t *testing.T) {
	mockBookdatabase := &mocks.Bookdatabase{}
	mockBookdatabase.On("Post", mock.Anything, &models.Book{Title: "SUper kniga", Author: "Igor", Publisher: "Superizdatel", PublishDate: "2020-02-02", Rating: 0x3, Status: "CheckedIn"}).Return(models.Book{Title: "SUper kniga", Author: "Igor", Publisher: "Superizdatel", PublishDate: "2020-02-02", Rating: 0x3, Status: "CheckedIn"})
	h := &Handler{mockBookdatabase}
	e := echo.New()
	validator := validator.New()
	e.Validator = &customValidator{validator}
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(bookJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h.CreateBook(c)
	mockBookdatabase.AssertExpectations(t)
	// Assertions
	assert.Equal(t, http.StatusCreated, rec.Code)
	assert.Equal(t, bookJSON, rec.Body.String())
}
func TestCreateBook(t *testing.T) {
	//setup
	e := echo.New()
	validator := validator.New()
	e.Validator = &customValidator{validator}
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(bookJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := &Handler{mockDB}

	// Assertions
	if assert.NoError(t, h.CreateBook(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, bookJSON, rec.Body.String())
	}

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
	mockBookdatabase.On("Get", mock.Anything).Return(&models.Book{Title: "SUper kniga", Author: "Igor", Publisher: "Superizdatel", PublishDate: "2020-02-02", Rating: 0x3, Status: "CheckedIn"})
	h.GetBook(c)
	mockBookdatabase.AssertExpectations(t)
	// Assertions
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, bookJSON, rec.Body.String())
}
func TestGetBook(t *testing.T) {
	// Setup
	e := echo.New()
	e.MaxParam(5)
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/books/:id")
	c.SetParamNames("id")
	c.SetParamValues("80369cc7-41af-48a0-be9e-a71377bcb337")
	mockDB.B = map[string]*models.Book{
		"80369cc7-41af-48a0-be9e-a71377bcb337": &models.Book{"SUper kniga", "Igor", "Superizdatel", "2020-02-02", 3, "CheckedIn"},
	}
	h := &Handler{mockDB}

	// Assertions
	if assert.NoError(t, h.GetBook(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, bookJSON, rec.Body.String())
	}
}

func TestDeleteBook(t *testing.T) {
	// Setup
	e := echo.New()
	e.MaxParam(5)
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/books/:id")
	c.SetParamNames("id")
	c.SetParamValues("80369cc7-41af-48a0-be9e-a71377bcb337")
	mockDB.B = map[string]*models.Book{
		"80369cc7-41af-48a0-be9e-a71377bcb337": &models.Book{"SUper kniga", "Igor", "Superizdatel", "2020-02-02", 3, "CheckedIn"},
	}
	h := &Handler{mockDB}
	exp := ""
	// Assertions
	if assert.NoError(t, h.DeleteBook(c)) {
		assert.Equal(t, http.StatusNoContent, rec.Code)
		assert.Equal(t, exp, rec.Body.String())
	}
}
