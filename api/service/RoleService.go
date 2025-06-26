package service

import "golnfuturecapacities/api/models"

type RoleService interface {
	Save(*models.Role) (*models.Role, error)
	Find(uint64) (*models.Role, error)
	FindAll() ([]*models.Role, error)
	Exists(name string) (*models.Role, error)
	Update(*models.Role) error
	Delete(uint64) error
}
