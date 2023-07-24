package admin_service

import "github.com/CobaKauPikirkan/cashier/models"

type Admin interface {
	FindAll()([]models.User, error)
	FindById(id int)(models.User, error)
	Update(Id int,user models.UserRequest)(models.User, error)
	Delete(Id int)(models.User, error)
}