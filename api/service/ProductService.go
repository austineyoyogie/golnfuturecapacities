package service

import (
	"golnfuturecapacities/api/models/products"
)

type ProductService interface {
	Save(product *products.Product) (*products.Product, error)
	Exists(name string) (*products.Product, error)
	Find(uint64) (*products.Product, error)
	FindAll() ([]*products.Product, error)
	Update(product *products.Product) error
	Delete(uint64) error
	AddToProductCategory(*products.ProductCategory) (*products.ProductCategory, error)
}
