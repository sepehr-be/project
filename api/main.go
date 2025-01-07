package main

import (
	"apiTest/config"
	"apiTest/server"
	"apiTest/server/root"
)

func main() {
	root.ReservationRoots()

	port := config.ConfingApi()

	server.Server(port)
}
