package main

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type User struct {
	Name  string `query:"name" json:"name" validate:"required"`
	Email string `query:"email" json:"email" validate:"required,email"`
	Umur int `json:"umur" validate:"lt=80"`
}

func main() {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	e.GET("/product", func(c echo.Context) error {
		return c.String(http.StatusOK, "Get product")
	})
	e.POST("/product", func(c echo.Context) error {
		return c.String(http.StatusOK, "Product created")
	})
	e.PATCH("/product", func(c echo.Context) error {
		return c.String(http.StatusOK, "Product updated")
	})
	e.DELETE("/product", func(c echo.Context) error {
		return c.String(http.StatusOK, "Product deleted")
	})
	e.GET("/user/:name", func(c echo.Context) error {
		name := c.Param("name")
		data := fmt.Sprintf("Halo %s", name)
		return c.String(http.StatusOK, data)
	})
	e.POST("/register", func(c echo.Context) error {
		var request User
		if err := c.Bind(&request); err != nil {
			return err
		}
		if err := c.Validate(&request); err != nil {
			return err
		}
		return c.JSON(http.StatusOK, request)
	})
	e.Logger.Fatal(e.Start(":8080"))
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}
