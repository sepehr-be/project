package router

import (
	"apiTest/repository/cache"
	"apiTest/server/handler"
	"fmt"
	"net/http"
)
func ReservationRoots() {

	repo := cache.NewReservationRepository()
    Handlers := handler.NewHandlers(repo)

	http.HandleFunc("/", WelcomHandler)
	http.HandleFunc("/reserve", Handlers.ReservationHandler)
	http.HandleFunc("/get/reserve", Handlers.SingleReservation)
}

func WelcomHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "Please send data to '/reserve' for reservation.")
}