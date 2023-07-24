package routes

import (
	"github.com/CobaKauPikirkan/cashier/handler"
	"github.com/CobaKauPikirkan/cashier/middleware"
	"github.com/CobaKauPikirkan/cashier/repository/manager"
	"github.com/CobaKauPikirkan/cashier/service/manager_service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ManagerRoutes(incomingRoutes *gin.Engine, db *gorm.DB) {
	repository := manager.NewManagerRepository(db)
	service := manager_service.NewManagerService(repository)
	handler := handler.NewManagerHandler(service)

	routes := incomingRoutes.Group("/manager")
	routes.Use(middleware.Authorization([]int{2}))
	routes.GET("/findalltransaksi", handler.FindAllTransaksiManager)
	routes.GET("/findbykasir/:id", handler.FindByIdKasir)
	routes.GET("/findbytanggal/search", handler.FindByTgl)
	routes.GET("/favouritproductlist", handler.FindMostFavouriteProduct)
}