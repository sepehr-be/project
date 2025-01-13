package handler

import (
	"apiTest/repository"
	"apiTest/repository/cache"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type ReservationHandler struct {
	Repo repository.ReservationInterface
}

func (h ReservationHandler) ReservationMetodHandler(w http.ResponseWriter, r *http.Request) {

	handlers := NewHandlers(h.Repo)

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


func (h ReservationHandler) SingleReservation(w http.ResponseWriter, r *http.Request) {

	handlers := NewHandlers(h.Repo)

	switch r.Method {
	case http.MethodGet:
		handlers.FindById(w, r)
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func NewHandlers(repo repository.ReservationInterface) *ReservationHandler {
	return &ReservationHandler{Repo: repo}
}

func (h *ReservationHandler) GetReservations(w http.ResponseWriter, r *http.Request) {
	reservation, err := h.Repo.Get()
	if err != nil {
		http.Error(w, "Error retrieving reservations", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(reservation); err != nil {
		http.Error(w, "Error encoding reservations", http.StatusInternalServerError)
		return
	}
}

func (h *ReservationHandler) CreateReservation(w http.ResponseWriter, r *http.Request) {
	var reservation cache.Reservation
	err := json.NewDecoder(r.Body).Decode(&reservation)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := h.Repo.Create(reservation); err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Reservation with id %d created", reservation.NationalID)

}

func (h *ReservationHandler) UpdateReservation(w http.ResponseWriter, r *http.Request) {

	var reservation cache.Reservation
	if err := json.NewDecoder(r.Body).Decode(&reservation); err != nil {
		http.Error(w, "Invalid Input", http.StatusBadRequest)
		return
	}
	num := r.URL.Query().Get("national_id")
	id, err := strconv.Atoi(num)
	if err != nil {
		return
	}

	if err := h.Repo.Update(id, reservation); err != nil {
		http.Error(w, "Invalid Input", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "reservation with id %d updated", id)

}

func (h *ReservationHandler) DeleteReservation(w http.ResponseWriter, r *http.Request) {
	num := r.URL.Query().Get("national_id")
	id, err := strconv.Atoi(num)
	if err != nil {
		return
	}
	if err := h.Repo.Delete(id); err != nil {
		http.Error(w, "Invalid Inpur", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "reservation with id %d deleted", id)
}

func (h *ReservationHandler) FindById(w http.ResponseWriter, r *http.Request) {
	num := r.URL.Query().Get("national_id")
	id, err := strconv.Atoi(num)
	if err != nil {
		return
	}
	reservation, err := h.Repo.FindById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(reservation); err != nil {
		http.Error(w, "Error encoding user", http.StatusInternalServerError)
		return
	}
}
