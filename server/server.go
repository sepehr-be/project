package server

import (
	"fmt"
	"goApi/graceful"
	database "goApi/repository/db"
	"goApi/server/router"
	"log"
	"net/http"
)


func Server(port string) {

	db := database.GetDB()
	router.ReservationRoots()
	server := &http.Server{Addr: ":" + port}

	handler := graceful.RecoverMidleware(server,db,router.Mux)
	server.Handler = handler

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server error: %v", err)
		}
	}()
	fmt.Printf("Server is running on http://localhost:%s\n", port)
	graceful.GracefulShutdown(server,db)
}
