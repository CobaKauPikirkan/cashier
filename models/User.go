package models

type Role struct {
	IdRole int    `json:"id"`
	Nama   string `json:"nama"`
}

type User struct {
	ID       int    `json:"id"`
	Nama     string `json:"nama"`
	Role     int  `json:"role"`
	Username string `json:"username"`
	Password string `json:"password"`
}
