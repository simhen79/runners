package main

import (
	"log"

	"runners/config"
	"runners/server"
)

func main() {
	log.Println("Starting the Runners App")
	log.Println("Initializing configuration")
	config := config.InitConfig("runners")
	log.Println("Initializing database")
	dbHander := server.InitDatabase(config)
	log.Println("Initializing HTTP Server")
	httpServer := server.InitHttpServer(config, dbHander)
	httpServer.Start()
}
