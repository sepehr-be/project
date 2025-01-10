package router

import (
	"apiTest/repository/cache"
	"apiTest/server/handler"
	"net/http"
)
func ReservationRoots() {

	repo := cache.NewReservationRepository()
    Handlers := handler.NewHandlers(repo)

	http.HandleFunc("/", Handlers.WelcomHandler)
	http.HandleFunc("/reserve", Handlers.ReservationHandler)
	http.HandleFunc("/get/reserve", Handlers.SingleReservation)
}
