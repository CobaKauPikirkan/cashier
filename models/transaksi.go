package models

import (
	"time"
)

type Status struct {
	ID   int    `json:"id"`
	Nama string `json:"nama"`
}

type Transaksi struct {
	ID             int       `json:"id"`
	Tgl_Transaksi  time.Time `json:"tgl_transaksi"`
	IdUser         int       `json:"id_user"`
	IdMeja         int       `json:"id_meja"`
	Nama_Pelanggan string    `json:"nama_pelanggan"`
	Status         int       `json:"status"`
}

type DetailTransaksi struct {
	ID           int   `json:"id"`
	Id_transaksi int   `json:"id_transaksi"`
	Id_menu      []int `json:"id_menu" gorm:"-"`
	Harga        int   `json:"harga"`
}

type DetailTransaksi2 struct {
	ID           int   `json:"id"`
	Id_transaksi int   `json:"id_transaksi"`
	Id_menu      string `json:"id_menu"`
	Harga        int   `json:"harga"`
}

type Result struct {
	Tgl_Transaksi  time.Time `json:"tgl_transaksi"`
	IdUser         int       `json:"id_user"`
	IdMeja         int       `json:"id_meja"`
	Nama_Pelanggan string    `json:"nama_pelanggan"`
	Status         int       `json:"status"`
	Id_transaksi int   `json:"id_transaksi"`
	Id_menu      string `json:"id_menu"`
	Harga        int   `json:"harga"`
}
