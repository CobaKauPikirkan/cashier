package kasir

import "github.com/CobaKauPikirkan/cashier/models"

type RepositoryKasir interface {
	FindAll() ([]models.Result, error)
	FindByIdTransaksi(id int) (models.Transaksi, error)
	FindByIdDetailTransaksi(id int) (models.DetailTransaksi, error)
	CreateTransaksi(transaksi models.Transaksi) (models.Transaksi, error)
	CreateDetailTransaksi(detail models.DetailTransaksi) (models.DetailTransaksi, error)
	SumMenu(idMenu []int) (int, error)
	UpdateTransaksi(transaksi models.Transaksi) (models.Transaksi, error)
	UpdateDetailTransaksi(DetailTransaksi models.DetailTransaksi) (models.DetailTransaksi, error)
	// Delete(transaksi models.Transaksi, detail models.DetailTransaksi) (models.Result, error)
	UpdateMeja(meja models.Meja) (models.Meja, error)
	FindMejaAvaible() ([]models.Meja, error)
	FindMeja(id int)(models.Meja, error)
}
