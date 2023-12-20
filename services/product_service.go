package services

import (
	"errors"
	"tesbaru/entity"
	"tesbaru/repository"
)

type ProductService struct {
	ProductRepo repository.ProductRepo
}

func (service ProductService) GetProductById(id string) (*entity.Product, error) {
	product := service.ProductRepo.FindById(id)
	if product == nil {
		return nil, errors.New("product not found")
	}

	return product, nil
}
