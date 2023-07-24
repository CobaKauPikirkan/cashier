package user_service

import "github.com/CobaKauPikirkan/cashier/models"

type User interface {
	FindAll()([]models.User, error)
	FindById(id int)(models.User, error)
	Create(user models.UserRequest)(models.User, error)
	Update(Id int,user models.UserRequest)(models.User, error)
	Delete(Id int)(models.User, error)
	FindByEmail(email string)(models.User, error)
}