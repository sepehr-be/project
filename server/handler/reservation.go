// package handler

// import (
// 	"apiTest/repository"
// 	"apiTest/repository/cache"
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"strconv"
// )

// type Handlers struct {
// 	Repo repository.ReservationInterface
// }

// func NewHandlers(repo repository.ReservationInterface) *Handlers {

// 	return &Handlers{Repo: repo}
// }

// func (h *Handlers) GetReservations(w http.ResponseWriter, r *http.Request) {

// 	reservation, err := h.Repo.Get()
// 	fmt.Println(reservation)
// 	if err != nil {
// 		fmt.Println(err)
// 		http.Error(w, "error retrieving reservations", http.StatusInternalServerError)
// 		return
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	if err := json.NewEncoder(w).Encode(reservation); err != nil {
// 		http.Error(w, "error encoding reservations", http.StatusInternalServerError)
// 		return
// 	}
// }

// func (h *Handlers) FindById(w http.ResponseWriter, r *http.Request) {
// 	num := r.URL.Query().Get("national_id")
// 	id, err := strconv.Atoi(num)
// 	if err != nil {
// 		fmt.Println("error:", err)
// 	}
// 	fmt.Println(id)
// 	reservation, err := h.Repo.FindById(id)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusNotFound)
// 		return
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	if err := json.NewEncoder(w).Encode(reservation); err != nil {
// 		http.Error(w, "Error encoding user", http.StatusInternalServerError)
// 		return
// 	}
// 	// json.NewEncoder(w).Encode(data.ReservationList)
// }

// func (h *Handlers) CreateReservation(w http.ResponseWriter, r *http.Request) {
// 	var reservation cache.Reservation
// 	err := json.NewDecoder(r.Body).Decode(&reservation)
// 	if err != nil {
// 		http.Error(w, "Invalid input", http.StatusBadRequest)
// 		return
// 	}
// 	// fmt.Println(data, reflect.TypeOf(data), reflect.TypeOf(data.FirstName), reflect.TypeOf(data.LastName), reflect.TypeOf(data.Email), reflect.TypeOf(data.NationalID), reflect.TypeOf(data.TicketCount))

// 	// errors := verification.VerificationData(data.FirstName, data.LastName, data.Email, data.NationalID, data.TicketCount)
// 	// if len(errors) > 0 {
// 	// 	fmt.Println(errors)
// 	// 	http.Error(w, fmt.Sprintf("Validation errors: %v", errors), http.StatusBadRequest)
// 	// 	return
// 	// }
// 	if err := h.Repo.Create(reservation); err != nil {
// 		http.Error(w, err.Error(), http.StatusConflict)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte("Reservation successful"))
// 	fmt.Fprintf(w, "Reservation with id %d created", reservation.NationalID)

// }

// // func saveReservation(res data.ReservationData) {
// // 	data.ReservationList = append(data.ReservationList, res)
// // }

// // func update(nationalID string, updatedRes data.ReservationData) bool {
// // 	for index, res := range data.ReservationList {
// // 		if res.NationalID == nationalID {
// // 			// بروزرسانی رزرو
// // 			data.ReservationList[index] = updatedRes
// // 			fmt.Println("Reservation updated successfully.")
// // 			return true
// // 		}
// // 	}
// // 	return false
// // }

// func (h *Handlers) UpdateReservation(w http.ResponseWriter, r *http.Request) {

// 	var reservation cache.Reservation
// 	if err := json.NewDecoder(r.Body).Decode(&reservation); err != nil {
// 		http.Error(w, "Invalid Input", http.StatusBadRequest)
// 		return
// 	}
// 	num := r.URL.Query().Get("national_id")
// 	id, err := strconv.Atoi(num)
// 	if err != nil {
// 		fmt.Println("error:", err)
// 	}
// 	fmt.Println(id)

// 	if err := h.Repo.Update(id, reservation); err != nil {
// 		http.Error(w, "Invalid Input", http.StatusNotFound)
// 		return
// 	}
// 	w.WriteHeader(http.StatusOK)
// 	fmt.Fprintf(w, "reservation with id %d updated",id)

// 	// nationalID := r.URL.Query().Get("national_id")
// 	// fmt.Println(nationalID)
// 	// if nationalID == "" {
// 	// 	http.Error(w, "Missing National ID", http.StatusBadRequest)
// 	// 	return
// 	// }

// 	// var updatedRes data.ReservationData
// 	// err := json.NewDecoder(r.Body).Decode(&updatedRes)
// 	// if err != nil {
// 	// 	http.Error(w, "Invalid input", http.StatusBadRequest)
// 	// 	return
// 	// }

// 	// if !update(nationalID, updatedRes) {
// 	// 	http.Error(w, "Reservation not found", http.StatusNotFound)
// 	// 	return
// 	// }

// 	// w.WriteHeader(http.StatusOK)
// 	// w.Write([]byte("Reservation updated"))
// }

// // func delete(nationalID string) {
// // 	for index, res := range data.ReservationList {
// // 		if res.NationalID == nationalID {
// // 			data.ReservationList = append(data.ReservationList[:index], data.ReservationList[index+1:]...)
// // 			fmt.Println("Reservation deleted successfully.")
// // 			return
// // 		}
// // 	}
// // }

// func (h *Handlers) DeleteReservation(w http.ResponseWriter, r *http.Request) {
// 	num := r.URL.Query().Get("national_id")
// 	id, err := strconv.Atoi(num)
// 	if err != nil {
// 		fmt.Println("error:", err)
// 	}
// 	fmt.Println(id)
// 	if err := h.Repo.Delete(id); err != nil {
// 		http.Error(w, "Invalid Inpur", http.StatusNotFound)
// 		return
// 	}
// 	w.WriteHeader(http.StatusOK)
// 	fmt.Fprintf(w, "reservation with id %d deleted", id)

// 	// nationalID := r.URL.Query().Get("national_id")
// 	// fmt.Println(nationalID, r.URL)
// 	// if nationalID == "" {
// 	// 	http.Error(w, "Missing National ID", http.StatusBadRequest)
// 	// 	return
// 	// }
// 	// delete(nationalID)
// 	// w.WriteHeader(http.StatusOK)
// 	// w.Write([]byte("Reservation deleted"))

// }

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
		fmt.Println("error:", err)
	}
	fmt.Println(id)

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
		fmt.Println("error:", err)
	}
	fmt.Println(id)
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
		fmt.Println("error:", err)
	}
	fmt.Println(id)
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
	// json.NewEncoder(w).Encode(data.ReservationList)
}
