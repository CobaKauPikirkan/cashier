package manager

import (
	"errors"
	"fmt"

	"github.com/CobaKauPikirkan/cashier/helper"
	"github.com/CobaKauPikirkan/cashier/models"
	"gorm.io/gorm"
)

type repositoryManagerImpl struct {
	db *gorm.DB
}

func NewManagerRepository(db *gorm.DB) *repositoryManagerImpl {
	return &repositoryManagerImpl{db}
}

func (r *repositoryManagerImpl) FindAllTransaksi() ([]models.Result, error) {
	var results []models.Result
	err := r.db.Table("detail_transaksi").Joins("left join transaksi on transaksi.id = detail_transaksi.id_transaksi").Scan(&results).Error
	helper.PanicIfError(err)

	return results, err
}

func (r *repositoryManagerImpl) FindByIdKasir(id int) ([]models.Result, error) {
    var results []models.Result
    err := r.db.Table("transaksi").Joins("left join detail_transaksi on transaksi.id = detail_transaksi.id_transaksi").Where("transaksi.id_user = ?", id).Scan(&results).Error
    helper.PanicIfError(err)

    return results, err
}

func (r *repositoryManagerImpl) FindByTanggal(tgl string, filter string) ([]models.Result, error) {
	var results []models.Result
	err := r.db.Table("transaksi").Joins("left join detail_transaksi on detail_transaksi.id_transaksi = transaksi.id").Where("DATE(tgl_transaksi) = ?", "%"+tgl+"%").Scan(&results).Error
	helper.PanicIfError(err)
	switch filter {
	case "date":
		err = r.db.Table("transaksi").
			Joins("LEFT JOIN detail_transaksi ON detail_transaksi.id_transaksi = transaksi.id").
			Where("DATE(tgl_transaksi) = ?", tgl).
			Scan(&results).
			Error
	case "month":
		err = r.db.Table("transaksi").
			Joins("LEFT JOIN detail_transaksi ON detail_transaksi.id_transaksi = transaksi.id").
			Where("MONTH(tgl_transaksi) = MONTH(?) AND YEAR(tgl_transaksi) = YEAR(?)", tgl, tgl).
			Scan(&results).
			Error
	case "year":
		err = r.db.Table("transaksi").
			Joins("LEFT JOIN detail_transaksi ON detail_transaksi.id_transaksi = transaksi.id").
			Where("YEAR(tgl_transaksi) = YEAR(?)", tgl).
			Scan(&results).
			Error
	default:
		err = errors.New("invalid filter")
	}

	if err != nil {
		return nil, err
	}

	return results, err
}

// err :=  r.db.Table("detail_transaksi").
    // Select("GROUP_CONCAT(DISTINCT detail_transaksi.id_menu SEPARATOR ',') as id_menus, COUNT(*) as count").
    // Joins("JOIN menu ON FIND_IN_SET(menu.id, detail_transaksi.id_menu)").
    // Group("detail_transaksi.id_menu").
    // Order("count DESC").
    // Scan(&results).
    // Error

// func (r *repositoryManagerImpl) FindMostFavourite() ([]models.MostFavouriteMenu, error) {
// 	var results []struct {
// 		IdMenus string 
// 		Count   int
// 	}

// 	err := r.db.Table("menu").
//         Select("GROUP_CONCAT(DISTINCT menu.id SEPARATOR ',') as id_menus, COUNT(detail_transaksi.id) as count").
//         Joins("LEFT JOIN detail_transaksi ON menu.id = detail_transaksi.id_menu").
//         Group("menu.id").
//         Order("count DESC").
//         Scan(&results).
//         Error

// 	if err != nil {
// 		return nil, err
// 	}

// 	var menus []models.MostFavouriteMenu
// 	menuMap := make(map[int]bool)
// 	for _, result := range results {
//     idMenus := strings.Split(result.IdMenus, ",")
//     for _, idMenu := range idMenus {
//         id, _ := strconv.Atoi(idMenu)
//         if _, ok := menuMap[id]; !ok {
//             menu, err := r.FindByIdMenu(id)
//             if err != nil {
//                 return nil, err
//             }
//             // Modify the mapping of the `Menu` struct to include the `id` field
//             menus = append(menus, models.MostFavouriteMenu{
//                 Menu: models.Menu{
//                     ID:          menu.ID,
//                     NamaMenu:    menu.NamaMenu,
//                     Jenis:       menu.Jenis,
//                     Deskripsi:   menu.Deskripsi,
//                     Gambar:      menu.Gambar,
//                     Harga:       menu.Harga,
//                 },
//                 Count: result.Count,
//             })
//             menuMap[id] = true
//         }
//     }
// }

// fmt.Println(results)

// 	return menus, nil
// }

func (r *repositoryManagerImpl) FindMostFavourite() ([]models.MostFavouriteMenu, error) {
    var results []struct {
        models.Menu
        Count int
    }

    err := r.db.Table("menu").
        Select("menu.*, (SELECT COUNT(*) FROM detail_transaksi WHERE FIND_IN_SET(menu.id, detail_transaksi.id_menu)) as count").
        Order("count DESC").
        Scan(&results).
        Error
    if err != nil {
        return nil, err
    }

    var menus []models.MostFavouriteMenu
    for _, result := range results {
        menus = append(menus, models.MostFavouriteMenu{
            Menu:  result.Menu,
            Count: result.Count,
        })
    }

	fmt.Println(results)

    return menus, nil
}

func (r *repositoryManagerImpl)FindByIdMenu(id int) (models.Menu, error) {
	var Menu models.Menu
	
	err:= r.db.Table("menu").Find(&Menu, id).Error
	helper.PanicIfError(err)
	
	return Menu, err 
}
