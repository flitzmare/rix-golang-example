package routes

import (
	"backend/clean-arch/deliveries/http/handlers"
	"backend/clean-arch/domain/repositories/postgres"
	"backend/clean-arch/domain/usecases"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"gorm.io/gorm"
)

type CleanArch struct {
	productRepo    postgres.IProductRepository
	productUsecase usecases.IProductUsecase
	productHandler handlers.IProductHandler
}

func NewCleanArch(db *gorm.DB) CleanArch {
	cleanArch := CleanArch{}

	cleanArch.productRepo = postgres.NewProductRepository(db)
	cleanArch.productUsecase = usecases.NewProductUsecase(cleanArch.productRepo)
	cleanArch.productHandler = handlers.NewProductHandler(cleanArch.productUsecase)

	return cleanArch
}

func SetupRouter(e *echo.Echo, c CleanArch) {

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.GET("/products", c.productHandler.GetProducts)
	e.POST("/product", c.productHandler.CreateProduct)
}
