package main

import (
	"apiTest/config"
	"apiTest/server"
)

func main() {

	port := config.ConfingApi()

	server.Server(port)
}
