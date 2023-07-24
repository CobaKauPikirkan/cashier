package admin_service

import (
	"github.com/CobaKauPikirkan/cashier/helper"
	"github.com/CobaKauPikirkan/cashier/models"
	"github.com/CobaKauPikirkan/cashier/repository/admin"
)
type adminService struct {
	repository admin.RepositoryAdmin
}

func NewAdminService(repository admin.RepositoryAdmin) *adminService {
	return &adminService{repository}
}

func (s *adminService) FindAll() ([]models.User, error) {
	user, err := s.repository.FindAll()
	helper.PanicIfError(err)
	
	return user, err
}

func (s *adminService)FindById(Id int) (models.User, error) {
	user, err := s.repository.FindById(Id)
	helper.PanicIfError(err)
	
	return user, err
}

func (s *adminService)Update(Id int, userrequest models.UserRequest) (models.User, error) {
	user, err:=s.repository.FindById(Id)
	helper.PanicIfError(err)

	user.Nama = userrequest.Nama
	user.Role = userrequest.Role
	user.Username = userrequest.Username
	user.Password = userrequest.Password
	newuser, err:= s.repository.Update(user)
	helper.PanicIfError(err)

	return newuser, err
}

func (s *adminService)Delete(Id int) (models.User, error) {
	User, err:=s.repository.FindById(Id)
	helper.PanicIfError(err)

	newUser, err := s.repository.Delete(User)
	helper.PanicIfError(err)

	return newUser, err
}