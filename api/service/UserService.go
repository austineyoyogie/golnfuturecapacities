package service

import "golnfuturecapacities/api/models"

type UserService interface {
	Save(user *models.User) (*models.User, error)
	Exists(email string) (*models.User, error)
	IsEnabled(email string) (*models.User, error)
	Find(uint64) (*models.User, error)
	FindAll() ([]*models.User, error)
	Update(user *models.User) error
	Delete(uint64) error
	AddToUserRole(*models.UserRole) (*models.UserRole, error)
	DeleteVerificationCode(code *models.TwoFactor) error
	SendVerificationCode(code *models.TwoFactor) (*models.TwoFactor, error)
}
