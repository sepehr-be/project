package router

import (
	"apiTest/server/handler"
	"net/http"
)
func ReservationRoots() {
	http.HandleFunc("/", handler.WelcomHandler)
	http.HandleFunc("/reserve", handler.ReservationHandler)
	http.HandleFunc("/get/reserve", handler.SingleReservation)
}
