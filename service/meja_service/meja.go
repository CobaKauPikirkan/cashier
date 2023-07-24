package meja_service

import "github.com/CobaKauPikirkan/cashier/models"

type Meja interface {
	FindAll()([]models.Meja, error)
	FindById(id int)(models.Meja, error)
	Create(meja models.MejaRequest)(models.Meja, error)
	Update(Id int,menu models.MejaRequest)(models.Meja, error)
	Delete(Id int)(models.Meja, error)
}