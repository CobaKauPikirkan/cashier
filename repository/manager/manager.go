package manager

import (
	"github.com/CobaKauPikirkan/cashier/models"
)

type RepositoryManager interface {
	FindAllTransaksi() ([]models.Result, error)
	FindByIdKasir(id int)([]models.Result, error)
	FindByTanggal(tgl string, filter string)([]models.Result, error)
	FindMostFavourite()([]models.MostFavouriteMenu, error)
	FindByIdMenu(id int)(models.Menu, error)
}