package main

import (
	"backend/gorm/database"
	"backend/gorm/entities"
	"backend/gorm/requests"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func main() {
	db := database.InitPostgres()

	e := echo.New()
	e.POST("/product", func(c echo.Context) error {
		var req = requests.Product{}
		err := c.Bind(&req)
		if err != nil {
			return err
		}

		result := db.Create(&entities.Product{
			Name:  req.Name,
			Price: req.Price,
		})
		if result.Error != nil {
			return result.Error
		}

		return c.String(http.StatusCreated, "Product created")
	})
	e.PATCH("/product/:id", func(c echo.Context) error {
		var req = requests.Product{}
		err := c.Bind(&req)
		if err != nil {
			return err
		}

		paramId := c.Param("id")

		id, err := strconv.Atoi(paramId)
		if err != nil {
			return err
		}

		result := db.Save(&entities.Product{
			ID:    uint(id),
			Name:  req.Name,
			Price: req.Price,
		})
		if result.Error != nil {
			return result.Error
		}

		return c.String(http.StatusNoContent, "Product updated")
	})

	e.GET("/product/:id", func(c echo.Context) error {
		paramId := c.Param("id")

		id, err := strconv.Atoi(paramId)
		if err != nil {
			return err
		}

		var product entities.Product
		result := db.First(&product, id)
		if result.Error != nil {
			if result.Error == gorm.ErrRecordNotFound {
				return c.String(http.StatusNotFound, "Product not found")
			}
			return result.Error
		}

		return c.JSON(http.StatusOK, product)
	})

	e.GET("/products", func(c echo.Context) error {

		req := requests.PaginationRequest{}
		err := c.Bind(&req)
		if err != nil {
			return err
		}

		var products []entities.Product

		db.Offset((req.Page - 1) * req.Size).Limit(req.Size).Preload("User").Find(&products)

		var totalData int64
		db.Model(&entities.Product{}).Count(&totalData)

		totalPages := int(totalData) / req.Size
		if int(totalData)%req.Size != 0 {
			totalPages++
		}

		return c.JSON(http.StatusOK, requests.PaginationResponse{
			TotalData:  totalData,
			TotalPages: totalPages,
			Data:       products,
		})
	})

	e.DELETE("/product/:id", func(c echo.Context) error {
		paramId := c.Param("id")

		id, err := strconv.Atoi(paramId)
		if err != nil {
			return err
		}

		result := db.Delete(&entities.Product{}, id)
		if result.Error != nil {
			if result.Error == gorm.ErrRecordNotFound {
				return c.String(http.StatusNotFound, "Product not found")
			}
			return result.Error
		}

		return c.JSON(http.StatusNoContent, nil)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
