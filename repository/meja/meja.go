package meja

import "github.com/CobaKauPikirkan/cashier/models"

type RepositoryMeja interface {
	FindAll() ([]models.Meja, error)
	FindById(id int)(models.Meja, error)
	Create(meja models.Meja)(models.Meja, error)
	Update(meja models.Meja)(models.Meja, error)
	Delete(meja models.Meja)(models.Meja, error)
}