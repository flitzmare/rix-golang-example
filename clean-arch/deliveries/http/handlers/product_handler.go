package handlers

import (
	"backend/clean-arch/deliveries/requests"
	"backend/clean-arch/domain/usecases"
	"net/http"

	"github.com/labstack/echo/v4"
)

type IProductHandler interface {
	GetProducts(c echo.Context) error
	CreateProduct(c echo.Context) error
}

type ProductHandler struct {
	productUsecase usecases.IProductUsecase
}

func NewProductHandler(productUsecase usecases.IProductUsecase) IProductHandler {
	return ProductHandler{
		productUsecase: productUsecase,
	}
}

// GetProducts godoc
// @Summary      Get products
// @Description  get products
// @Tags         product
// @Accept       json
// @Produce      json
// @Success      200  {array}  responses.GetProduct
// @Failure      400  {object}  error
// @Failure      404  {object}  error
// @Failure      500  {object}  error
// @Router       /products [get]
func (p ProductHandler) GetProducts(c echo.Context) error {
	data, err := p.productUsecase.GetProducts()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, data)
}

// CreateProduct godoc
// @Summary      Create product
// @Description  create product
// @Tags         product
// @Accept       json
// @Produce      json
// @Param 		 request body requests.CreateProduct true "body"
// @Success      201  {object}  nil
// @Failure      400  {object}  error
// @Failure      404  {object}  error
// @Failure      500  {object}  error
// @Router       /product [post]
func (p ProductHandler) CreateProduct(c echo.Context) error {
	req := requests.CreateProduct{}
	err := c.Bind(&req)
	if err != nil {
		return err
	}
	err = p.productUsecase.CreateProducts(req)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, nil)
}
