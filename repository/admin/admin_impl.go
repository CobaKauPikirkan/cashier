package admin

import (
	"github.com/CobaKauPikirkan/cashier/models"
	"github.com/CobaKauPikirkan/cashier/helper"
	"gorm.io/gorm"
)

type repositoryAdminImpl struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) *repositoryAdminImpl {
	return &repositoryAdminImpl{db}
}

func (r *repositoryAdminImpl)FindAll() ([]models.User, error) {
	var user []models.User
	
	err:= r.db.Table("user").Find(&user).Error
	helper.PanicIfError(err)
	
	return user, err 
}
func (r *repositoryAdminImpl)FindById(id int) (models.User, error) {
	var user models.User
	
	err:= r.db.Table("user").Find(&user, id).Error
	helper.PanicIfError(err)
	
	return user, err 
}

func (r *repositoryAdminImpl)Update(user models.User) (models.User, error) {
	err:= r.db.Table("user").Save(&user).Error
	helper.PanicIfError(err)
	return user, err 
}

func (r *repositoryAdminImpl)Delete(user models.User) (models.User, error) {
	err:= r.db.Table("user").Delete(&user).Error
	helper.PanicIfError(err)
	return user, err
}