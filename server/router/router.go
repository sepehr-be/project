package router

import (
  "apiTest/repository/cache"
  "apiTest/server/handler"
  "apiTest/server/router/app"
  "fmt"
  "net/http"
)

func ReservationRoots() {

  repo := cache.NewReservationRepository()
  Handlers := handler.NewHandlers(repo)

  http.HandleFunc("/", WelcomHandler)
  // http.HandleFunc("/reserve", Handlers.ReservationHandler)
  http.Handle("/reserve", ReservationNewRoots())

  http.HandleFunc("/get/reserve", Handlers.SingleReservation)
}

func WelcomHandler(response http.ResponseWriter, request *http.Request) {
  fmt.Fprintf(response, "Please send data to '/reserve' for reservation.")
}

// In package router
func ReservationNewRoots() http.Handler {
  repo := cache.NewReservationRepository()
  Handlers := handler.NewHandlers(repo)

  r := app.NewRouter()
  r.Get("/reserve", Handlers.GetReservations)
  r.Post("/reserve", Handlers.CreateReservation)
  r.Put("/reserve", Handlers.UpdateReservation)
  r.Delete("/reserve", Handlers.DeleteReservation)

  return r
}