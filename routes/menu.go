package routes

import (
	"github.com/CobaKauPikirkan/cashier/handler"
	"github.com/CobaKauPikirkan/cashier/middleware"
	"github.com/CobaKauPikirkan/cashier/repository/menu"
	"github.com/CobaKauPikirkan/cashier/service/menu_service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func MenuRoutes(incomingRoutes *gin.Engine, db *gorm.DB) {
	repository := menu.NewMenuRepository(db)
	service := menu_service.NewMenuService(repository)
	handler := handler.NewMenuHandler(service)

	routes := incomingRoutes.Group("/menu")
	routes.Use(middleware.Authorization([]int{1}))
	routes.POST("/", handler.CreateMenu)
	routes.GET("/", handler.GetAllMenu)
	routes.GET("/:id", handler.GetMenuById)
	routes.PUT("/:id", handler.UpdateMenu)
	routes.DELETE("/:id", handler.DeleteMenu)
}