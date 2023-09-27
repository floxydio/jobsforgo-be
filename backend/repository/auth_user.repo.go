package repository

import (
	"jobsforgobe/database"
	"jobsforgobe/entities"

	"gorm.io/gorm"
)

type SqlStruct struct {
	Db *gorm.DB
}

type AuthUserRepository interface {
	SignUpAccount(entities.AuthUser) error
	SignInAccout(entities.AuthUser) (entities.AuthUser, error)
}

func AuthUserInterface(db *gorm.DB) AuthUserRepository {
	return &SqlStruct{Db: database.DBAccess}
}

func (r *SqlStruct) SignUpAccount(form entities.AuthUser) error {
	err := r.Db.Create(&form).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *SqlStruct) SignInAccout(form entities.AuthUser) (entities.AuthUser, error) {
	var user entities.AuthUser
	err := r.Db.Where("username = ? OR email = ?", form.Username, form.Email).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
