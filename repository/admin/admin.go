package admin

import "github.com/CobaKauPikirkan/cashier/models"

type RepositoryAdmin interface {
	FindAll() ([]models.User, error)
	FindById(id int)(models.User, error)
	Update(user models.User)(models.User, error)
	Delete(user models.User)(models.User, error)
}