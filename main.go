package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/likejehu/crudapi/handlers"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.POST("/books", handlers.CreateBook)
	e.GET("/books/:id", handlers.GetBook)
	e.PUT("/books/:id", handlers.UpdateBook)
	e.DELETE("/books/:id", handlers.DeleteBook)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
