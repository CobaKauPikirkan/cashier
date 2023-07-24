package routes

import (
	"github.com/CobaKauPikirkan/cashier/handler"
	"github.com/CobaKauPikirkan/cashier/middleware"
	"github.com/CobaKauPikirkan/cashier/repository/meja"
	"github.com/CobaKauPikirkan/cashier/service/meja_service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func MejaRoutes(incomingRoutes *gin.Engine, db *gorm.DB) {
	repository := meja.NewMejaRepository(db)
	service := meja_service.NewMejaService(repository)
	handler := handler.NewMejaHandler(service)

	routes := incomingRoutes.Group("/meja")
	routes.Use(middleware.Authorization([]int{1}))
	routes.POST("/", handler.CreateMeja)
	routes.GET("/", handler.GetAllMeja)
	routes.GET("/:id", handler.GetMejaById)
	routes.PUT("/:id", handler.UpdateMeja)
	routes.DELETE("/:id", handler.DeleteMeja)
}