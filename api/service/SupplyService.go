package service

import (
	"golnfuturecapacities/api/models/products"
)

type SupplyService interface {
	Save(supply *products.Supply) (*products.Supply, error)
	Exists(name string) (*products.Supply, error)
	Find(uint64) (*products.Supply, error)
	FindAll() ([]*products.Supply, error)
	Update(supply *products.Supply) error
	Delete(uint64) error
}
