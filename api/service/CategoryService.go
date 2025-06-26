package service

import "golnfuturecapacities/api/models/products"

type CategoryService interface {
	Save(category *products.Category) (*products.Category, error)
	Exists(name string) (*products.Category, error)
	Find(uint64) (*products.Category, error)
	FindAll() ([]*products.Category, error)
	Update(category *products.Category) error
	Delete(uint64) error
}
