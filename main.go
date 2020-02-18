package main

import (
	"net/http"
	"strconv"

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
	id, _ := strconv.Atoi(c.Param("id"))
	models.Books[id].Title = b.Title
	return c.JSON(http.StatusOK, models.Books[id])
}

func createBook(c echo.Context) error {
	b := &models.Book{
		ID: uuid.New(),
	}
	if err := c.Bind(b); err != nil {
		return err
	}

	db.BooksDB[b.ID] = b
	return c.JSON(http.StatusCreated, b)
}

func getBook(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, models.Books[id])
}

func deleteBook(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	delete(models.Books, id)
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
