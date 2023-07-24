package manager_service

import "github.com/CobaKauPikirkan/cashier/models"

type Manager interface {
	FindAllTransaksi()([]models.Result, error)
	FindByIdKasir(id int)([]models.Result, error)
	FindByTanggal(tgl string, filter string)([]models.Result, error)
	FindMostFavourite()([]models.MostFavouriteMenu, error)
	FindMenuById()(models.Menu, error)
}