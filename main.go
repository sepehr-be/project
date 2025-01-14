package main

import (
	"apiTest/config"
	"apiTest/server"
	"log"
)

func main() {

	cfg,err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("Failed to load config %v " , err)
	}

	server.Server(cfg.Server.Port)
}
