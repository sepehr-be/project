package router

import (
	"fmt"
	"goApi/repository/cache"
	"goApi/server/handler"
	"goApi/server/router/wrapper"
	"net/http"
)

var Mux = http.NewServeMux()

func ReservationRoots() {
	repo := cache.NewReservationRepository()
	Handlers := handler.NewHandlers(repo)

	Mux.HandleFunc("/", WelcomHandler)
	Mux.HandleFunc("/search", handler.SearchReservation)
	Mux.HandleFunc("/get/reserve", Handlers.SingleReservation)
	Mux.Handle("/reserve", ReservationNewRoots())

}

func WelcomHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "Please send data to '/reserve' for reservation.")
}


// In package router
func ReservationNewRoots() http.Handler {
	repo := cache.NewReservationRepository()
	Handlers := handler.NewHandlers(repo)

	r := routewrapper.NewRoutr()
	r.Get("/reserve", Handlers.GetReservations)
	r.Post("/reserve", Handlers.CreateReservation)
	r.Put("/reserve", Handlers.UpdateReservation)
	r.Delete("/reserve", Handlers.DeleteReservation)
	return r
}
