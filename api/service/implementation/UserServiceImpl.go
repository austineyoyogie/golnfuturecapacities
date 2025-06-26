package implementation

import (
	"fmt"
	"golnfuturecapacities/api/messages"
	"golnfuturecapacities/api/models"
	"golnfuturecapacities/api/service"
	"golnfuturecapacities/api/utils"
	"gorm.io/gorm"
	"time"
)

type UserServiceImpl struct {
	db *gorm.DB
}

func NewUserServiceImpl(db *gorm.DB) service.UserService {
	return &UserServiceImpl{db}
}
func (u *UserServiceImpl) AddToUserRole(userRole *models.UserRole) (*models.UserRole, error) {
	var user models.User
	tx := u.db.Begin()
	tx.Debug().Model(&models.User{}).Where("id = ?", userRole.UserId).First(&user)

	columns := map[string]interface{}{
		"user_id": user.ID,
		"role_id": true,
	}
	err := tx.Debug().Model(&models.UserRole{}).Create(&columns).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	return userRole, tx.Commit().Error
}

func (u UserServiceImpl) Save(user *models.User) (*models.User, error) {
	tx := u.db.Begin()
	user.FirstName = utils.Escape(user.FirstName)
	user.LastName = utils.Escape(user.LastName)
	user.Email = utils.IsToLower(user.Email)
	user.Phone = user.Phone // need to be validate with regex all field
	hash, _ := utils.HashPassword(user.Password)
	user.Password = string(hash)

	token := utils.RandomString(10)
	user.Token = token
	err := tx.Debug().Model(&models.User{}).Create(user).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	subject := "Account Email Verify"
	msg := messages.Deliver([]string{user.Email}, subject)
	activate := fmt.Sprintf("http://localhost:8080/api/v1/verify?email=%s&token=%s", user.Email, user.Token)
	msg.EmailTemplate("api/messages/verify.html", activate)
	return user, tx.Commit().Error
}
func (u UserServiceImpl) Exists(email string) (*models.User, error) {
	user := &models.User{}
	err := u.db.Debug().Model(&models.User{}).Where("email = ?", email).Preload("Roles").Find(user).Error
	if err != nil {
		return nil, err
	}
	return user, err
}
func (u UserServiceImpl) IsEnabled(email string) (*models.User, error) {
	user := &models.User{}
	err := u.db.Debug().Model(&models.User{}).Where("email = ?", email).Find(user).Where("enabled = ?", true).Find(user).Error
	if err != nil {
		return nil, err
	}
	return user, err
}
func (u *UserServiceImpl) Find(userId uint64) (*models.User, error) {
	user := &models.User{}
	err := u.db.Debug().Model(&models.User{}).Where("id = ?", userId).Preload("Roles").Find(&user).Error
	if err != nil {
		return nil, err
	}
	return user, err
}
func (u *UserServiceImpl) FindAll() ([]*models.User, error) {
	var users []*models.User
	err := u.db.Debug().Model(&models.User{}).Preload("Roles").Find(&users).Error
	return users, err
}
func (u *UserServiceImpl) Update(user *models.User) error {
	tx := u.db.Begin()
	columns := map[string]interface{}{
		"first_name": user.FirstName,
		"last_name":  user.LastName,
		"phone":      user.Phone,
		"updated_at": time.Now(),
	}
	err := tx.Debug().Model(&models.User{}).Where("id = ?", user.ID).UpdateColumns(columns).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
func (u *UserServiceImpl) Delete(userId uint64) error {
	tx := u.db.Begin()
	columns := map[string]interface{}{
		"active":     false,
		"enabled":    false,
		"deleted":    true,
		"deleted_at": time.Now(),
	}
	err := tx.Debug().Model(&models.User{}).Where("id = ?", userId).UpdateColumns(&columns).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (u UserServiceImpl) SendVerificationCode(code *models.TwoFactor) (*models.TwoFactor, error) {
	tx := u.db.Begin()
	expireTime := 15 * time.Minute
	factor := utils.TwoFactorCode(6)
	code.Code = factor
	code.ExpiredAt = time.Now().Add(expireTime) // Not sure

	err := tx.Debug().Model(&models.TwoFactor{}).Create(code).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	return code, tx.Commit().Error
}

// Rework Needed -> DeleteVerificationCode
func (u *UserServiceImpl) DeleteVerificationCode(userId *models.TwoFactor) error {
	tx := u.db.Begin()
	err := tx.Debug().Model(&models.TwoFactor{}).Where("id = ?", userId).Delete(&models.TwoFactor{}).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
