package handler

import (
	"log"
	"net/http"

	"github.com/CobaKauPikirkan/cashier/middleware"
	"github.com/CobaKauPikirkan/cashier/models"
	"github.com/CobaKauPikirkan/cashier/service/user_service"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	userService user_service.User
}

func NewUserHandler(userService user_service.User) *UserHandler {
	return &UserHandler{userService}
}

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}

func VerifyPassword(userPassword string, givenPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(givenPassword), []byte(userPassword))
	valid := true
	msg := ""

	if err != nil {
		msg = "login or password incorrect"
		valid = false
	}

	return valid, msg
}

func (s *UserHandler) SignUp(c *gin.Context) {
	var request models.UserRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	hashedPassword := HashPassword(request.Password)
	request.Password = hashedPassword

	user, err := s.userService.Create(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"user": user})
}

func (s *UserHandler) Login(c *gin.Context) {
	var request models.UserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := s.userService.FindByEmail(request.Username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Username or password is incorrect"})
		return
	}

	if isValid, _ := VerifyPassword(request.Password, user.Password); !isValid {
		c.JSON(http.StatusUnauthorized, gin.H{"errorVerify": "Username or password is incorrect"})
		return
	}

	// create JWT token here
	token, err := middleware.GenerateToken(uint(user.ID), []int{user.Role})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
