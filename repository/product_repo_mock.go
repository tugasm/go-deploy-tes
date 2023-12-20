package repository

import (
	"github.com/stretchr/testify/mock"
	"tesbaru/entity"
)

type ProductRepoMock struct {
	Mock mock.Mock
}

func (repoMock *ProductRepoMock) FindById(id string) *entity.Product {
	arguments := repoMock.Mock.Called(id)

	if arguments.Get(0) == nil {
		return nil
	}
	product := arguments.Get(0).(entity.Product)

	return &product
}
