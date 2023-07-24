package models

type MejaRequest struct {
	NomorMeja string `json:"nomor_meja"`
	IsAvaible string `json:"is_avaible"`
}

type MenuRequest struct {
	NamaMenu  string `json:"nama_menu"`
	Jenis     int    `json:"jenis"`
	Deskripsi string `json:"deskripsi"`
	Gambar    string `json:"gambar"`
	Harga     int    `json:"harga"`
}

type UserRequest struct {
	Nama         string `json:"nama"`
	Role         int    `json:"role"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	RefreshToken string `json:"refresh_token"`
}

type KasirRequest struct {
	IdUser         int       `json:"id_user"`
	IdMeja         int       `json:"id_meja"`
	Nama_Pelanggan string    `json:"nama_pelanggan"`
	Status         int       `json:"status"`
	Id_menu        []int     `json:"id_menu"`
	Harga          int       `json:"harga"`
}