package user_service

import (
	"github.com/CobaKauPikirkan/cashier/helper"
	"github.com/CobaKauPikirkan/cashier/models"
	"github.com/CobaKauPikirkan/cashier/repository/user"
)
type userService struct {
	repository user.RepositoryUser
}

func NewUserService(repository user.RepositoryUser) *userService {
	return &userService{repository}
}

func (s *userService) FindAll() ([]models.User, error) {
	user, err := s.repository.FindAll()
	helper.PanicIfError(err)
	
	return user, err
}

func (s *userService)FindById(Id int) (models.User, error) {
	user, err := s.repository.FindById(Id)
	helper.PanicIfError(err)
	
	return user, err
}

func (s *userService)Create(userRequest models.UserRequest) (models.User, error) {
	user := models.User{
		Nama: userRequest.Nama,
		Role: userRequest.Role,
		Username: userRequest.Username,
		Password: userRequest.Password,
	}
	newuser, err := s.repository.Create(user)
	helper.PanicIfError(err)

	return newuser, err
}

func (s *userService)Update(Id int, userRequest models.UserRequest) (models.User, error) {
	user, err:=s.repository.FindById(Id)
	helper.PanicIfError(err)

	user.Nama = userRequest.Nama
	user.Role = userRequest.Role
	user.Username = userRequest.Username
	user.Password = userRequest.Password

	newuser, err:= s.repository.Update(user)
	helper.PanicIfError(err)

	return newuser, err
}

func (s *userService)Delete(Id int) (models.User, error) {
	user, err:=s.repository.FindById(Id)
	helper.PanicIfError(err)

	newuser, err := s.repository.Delete(user)
	helper.PanicIfError(err)

	return newuser, err
}

func (s *userService)FindByEmail(email string) (models.User, error) {
	user, err := s.repository.FindByEmail(email)
	helper.PanicIfError(err)
	
	return user, err
}