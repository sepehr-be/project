package main

import (
	"goApi/config"
	"goApi/repository/db"
	"goApi/server"
	"log"
)

func main() {
	database.Main()
	cfg,err := config.LoadConfig("./")
	if err != nil {
		log.Fatalf("Failed to load config %v " , err)
	}
	server.Server(cfg.Server.Port)
}