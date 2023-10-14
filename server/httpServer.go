package server

import (
	"database/sql"
	"log"

	"codeid.northwind/controllers"
	"codeid.northwind/repositories"
	"codeid.northwind/services"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type HttpServer struct {
	config             *viper.Viper
	router             *gin.Engine
	// categoryController *controllers.CategoryController
	ControllerManager controllers.ControllerManager
}

func InitHttpServer(config *viper.Viper, dbHandler *sql.DB) HttpServer{
	
	repositoryManager := repositories.NewRepositoryManager(dbHandler)
	serviceManager := services.NewServiceManager(repositoryManager)
	controllerManager := controllers.NewControllerManager(serviceManager)

	// create object router nly one
	router := gin.Default()
	InitRouter(router, controllerManager)

	return HttpServer{
		config: config,
		router: router,
		ControllerManager: *controllerManager,
	}
}

// Running gin httpserver
func (hs HttpServer)Start(){
	err := hs.router.Run(hs.config.GetString("http.server_address"))
	if err!= nil {
		log.Fatalf("Error While Starting HTTP Server : %v", err)
	}
}