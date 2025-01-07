package handler

import (
	"apiTest/data"
	"apiTest/verification"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
)

func PostHandler(w http.ResponseWriter, r *http.Request) {
	var res data.ReservationData

	err := json.NewDecoder(r.Body).Decode(&res)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	fmt.Println(res, reflect.TypeOf(res), reflect.TypeOf(res.FirstName), reflect.TypeOf(res.LastName), reflect.TypeOf(res.Email), reflect.TypeOf(res.NationalID), reflect.TypeOf(res.TicketCount))

	errors := verification.VerificationData(res.FirstName, res.LastName, res.Email, res.NationalID, res.TicketCount)
	if len(errors) > 0 {
		fmt.Println(errors)
		http.Error(w, fmt.Sprintf("Validation errors: %v", errors), http.StatusBadRequest)
		return
	}

	saveReservation(res)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Reservation successful"))

}

func saveReservation(res data.ReservationData) {
	data.ReservationList = append(data.ReservationList, res)
}
