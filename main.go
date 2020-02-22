package main

import (
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

func updateBook(c echo.Context) (err error) {
	b := new(models.Book)
	if err := c.Bind(b); err != nil {
		return err
	}
	id := c.Param("id")
	db.BooksDB[id].Title = b.Title
	return c.JSON(http.StatusOK, db.BooksDB[id])
}

func createBook(c echo.Context) (err error) {

	b := new(models.Book)
	if err := c.Bind(b); err != nil {
		return err
	}
	id := uuid.New().String()
	b.ID = id

	db.Library.Put(id, b)
	return c.JSON(http.StatusCreated, b)
}

func getBook(c echo.Context) (err error) {
	id := c.Param("id")
	return c.JSON(http.StatusOK, db.BooksDB[id])
}

func deleteBook(c echo.Context) (err error) {
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
