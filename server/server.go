package server

import (
	"apiTest/server/router"
	"fmt"
	"log"
	"net/http"
)

func Server(port string) {
	
	router.ReservationRoots()
	
	fmt.Printf("Server is running on http://localhost:%s \n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
