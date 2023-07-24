package user

import (
	"github.com/CobaKauPikirkan/cashier/models"
	"github.com/CobaKauPikirkan/cashier/helper"
	"gorm.io/gorm"
)

type repositoryUserImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *repositoryUserImpl {
	return &repositoryUserImpl{db}
}

func (r *repositoryUserImpl)FindAll() ([]models.User, error) {
	var User []models.User
	
	err:= r.db.Table("user").Find(&User).Error
	helper.PanicIfError(err)
	
	return User, err 
}
func (r *repositoryUserImpl)FindById(id int) (models.User, error) {
	var User models.User
	
	err:= r.db.Table("user").Find(&User, id).Error
	helper.PanicIfError(err)
	
	return User, err 
}

func (r *repositoryUserImpl)Create(User models.User) (models.User, error) {
	err:= r.db.Table("user").Create(&User).Error
	helper.PanicIfError(err)	
	return User, err 
}

func (r *repositoryUserImpl)Update(User models.User) (models.User, error) {
	err:= r.db.Table("user").Save(&User).Error
	helper.PanicIfError(err)
	return User, err 
}

func (r *repositoryUserImpl)Delete(User models.User) (models.User, error) {
	err:= r.db.Table("user").Delete(&User).Error
	helper.PanicIfError(err)
	return User, err
}

func (r *repositoryUserImpl)FindByEmail(email string) (models.User, error) {
	var User models.User
	
	err:= r.db.Table("user").Where("username = ?", email).Find(&User).Error
	helper.PanicIfError(err)
	
	return User, err 
}