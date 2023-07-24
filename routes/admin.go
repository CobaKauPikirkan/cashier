package routes

import (
	"github.com/CobaKauPikirkan/cashier/handler"
	"github.com/CobaKauPikirkan/cashier/middleware"
	"github.com/CobaKauPikirkan/cashier/repository/admin"
	"github.com/CobaKauPikirkan/cashier/service/admin_service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AdminRoutes(incomingRoutes *gin.Engine, db *gorm.DB )  {
	repository := admin.NewAdminRepository(db)
	service := admin_service.NewAdminService(repository)
	handler := handler.NewAdminHandler(service)

	routes := incomingRoutes.Group("/admin")
	routes.Use(middleware.Authorization([]int{1}))
	routes.GET("/user", handler.FindAllUser)
	routes.GET("/user/:id", handler.FindByUser)
	routes.PUT("/user/:id", handler.UpdateUser)
	routes.DELETE("/user/:id", handler.DeleteUser)
}