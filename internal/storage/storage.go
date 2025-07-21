package storage

import (
	"github.com/vandannandwana/Basic-E-Commerce/internal/types"
)

type Storage interface {
	CreateProduct(productName string, productPrice int64, productDescription string) (int64, error)
	GetProductById(productId int64) (types.Product, error)
	GetProducts () ([] types.Product, error)
	DeleteProductById (productId int64) (bool, error)
}