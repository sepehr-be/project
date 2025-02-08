package handler

import (
	"encoding/json"
	"fmt"
	"goApi/repository"
	"goApi/repository/cache"
	"goApi/repository/db"
	"goApi/verification"
	"net/http"
	"strconv"
)

type ReservationHandler struct {
	Repo repository.ReservationInterface
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

var remainingTicket int = 100

func NewHandlers(repo repository.ReservationInterface) *ReservationHandler {
	return &ReservationHandler{Repo: repo}
}

func (h *ReservationHandler) GetReservations(res http.ResponseWriter, req *http.Request) {
	pageSizeStr := req.URL.Query().Get("page_size")
	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		pageSize = 0
		http.Error(res, "extracting page size number error", http.StatusBadRequest)
	}
	pagestr := req.URL.Query().Get("page")
	page, err := strconv.Atoi(pagestr)
	if err != nil {
		page = 0
		http.Error(res, "extracting page number error", http.StatusBadRequest)
	}
	reservation, err := h.Repo.Get()
	if err != nil {
		http.Error(res, "reservation not found!", http.StatusNotFound)
		return
	}
	res.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(res).Encode(reservation); err != nil {
		http.Error(res, "Encoding Reservation Error From Repository", http.StatusInternalServerError)
	}

	data, err := database.GetAllReservations(pageSize, page)
	if err != nil {
		http.Error(res, "returning data error!", http.StatusBadRequest)
	}
	res.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(res).Encode(data); err != nil {
		http.Error(res, "Encoding Reservation Error From Database", http.StatusInternalServerError)
	}
}

func (h *ReservationHandler) CreateReservation(res http.ResponseWriter, req *http.Request) {
	var reservation cache.Reservation
	if err := json.NewDecoder(req.Body).Decode(&reservation); err != nil {
		http.Error(res, "Invalid input", http.StatusBadRequest)
		return
	}

	errors, correctName := verification.VerificationData(reservation.FirstName, reservation.LastName, reservation.Email, reservation.NationalID, reservation.TicketCount, remainingTicket)
	if len(errors) > 0 {
		for _, err := range errors {
			http.Error(res, err, http.StatusBadRequest)
			return
		}
	}

	reservation.FirstName = correctName.FirstName
	reservation.LastName = correctName.LastName
	remainingTicket = remainingTicket - reservation.TicketCount

	if err := h.Repo.Create(reservation); err != nil {
		http.Error(res, err.Error(), http.StatusConflict)
		return
	}

	if err := database.InsertReservation(reservation); err != nil {
		http.Error(res, err.Error(), http.StatusConflict)
		return
	}

	res.WriteHeader(http.StatusOK)
	fmt.Fprintf(res, "Congratulations \n Reservation with id %d created", reservation.NationalID)
}

func (h *ReservationHandler) UpdateReservation(res http.ResponseWriter, req *http.Request) {
	var reservation cache.Reservation
	if err := json.NewDecoder(req.Body).Decode(&reservation); err != nil {
		http.Error(res, "Invalid Input", http.StatusBadRequest)
	}
	ID := req.URL.Query().Get("national_id")
	num, err := strconv.Atoi(ID)
	if err != nil {
		fmt.Println(err)
	}
	if err := h.Repo.Update(num, reservation); err != nil {
		http.Error(res, "Invalid Input", http.StatusBadRequest)
		return
	}

	er := database.UpdateReservation(ID, reservation)
	if er != nil {
		http.Error(res, "Invalid Input", http.StatusBadRequest)
		return
	}
	res.WriteHeader(http.StatusOK)
	fmt.Fprintf(res, "reservation with id %v updated", ID)
}

func (h *ReservationHandler) DeleteReservation(res http.ResponseWriter, req *http.Request) {
	ID := req.URL.Query().Get("national_id")
	num, err := strconv.Atoi(ID)
	if err != nil {
		fmt.Println(err)
	}
	if err := h.Repo.Delete(num); err != nil {
		http.Error(res, "Invalid Input", http.StatusNotFound)
		return
	}

	er := database.DeleteReservation(ID)
	if er != nil {
		http.Error(res, "Invalid Input", http.StatusBadRequest)
		return
	}

	res.WriteHeader(http.StatusOK)
	fmt.Fprintf(res, "reservation with id %v deleted", ID)
}

func (h *ReservationHandler) FindById(res http.ResponseWriter, req *http.Request) {
	ID := req.URL.Query().Get("national_id")
	num, err := strconv.Atoi(ID)
	if err != nil {
		fmt.Println(err)
	}
	reserv, er := database.GetReservationByID(ID)
	if er != nil {
		http.Error(res, "not found", http.StatusNotFound)
		return
	}
	res.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(res).Encode(reserv); err != nil {
		http.Error(res, "Error encoding user", http.StatusInternalServerError)
		return
	}

	reservation, err := h.Repo.FindById(num)
	if err != nil {
		http.Error(res, err.Error(), http.StatusNotFound)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(res).Encode(reservation); err != nil {
		http.Error(res, "Error encoding user", http.StatusInternalServerError)
		return
	}
}

func SearchReservation(res http.ResponseWriter, req *http.Request) {
	value := req.URL.Query().Get("value")
	pageStr := req.URL.Query().Get("page")
	pageSizeStr := req.URL.Query().Get("page_size")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		http.Error(res, "extracting page number error", http.StatusBadRequest)
	}
	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		http.Error(res, "extracting page size numbrt error", http.StatusBadRequest)
	}
	reservation, err := database.SearchReservation(value, page, pageSize)
	if err != nil {
		http.Error(res, "not found", http.StatusNotFound)
	}
	if err = json.NewEncoder(res).Encode(reservation); err != nil {
		http.Error(res, "Encoding Reservation Error From Database", http.StatusInternalServerError)
	}
}
