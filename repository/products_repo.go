package repository

import "tesbaru/entity"

type ProductRepo interface {
	FindById(id string) *entity.Product
}
