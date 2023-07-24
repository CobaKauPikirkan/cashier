package models

type Test string

type Jenis struct {
	ID   int    `json:"id"`
	Name string `json:"nama"`
}

type Menu struct {
	ID        int    `json:"id"`
	NamaMenu  string `json:"nama_menu"`
	Jenis     int    `json:"jenis"`
	Deskripsi string `json:"deskripsi"`
	Gambar    string `json:"gambar"`
	Harga     int    `json:"harga"`
}

type MostFavouriteMenu struct {
    Menu  Menu
    Count int
}