package usecases

import (
	"backend/clean-arch/deliveries/requests"
	"backend/clean-arch/deliveries/responses"
	"backend/clean-arch/domain/entities"
	"backend/clean-arch/domain/repositories/postgres"
)

type IProductUsecase interface {
	GetProducts() ([]responses.GetProduct, error)
	CreateProducts(req requests.CreateProduct) error
}

type ProductUsecase struct {
	repo postgres.IProductRepository
}

func NewProductUsecase(repo postgres.IProductRepository) IProductUsecase {
	return ProductUsecase{
		repo: repo,
	}
}

func (p ProductUsecase) GetProducts() ([]responses.GetProduct, error) {
	products, err := p.repo.FetchProducts()
	if err != nil {
		return nil, err
	}
	return p.Pack(products), nil
}

func (p ProductUsecase) Pack(req []entities.Product) (result []responses.GetProduct) {
	for _, item := range req {
		result = append(result, responses.GetProduct{
			ID:    item.ID,
			Name:  item.Name,
			Price: item.Price,
		})
	}
	return result
}

func (p ProductUsecase) CreateProducts(req requests.CreateProduct) error {
	return p.repo.InsertProducts(p.Unpack(req))
}

func (p ProductUsecase) Unpack(req requests.CreateProduct) entities.Product {
	return entities.Product{
		Name:   req.Name,
		Price:  req.Price,
		UserID: req.UserID,
	}
}
