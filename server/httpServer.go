package server

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"runners/controllers"
	"runners/repositories"
	"runners/services"
)

type HttpServer struct {
	config            *viper.Viper
	router            *gin.Engine
	runnersController *controllers.RunnersController
	resultsController *controllers.ResultsController
}

func InitHttpServer(config *viper.Viper, dbHandler *sql.DB) HttpServer {
	runnersRepository := repositories.NewRunnersRepository(dbHandler)
	resultsRepository := repositories.NewResultsRepository(dbHandler)
	runnersService := services.NewRunnersService(runnersRepository, resultsRepository)
	resultsService := services.NewResultsService(runnersRepository, resultsRepository)
	runnersController := controllers.NewRunnersController(runnersService)
	resultsController := controllers.NewResultsController(resultsService)

	router := gin.Default()
	router.POST("/runner", runnersController.CreateRunner)
	router.PUT("/runner", runnersController.UpdateRunner)
	router.DELETE("/runner/:id", runnersController.DeleteRunner)
	router.GET("/runner/:id", runnersController.GetRunner)
	router.GET("/runner/", runnersController.GetRunnersBatch)

	router.POST("/result", resultsController.CreateResult)
	router.DELETE("/result/:id", resultsController.DeleteResult)

	return HttpServer{
		config:            config,
		router:            router,
		runnersController: runnersController,
		resultsController: resultsController,
	}
}

func (hs HttpServer) Start() {
	err := hs.router.Run(hs.config.GetString("http.server_address"))
	if err != nil {
		log.Fatal("Error while starting HTTP Server: %v", err)
	}
}
