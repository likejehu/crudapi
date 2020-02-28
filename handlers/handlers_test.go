package handlers

import (
	"testing"

	"github.com/likejehu/crudapi/models"
	_ "github.com/stretchr/testify/assert"
)

var (
	mockDB = map[string]*models.Book{
		"80369cc7-41af-48a0-be9e-a71377bcb337": &models.Book{"80369cc7-41af-48a0-be9e-a71377bcb337", "SUper kniga", "Igor", "Superizdatel", "2020-02-02", 3, "CheckedIn"},
	}

	bookJSON = `{"id": "80369cc7-41af-48a0-be9e-a71377bcb337", "title": "SUper kniga", "author": "Igor", "publisher": "Superizdatel", "date": "2020-02-02", "rating": 3, "status": "CheckedIn"}`
)

func TestCreateBook(t *testing.T) {
	//setup
	/*
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(bookJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
	*/
}
