package kasir

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/CobaKauPikirkan/cashier/helper"
	"github.com/CobaKauPikirkan/cashier/models"
	"gorm.io/gorm"
)

type repositoryKasirImpl struct {
	db *gorm.DB
}

func NewKasirRepository(db *gorm.DB) *repositoryKasirImpl {
	return &repositoryKasirImpl{db}
}

func (r *repositoryKasirImpl)FindAll() ([]models.Result, error){
	var results []models.Result
	err:= r.db.Table("detail_transaksi").Joins("left join transaksi on transaksi.id = detail_transaksi.id_transaksi").Scan(&results).Error
	helper.PanicIfError(err)

	return results, err
}

func (r *repositoryKasirImpl)FindByIdTransaksi(id int) (models.Transaksi, error) {
	var transaksi models.Transaksi
	
	err:= r.db.Table("transaksi").Find(&transaksi, id).Error
	helper.PanicIfError(err)
	
	return transaksi, err 
}

func (r *repositoryKasirImpl)FindByIdDetailTransaksi(id int) (models.DetailTransaksi, error) {
	var detailString models.DetailTransaksi2
	
	err:= r.db.Table("detail_transaksi").Find(&detailString, id).Error
	helper.PanicIfError(err)

	idMenuStr := strings.Split(detailString.Id_menu, ",")
	idMenu := make([]int, len(idMenuStr))
	for i, id := range idMenuStr {
		num, err := strconv.Atoi(id)
		if err != nil {
			return models.DetailTransaksi{}, err
		}
		idMenu[i] = num
	}
	detail :=  models.DetailTransaksi{
		ID: detailString.ID,
		Id_transaksi: detailString.Id_transaksi,
		Id_menu: idMenu,
		Harga: detailString.Harga,
	}
	
	return detail, err 
}

func (r *repositoryKasirImpl)CreateTransaksi(transaksi models.Transaksi) (models.Transaksi, error) {
	err:= r.db.Table("transaksi").Create(&transaksi).Error
	helper.PanicIfError(err)	

	return transaksi, err 
}

func (r *repositoryKasirImpl)CreateDetailTransaksi(detail models.DetailTransaksi) (models.DetailTransaksi, error) {
	// idMenuStr := make([]string, len(detail.Id_menu))
    // for i, id := range detail.Id_menu {
    //     idMenuStr[i] = strconv.Itoa(id)
    // }
	idMenuStr := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(detail.Id_menu)), ","), "[]")


	var insert models.DetailTransaksi2

	insert.Id_transaksi = detail.Id_transaksi
    insert.Id_menu = idMenuStr
	insert.Harga = detail.Harga

	err := r.db.Table("detail_transaksi").Create(&insert).Error
	helper.PanicIfError(err)

	return detail, err 
} 
func (r *repositoryKasirImpl) SumMenu(idMenu []int) (int, error) {
	var total int
	err := r.db.Table("menu").
		Select("SUM(harga) as total").
		Where("id IN (?)", idMenu).
		Row().
		Scan(&total)
	if err != nil {
		return 0, err
	}
	return total, nil
}

func (r *repositoryKasirImpl)UpdateTransaksi(transaksi models.Transaksi) (models.Transaksi, error) {
	err := r.db.Table("transaksi").Save(&transaksi).Error
	helper.PanicIfError(err)

	return transaksi, err 
} 

func (r *repositoryKasirImpl)UpdateDetailTransaksi(detail models.DetailTransaksi) (models.DetailTransaksi, error) {
	// err := r.db.Table("detail_transaksi").Save(&DetailTransaksi).Error
	// helper.PanicIfError(err)

	// return DetailTransaksi, err 
	idMenuStr := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(detail.Id_menu)), ","), "[]")


	var insert models.DetailTransaksi2

	insert.Id_transaksi = detail.Id_transaksi
    insert.Id_menu = idMenuStr
	insert.Harga = detail.Harga

	err := r.db.Table("detail_transaksi").Save(&insert).Error
	helper.PanicIfError(err)

	return detail, err 
} 

func (r *repositoryKasirImpl)FindMejaAvaible() ([]models.Meja, error) {
	var meja []models.Meja
	
	err := r.db.Table("meja").Where("is_avaible = ?", "avaible").Find(&meja).Error
	helper.PanicIfError(err)
	
	return meja, err 
}

func (r *repositoryKasirImpl)UpdateMeja(meja models.Meja) (models.Meja, error) {
	err:= r.db.Table("meja").Save(&meja).Error
	helper.PanicIfError(err)
	return meja, err 
}
func (r *repositoryKasirImpl)FindMeja(id int) (models.Meja, error) {
	var meja models.Meja
	
	err:= r.db.Table("meja").Find(&meja, id).Error
	helper.PanicIfError(err)
	
	return meja, err 
}