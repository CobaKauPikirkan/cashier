package routes

import (
	"github.com/CobaKauPikirkan/cashier/handler"
	"github.com/CobaKauPikirkan/cashier/middleware"
	"github.com/CobaKauPikirkan/cashier/repository/kasir"
	"github.com/CobaKauPikirkan/cashier/service/kasir_service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func KasirRoutes(incomingRoutes *gin.Engine, db *gorm.DB) {
	repository := kasir.NewKasirRepository(db)
	service := kasir_service.NewKasirService(repository)
	handler := handler.NewKasirHandler(service)

	routes := incomingRoutes.Group("/kasir")
	routes.Use(middleware.Authorization([]int{3}))
	routes.GET("/", handler.FindAllTransaksi)
	routes.POST("/transaksi", handler.CreateTransaksi)
	routes.PUT("/transaksi/:id", handler.UpdateTransaksi)
	routes.GET("/transaksi/:id", handler.FindTransaksiByID)
	routes.GET("/meja/avaible",handler.FindMejaAvaible)
	routes.PUT("/meja/:id",handler.UpdateMejaAvaible)
}