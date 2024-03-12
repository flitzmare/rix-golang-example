package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Use(middleware)
	e.GET("/product", func(c echo.Context) error {
		fmt.Println("Get product")
		return c.String(http.StatusOK, "Get product")
	})
	e.Logger.Fatal(e.Start(":8080"))
}

func middleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)
		fmt.Println("Saya dipanggil sesudah hit endpoint")
		return err
	}
}
