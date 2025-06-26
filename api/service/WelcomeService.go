package service

import "golnfuturecapacities/api/models"

type WelcomeService interface {
	Find(uint64) (*models.User, error)
}
