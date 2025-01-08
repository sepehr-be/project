package router

import (
	"apiTest/server/service"
	"net/http"
)

func ReservationRoots() {
	http.HandleFunc("/", service.WelcomHandler)
	http.HandleFunc("/reserve", service.ReservationHandler)
}
