package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/likejehu/crudapi/handlers"
	"gopkg.in/go-playground/validator.v10"
)

// CustomValidator struct is for storing the custom validator that will be registered to echo server
type CustomValidator struct {
	Validator *validator.Validate
}

// Validate is struct method that is called by registered validator in echo to validate
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}
func main() {
	e := echo.New()
	//create new validator
	validator := validator.New()
	//link our custom validator to echo framework
	e.Validator = &CustomValidator{validator}

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	handler := handlers.Handler{}
	e.POST("/books", handler.CreateBook)
	e.GET("/books", handler.GetLibrary)
	e.GET("/books/:id", handler.GetBook)
	e.PUT("/books/", handler.UpdateBook)
	e.DELETE("/books/:id", handler.DeleteBook)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
