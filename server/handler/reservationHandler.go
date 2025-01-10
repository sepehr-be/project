package handler

import (
	"apiTest/repository"
	"fmt"
	"net/http"
)

func ReservationHandler(w http.ResponseWriter, r *http.Request) {
    repo := repository.GetReservationRepository() 
    handlers := NewHandlers(repo)

    switch r.Method {
    case http.MethodPost:
        handlers.CreateReservation(w, r)
    case http.MethodGet:
        handlers.GetReservations(w, r)
    case http.MethodDelete:
        handlers.DeleteReservation(w, r)
    case http.MethodPut:
        handlers.UpdateReservation(w, r)
    default:
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
    }
}

func WelcomHandler(response http.ResponseWriter, request *http.Request) {
    fmt.Fprintf(response, "Please send data to '/reserve' for reservation.")
}

func SingleReservation(w http.ResponseWriter, r *http.Request) {
    repo := repository.GetReservationRepository()
    handlers := NewHandlers(repo)

    switch r.Method {
    case http.MethodGet:
        handlers.FindById(w, r)
    default:
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
    }
}
