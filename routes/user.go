package routes

import (
	"github.com/CobaKauPikirkan/cashier/handler"
	"github.com/CobaKauPikirkan/cashier/repository/user"
	"github.com/CobaKauPikirkan/cashier/service/user_service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRoutes(incomingRoutes *gin.Engine, db *gorm.DB) {
	repository := user.NewUserRepository(db)
	service := user_service.NewUserService(repository)
	handler := handler.NewUserHandler(service)
	routes := incomingRoutes.Group("/user")
	routes.POST("/signup", handler.SignUp)
	routes.POST("/login", handler.Login)
}