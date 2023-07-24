package meja_service

import (
	"github.com/CobaKauPikirkan/cashier/helper"
	"github.com/CobaKauPikirkan/cashier/models"
	"github.com/CobaKauPikirkan/cashier/repository/meja"
)
type mejaService struct {
	repository meja.RepositoryMeja
}

func NewMejaService(repository meja.RepositoryMeja) *mejaService {
	return &mejaService{repository}
}

func (s *mejaService) FindAll() ([]models.Meja, error) {
	meja, err := s.repository.FindAll()
	helper.PanicIfError(err)
	
	return meja, err
}

func (s *mejaService)FindById(Id int) (models.Meja, error) {
	meja, err := s.repository.FindById(Id)
	helper.PanicIfError(err)
	
	return meja, err
}

func (s *mejaService)Create(mejaRequest models.MejaRequest) (models.Meja, error) {
	meja := models.Meja{
		NomorMeja: mejaRequest.NomorMeja,
		IsAvaible: mejaRequest.IsAvaible,
	}
	newMeja, err := s.repository.Create(meja)
	helper.PanicIfError(err)

	return newMeja, err
}

func (s *mejaService)Update(Id int, mejarequest models.MejaRequest) (models.Meja, error) {
	meja, err:=s.repository.FindById(Id)
	helper.PanicIfError(err)

	meja.NomorMeja = mejarequest.NomorMeja
	meja.IsAvaible = mejarequest.IsAvaible

	newMeja, err:= s.repository.Update(meja)
	helper.PanicIfError(err)

	return newMeja, err
}

func (s *mejaService)Delete(Id int) (models.Meja, error) {
	meja, err:=s.repository.FindById(Id)
	helper.PanicIfError(err)

	newMeja, err := s.repository.Delete(meja)
	helper.PanicIfError(err)

	return newMeja, err
}