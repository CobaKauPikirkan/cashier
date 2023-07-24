package main

import (
	"log"
	"os"
	"time"

	"github.com/CobaKauPikirkan/cashier/initializer"
	"github.com/CobaKauPikirkan/cashier/middleware"
	"github.com/CobaKauPikirkan/cashier/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Init() {

}

func main() {
	db, err := initializer.ConnectDatabase()
	if err != nil {
		log.Fatal(err)
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.MaxMultipartMemory = 50 << 20 // 50 MiB

	routes.UserRoutes(router, db)

	router.Use(middleware.RBACMiddleware())

	routes.MejaRoutes(router, db)
	routes.MenuRoutes(router, db)
	routes.KasirRoutes(router, db)
	routes.ManagerRoutes(router, db)
	routes.AdminRoutes(router, db)

	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	log.Fatal(router.Run(":" + port))
}
