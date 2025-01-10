package handler

import (
	"apiTest/repository"
	"apiTest/repository/cache"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Handlers struct {
	Repo repository.ReservationInterface
}


func NewHandlers(repo repository.ReservationInterface) *Handlers {
	return &Handlers{Repo: repo}
}

func (h *Handlers) GetReservations(w http.ResponseWriter, r *http.Request) {
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

func (h *Handlers) CreateReservation(w http.ResponseWriter, r *http.Request) {
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


func (h *Handlers) UpdateReservation(w http.ResponseWriter, r *http.Request) {

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
	fmt.Fprintf(w, "reservation with id %d updated",id)

}


func (h *Handlers) DeleteReservation(w http.ResponseWriter, r *http.Request) {
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


func (h *Handlers) FindById(w http.ResponseWriter, r *http.Request) {
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
