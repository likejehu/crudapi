package main

import (
	"io/ioutil"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/likejehu/crudapi/db"
	"github.com/likejehu/crudapi/models"
)

//----------
// Handlers
//----------

func updateBook(c echo.Context) error {
	b := new(models.Book)
	if err := c.Bind(b); err != nil {
		return err
	}
	id := c.Param("id")
	db.BooksDB[id].Title = b.Title
	return c.JSON(http.StatusOK, db.BooksDB[id])
}

func createBook(c echo.Context) error {
	id := uuid.New().String()
	b := &models.Book{
		ID: id,
	}
	var bodyBytes []byte
	if c.Request().Body != nil {
		bodyBytes, _ = ioutil.ReadAll(c.Request().Body)
	}
	if err := c.Bind(b); err != nil {
		return err
	}

	db.BooksDB[id] = b
	return c.JSON(http.StatusCreated, bodyBytes)
}

func getBook(c echo.Context) error {
	id := c.Param("id")
	return c.JSON(http.StatusOK, db.BooksDB[id])
}

func deleteBook(c echo.Context) error {
	id := c.Param("id")
	delete(db.BooksDB, id)
	return c.NoContent(http.StatusNoContent)
}

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.POST("/books", createBook)
	e.GET("/books/:id", getBook)
	e.PUT("/books/:id", updateBook)
	e.DELETE("/books/:id", deleteBook)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
