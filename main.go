package main

import (
	"github.com/joho/godotenv"
	"github.com/jpbmdev/payment-api/config"
	"github.com/jpbmdev/payment-api/core"
)

func main() {
	//Load enviroment variables
	godotenv.Load(".env")
	config.LoadConfig()

	//Create server instance
	serverInstace := core.NewServerInstance()

	//Start configured server
	serverInstace.InitServer()
}
