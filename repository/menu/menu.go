package menu

import "github.com/CobaKauPikirkan/cashier/models"

type RepositoryMenu interface {
	FindAll() ([]models.Menu, error)
	FindById(id int) (models.Menu, error)
	Create(menu models.Menu) (models.Menu, error)
	Update(menu models.Menu) (models.Menu, error)
	Delete(menu models.Menu) (models.Menu, error)
}