package kasir_service

import "github.com/CobaKauPikirkan/cashier/models"

type Kasir interface {
	FindAll() ([]models.Result, error)
	FindById(id int)(models.Transaksi,models.DetailTransaksi, error)
	CreateTransaksi(models.KasirRequest)(models.Transaksi,models.DetailTransaksi, error)
	UpdateTransaksi(Id int, request models.KasirRequest)(models.Transaksi, models.DetailTransaksi, error)
	FindMejaAvaible()([]models.Meja, error)
	UpdateMeja(id int, request models.MejaRequest)(models.Meja, error)
}