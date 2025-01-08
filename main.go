package main

import (
	"apiTest/config"
	"apiTest/server"
	"apiTest/server/router"
)

func main() {
	router.ReservationRoots()

	port := config.ConfingApi()

	server.Server(port)
}
