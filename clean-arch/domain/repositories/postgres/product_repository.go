package postgres

import (
	"backend/clean-arch/domain/entities"

	"gorm.io/gorm"
)

type IProductRepository interface {
	FetchProducts() ([]entities.Product, error)
	InsertProducts(entities.Product) error
}

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) IProductRepository {
	return ProductRepository{
		db: db,
	}
}

func (p ProductRepository) FetchProducts() (products []entities.Product, err error) {
	err = p.db.Find(&products).Error
	return products, err
}

func (p ProductRepository) InsertProducts(req entities.Product) error {
	return p.db.Create(&req).Error
}
