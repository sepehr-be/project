package handler

import (
	"apiTest/data"
	"encoding/json"
	"net/http"
)

func Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data.ReservationList)
}