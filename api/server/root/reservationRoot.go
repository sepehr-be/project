package root

import (
	"apiTest/server/rootHandler"
	"net/http"
)

func ReservationRoots() {
	http.HandleFunc("/", rootHandler.WelcomHandler)
	http.HandleFunc("/reserve", rootHandler.ReservationHandler)
}
