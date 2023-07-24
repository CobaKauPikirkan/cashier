package menu_service

import "github.com/CobaKauPikirkan/cashier/models"

type Menu interface {
	FindAll()([]models.Menu, error)
	FindById(id int)(models.Menu, error)
	Create(menu models.MenuRequest)(models.Menu, error)
	Update(Id int,menu models.MenuRequest)(models.Menu, error)
	Delete(Id int)(models.Menu, error)
}