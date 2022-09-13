package main

import (
	"github.com/joho/godotenv"
	"github.com/jpbmdev/payment-api/core"
)

func main() {
	godotenv.Load(".env")

	serverInstace := core.NewServerInstance()

	serverInstace.InitServer()
}
