package manager_service

import (
	"github.com/CobaKauPikirkan/cashier/helper"
	"github.com/CobaKauPikirkan/cashier/models"
	"github.com/CobaKauPikirkan/cashier/repository/manager"
)

type ManagerService struct {
	repository manager.RepositoryManager
}

func NewManagerService(repository manager.RepositoryManager) *ManagerService {
	return &ManagerService{repository}
}

func (s *ManagerService) FindAllTransaksi() ([]models.Result, error) {
	result, err := s.repository.FindAllTransaksi()
	helper.PanicIfError(err)

	return result, err
}

func (s *ManagerService) FindByIdKasir(id int) ([]models.Result, error) {
	result, err := s.repository.FindByIdKasir(id)
	helper.PanicIfError(err)

	return result, err
}

func (s *ManagerService) FindByTanggal(tgl string, filter string) ([]models.Result, error) {
	result, err := s.repository.FindByTanggal(tgl, filter)
	helper.PanicIfError(err)

	return result, err
}

func (s *ManagerService) FindMostFavourite() ([]models.MostFavouriteMenu, error) {
	result, err := s.repository.FindMostFavourite()
	helper.PanicIfError(err)

	return result, err
}

func (s *ManagerService) FindMenuById() (models.Menu, error) {
	id:= 1
	menu, err := s.repository.FindByIdMenu(id)
	helper.PanicIfError(err)

	return menu, err
}