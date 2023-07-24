package menu

import (
	"github.com/CobaKauPikirkan/cashier/models"
	"github.com/CobaKauPikirkan/cashier/helper"
	"gorm.io/gorm"
)

type repositoryMenuImpl struct {
	db *gorm.DB
}

func NewMenuRepository(db *gorm.DB) *repositoryMenuImpl {
	return &repositoryMenuImpl{db}
}

func (r *repositoryMenuImpl)FindAll() ([]models.Menu, error) {
	var Menu []models.Menu
	
	err:= r.db.Table("menu").Find(&Menu).Error
	helper.PanicIfError(err)
	
	return Menu, err 
}
func (r *repositoryMenuImpl)FindById(id int) (models.Menu, error) {
	var Menu models.Menu
	
	err:= r.db.Table("menu").Find(&Menu, id).Error
	helper.PanicIfError(err)
	
	return Menu, err 
}

func (r *repositoryMenuImpl)Create(Menu models.Menu) (models.Menu, error) {
	err:= r.db.Table("menu").Create(&Menu).Error
	helper.PanicIfError(err)	
	return Menu, err 
}

func (r *repositoryMenuImpl)Update(Menu models.Menu) (models.Menu, error) {
	err:= r.db.Table("menu").Save(&Menu).Error
	helper.PanicIfError(err)
	return Menu, err 
}

func (r *repositoryMenuImpl)Delete(Menu models.Menu) (models.Menu, error) {
	err:= r.db.Table("menu").Delete(&Menu).Error
	helper.PanicIfError(err)
	return Menu, err
}