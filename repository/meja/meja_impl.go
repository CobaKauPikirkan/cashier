package meja

import (
	"github.com/CobaKauPikirkan/cashier/models"
	"github.com/CobaKauPikirkan/cashier/helper"
	"gorm.io/gorm"
)

type repositoryMejaImpl struct {
	db *gorm.DB
}

func NewMejaRepository(db *gorm.DB) *repositoryMejaImpl {
	return &repositoryMejaImpl{db}
}

func (r *repositoryMejaImpl)FindAll() ([]models.Meja, error) {
	var meja []models.Meja
	
	err:= r.db.Table("meja").Find(&meja).Error
	helper.PanicIfError(err)
	
	return meja, err 
}
func (r *repositoryMejaImpl)FindById(id int) (models.Meja, error) {
	var meja models.Meja
	
	err:= r.db.Table("meja").Find(&meja, id).Error
	helper.PanicIfError(err)
	
	return meja, err 
}

func (r *repositoryMejaImpl)Create(meja models.Meja) (models.Meja, error) {
	err:= r.db.Table("meja").Create(&meja).Error
	helper.PanicIfError(err)	
	return meja, err 
}

func (r *repositoryMejaImpl)Update(meja models.Meja) (models.Meja, error) {
	err:= r.db.Table("meja").Save(&meja).Error
	helper.PanicIfError(err)
	return meja, err 
}

func (r *repositoryMejaImpl)Delete(meja models.Meja) (models.Meja, error) {
	err:= r.db.Table("meja").Delete(&meja).Error
	helper.PanicIfError(err)
	return meja, err
}