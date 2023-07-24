package user

import "github.com/CobaKauPikirkan/cashier/models"

type RepositoryUser interface {
	FindAll() ([]models.User, error)
	FindById(id int) (models.User, error)
	Create(user models.User) (models.User, error)
	Update(user models.User) (models.User, error)
	Delete(user models.User) (models.User, error)
	FindByEmail(email string) (models.User, error)
}