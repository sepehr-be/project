package router

import (
	"apiTest/server/service"
	"net/http"
)

func ReservationRoots() {
	http.HandleFunc("/", service.WelcomHandler)
	http.HandleFunc("/reserve", service.ReservationHandler)
	http.HandleFunc("/get/reserve", service.SingleReservation)
}
