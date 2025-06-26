package implementation

import (
	"golnfuturecapacities/api/models"
	"golnfuturecapacities/api/service"
	"gorm.io/gorm"
	"time"
)

type RoleServiceImpl struct {
	db *gorm.DB
}

func NewRoleServiceImpl(db *gorm.DB) service.RoleService {
	return &RoleServiceImpl{db}
}

func (r *RoleServiceImpl) Save(role *models.Role) (*models.Role, error) {
	tx := r.db.Begin()
	err := tx.Debug().Model(&models.Role{}).Create(role).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	return role, tx.Commit().Error
}

func (r *RoleServiceImpl) Find(roleId uint64) (*models.Role, error) {
	role := &models.Role{}
	err := r.db.Debug().Model(&models.Role{}).Where("id = ?", roleId).Preload("Users", role.Users).Find(&role).Error
	if err != nil {
		return nil, err
	}
	return role, err
}

func (r *RoleServiceImpl) FindAll() ([]*models.Role, error) {
	var roles []*models.Role
	err := r.db.Debug().Model(&models.Role{}).Find(&roles).Error
	return roles, err
}

func (r *RoleServiceImpl) Exists(name string) (*models.Role, error) {
	role := &models.Role{}
	err := r.db.Debug().Model(&models.Role{}).Where("name = ?", name).Find(role).Error
	if err != nil {
		return nil, err
	}
	return role, err
}

func (r *RoleServiceImpl) Update(role *models.Role) error {
	tx := r.db.Begin()
	columns := map[string]interface{}{
		"name":       role.Name,
		"permission": role.Permission,
		"updated_at": time.Now(),
	}
	err := tx.Debug().Model(&models.Role{}).Where("id = ?", role.ID).UpdateColumns(columns).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (r *RoleServiceImpl) Delete(roleId uint64) error {
	tx := r.db.Begin()
	err := tx.Debug().Model(&models.Role{}).Where("id = ?", roleId).Delete(&models.Role{}).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
