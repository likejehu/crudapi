package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/likejehu/crudapi/models"
	"github.com/stretchr/testify/assert"
	"gopkg.in/go-playground/validator.v10"
)

/*
var mockBook = models.Book{
	Title:       "Super Kniga",
	Author:      "Igor",
	Publisher:   "Superizdatel",
	PublishDate: "2020-02-02",
	Rating:      2,
	Status:      "CheckedOut",
}

var bookBytes, _ = json.Marshal(mockBook)
var bookJSON = string(bookBytes)
*/
var mockDB = map[string]*models.Book{
	"80369cc7-41af-48a0-be9e-a71377bcb337": &models.Book{"SUper kniga", "Igor", "Superizdatel", "2020-02-02", 3, "CheckedIn"},
}

var bookJSON = `{"title":"SUper kniga","author":"Igor","publisher":"Superizdatel","date":"2020-02-02","rating":3,"status":"CheckedIn"}
`

type customValidator struct {
	Validator *validator.Validate
}

func (cv *customValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
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
