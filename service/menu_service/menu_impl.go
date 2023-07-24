package menu_service

import (
	"github.com/CobaKauPikirkan/cashier/helper"
	"github.com/CobaKauPikirkan/cashier/models"
	"github.com/CobaKauPikirkan/cashier/repository/menu"
)
type menuService struct {
	repository menu.RepositoryMenu
}

func NewMenuService(repository menu.RepositoryMenu) *menuService {
	return &menuService{repository}
}

func (s *menuService) FindAll() ([]models.Menu, error) {
	Menu, err := s.repository.FindAll()
	helper.PanicIfError(err)
	
	return Menu, err
}

func (s *menuService)FindById(Id int) (models.Menu, error) {
	Menu, err := s.repository.FindById(Id)
	helper.PanicIfError(err)
	
	return Menu, err
}

func (s *menuService)Create(MenuRequest models.MenuRequest) (models.Menu, error) {
	Menu := models.Menu{
		NamaMenu: MenuRequest.NamaMenu,
		Jenis: MenuRequest.Jenis,
		Deskripsi: MenuRequest.Deskripsi,
		Gambar: MenuRequest.Gambar,
		Harga: MenuRequest.Harga,
	}
	newMenu, err := s.repository.Create(Menu)
	helper.PanicIfError(err)

	return newMenu, err
}

func (s *menuService)Update(Id int, Menurequest models.MenuRequest) (models.Menu, error) {
	Menu, err:=s.repository.FindById(Id)
	helper.PanicIfError(err)

	Menu.NamaMenu = Menurequest.NamaMenu
	Menu.Jenis = Menurequest.Jenis
	Menu.Deskripsi = Menurequest.Deskripsi
	Menu.Gambar = Menurequest.Gambar
	Menu.Harga = Menurequest.Harga

	newMenu, err:= s.repository.Update(Menu)
	helper.PanicIfError(err)

	return newMenu, err
}

func (s *menuService)Delete(Id int) (models.Menu, error) {
	Menu, err:=s.repository.FindById(Id)
	helper.PanicIfError(err)

	newMenu, err := s.repository.Delete(Menu)
	helper.PanicIfError(err)

	return newMenu, err
}