package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/likejehu/crudapi/handlers"
	"gopkg.in/go-playground/validator.v10"
)

// CustomValidator struct is for storing the custom validator that will be registered to echo server
type CustomValidator struct {
	validator *validator.Validate
}

// Validate is for validation
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
func main() {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

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
