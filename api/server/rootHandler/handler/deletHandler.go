package handler

import (
	"apiTest/data"
	"fmt"
	"net/http"
)

func deleteReservation(nationalID string) {
	for index, res := range data.ReservationList {
		if res.NationalID == nationalID {
			// حذف رزرو با استفاده از slicing
			data.ReservationList = append(data.ReservationList[:index], data.ReservationList[index+1:]...)
			fmt.Println("Reservation deleted successfully.")
			return
		}
	}
}

func DeletHandler(w http.ResponseWriter, r *http.Request) {

	nationalID := r.URL.Query().Get("national_id")
	fmt.Println(nationalID, r.URL)
	if nationalID == "" {
		http.Error(w, "Missing National ID", http.StatusBadRequest)
		return
	}
	deleteReservation(nationalID)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Reservation deleted"))

}
