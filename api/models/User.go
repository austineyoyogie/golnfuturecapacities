package models

import (
	"database/sql"
	"errors"
	"gorm.io/gorm"
	"time"
)

var (
	ErrRoleEmptyName = errors.New("role.permission name can't be empty")
)

type User struct {
	ID        uint64       `gorm:"column:id;primary_key;auto_increment" json:"id"`
	FirstName string       `gorm:"column:first_name;size:45;not null;" validate:"required,min=2,max=45" json:"first_name"`
	LastName  string       `gorm:"column:last_name;size:45;not null;" validate:"required,min=2,max=45" json:"last_name" `
	Email     string       `gorm:"column:email;size:45;not null;unique" validate:"required,email" json:"email" `
	Password  string       `gorm:"column:password;size:255;not null;" validate:"required,min=8,max=255" json:"password" binding:"_"`
	Phone     string       `gorm:"column:phone;size:255;not null;" json:"phone"`
	Token     string       `gorm:"column:token;size:255;not null;" json:"token"`
	Verify    sql.NullBool `gorm:"column:verify;default:false" json:"verify"`
	Roles     *[]Role      `gorm:"many2many:user_roles" json:"role"`
	Enabled   sql.NullBool `gorm:"column:enabled;default:false" json:"enabled"`
	Deleted   sql.NullBool `gorm:"column:deleted;default:false" json:"deleted"`
	LoginAt   string       `gorm:"column:login_at;size:45;not null;" json:"login_at"`
	Model
}

type Role struct {
	ID         uint64  `gorm:"column:id;primary_key;auto_increment" json:"id"`
	Name       string  `gorm:"column:name;size:100;not null;" validate:"required,min=2,max=255" json:"name"`
	Permission string  `gorm:"column:permission;size:255;not null;" validate:"required,min=2,max=255" json:"permission"`
	Users      []*User `gorm:"many2many:user_roles" json:"user"`
	Model
}

type UserRole struct {
	ID     uint64 `gorm:"column:id;primary_key;auto_increment" json:"id"`
	UserId uint
	RoleId uint
	Model
}

func (r *Role) Validate() error {
	if r.Name == "" {
		return ErrRoleEmptyName
	}
	return nil
}

func AutoMigration(db *gorm.DB) {
	err := db.Debug().Migrator().AutoMigrate(&Role{}, &User{}, &UserRole{})
	//err := db.Debug().Migrator().DropTable(&User{}, &UserRole{})
	if err != nil {
		return
	}
}

/*
	type User struct {
		ID        uint64       `gorm:"column:id;primary_key;auto_increment" json:"id"`
		FirstName string       `gorm:"column:first_name;size:45;not null;" validate:"required,min=2,max=45" json:"first_name"`
		LastName  string       `gorm:"column:last_name;size:45;not null;" validate:"required,min=2,max=45" json:"last_name" `
		Email     string       `gorm:"column:email;size:45;not null;unique" validate:"required,email" json:"email" `
		Password  string       `gorm:"column:password;size:255;not null;" validate:"required,min=8,max=255" json:"password" binding:"_"`
		Phone     string       `gorm:"column:phone;size:255;not null;" json:"phone"`
		Token     string       `gorm:"column:token;size:255;not null;" json:"token"`
		Verify    sql.NullBool `gorm:"column:verify;default:false" json:"verify"`
		RoleID    uint64       `gorm:"not null" json:"role_id"`
		Roles     *[]Role      `gorm:"ForeignKey:ID" json:"roles"` // It look like it fixed. Need to be fixed A must for user output role permission
		TwoFactor *[]TwoFactor `gorm:"ForeignKey:ID" json:"two_factor"`
		Enabled   sql.NullBool `gorm:"column:enabled;default:false" json:"enabled"`
		LoginAt   string       `gorm:"column:login_at;size:45;not null;" json:"login_at"`
		Model
	}

	type Role struct {
		ID         uint64  `gorm:"column:id;primary_key;auto_increment" json:"id"`
		Name       string  `gorm:"column:name;size:100;not null;" json:"name" validate:"required"`
		Permission string  `gorm:"column:permission;size:255;not null;" validate:"required,min=2,max=255" json:"permission"`
		Users      []*User `gorm:"ForeignKey:RoleID" json:"users"`
		Model
	}
*/
type TwoFactor struct {
	ID        uint64    `gorm:"column:id;primary_key;auto_increment" json:"id"`
	UserId    uint64    `gorm:"column:user_id;size:4;not null;" json:"user_id"`
	Code      string    `gorm:"column:code;size:45;not null;unique" json:"code"`
	ExpiredAt time.Time `gorm:"column:expired_at;not null;" json:"expired_at"`
	Model
}
