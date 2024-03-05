package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/luuisavelino/rinha-de-backend-2024-q1/internal/api/controllers"
	"github.com/luuisavelino/rinha-de-backend-2024-q1/internal/api/middleware"
	"github.com/luuisavelino/rinha-de-backend-2024-q1/internal/api/models/repository"
	"github.com/luuisavelino/rinha-de-backend-2024-q1/internal/api/models/service"
	"github.com/luuisavelino/rinha-de-backend-2024-q1/internal/api/routes"
	"github.com/luuisavelino/rinha-de-backend-2024-q1/pkg/database"
	"github.com/luuisavelino/rinha-de-backend-2024-q1/pkg/utils/env"
)

func main() {
	dbConfig := database.DBConfig{
		Host:     env.POSTGRES_HOST,
		Port:     env.POSTGRES_PORT,
		Dbname:   env.POSTGRES_DB_NAME,
		User:     env.POSTGRES_USER,
		Password: env.POSTGRES_PASSWORD,
	}
	db := database.NewDatabase("postgres", dbConfig)
	postgresConn, err := db.NewConnection()
	if err != nil {
		log.Fatal(err)
	}

	bankRepository := repository.NewBankRepository(postgresConn)
	service := service.NewBankServiceInterface(bankRepository)
	bankController := controllers.NewBankControllerInterface(service)

	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(middleware.LoggerMiddleware())
	router.SetTrustedProxies([]string{env.ALLOWED_ORIGINS})

	routes.InitRoutes(&router.RouterGroup, bankController)

	if err := router.Run(":" + env.API_PORT); err != nil {
		log.Fatal(err)
	}
}
