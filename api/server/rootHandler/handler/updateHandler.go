package handler

import (
	"apiTest/data"
	"encoding/json"
	"fmt"
	"net/http"
)

func updateReservation(nationalID string, updatedRes data.ReservationData) bool {
	for index, res := range data.ReservationList {
		if res.NationalID == nationalID {
			// بروزرسانی رزرو
			data.ReservationList[index] = updatedRes
			fmt.Println("Reservation updated successfully.")
			return true
		}
	}
	return false
}

func UpdateHandler(w http.ResponseWriter, r *http.Request) {

	nationalID := r.URL.Query().Get("national_id")
	fmt.Println(nationalID)
	if nationalID == "" {
		http.Error(w, "Missing National ID", http.StatusBadRequest)
		return
	}

	var updatedRes data.ReservationData
	err := json.NewDecoder(r.Body).Decode(&updatedRes)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if !updateReservation(nationalID, updatedRes) {
		http.Error(w, "Reservation not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Reservation updated"))
}
