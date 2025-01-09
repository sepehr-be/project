// package cache

// import (
// 	"errors"
// 	"fmt"
// )

// type Reservation struct {
// 	FirstName   string `json:"first_name"`
// 	LastName    string `json:"last_name"`
// 	Email       string `json:"email"`
// 	NationalID  int `json:"national_id"`
// 	TicketCount int    `json:"ticket_count"`
// }

// type ReservationRepository struct {
// 	data map[int]Reservation
// }

// func NewReservationRepository() *ReservationRepository {
// 	return &ReservationRepository{
// 		data: make(map[int]Reservation),
// 	}
// }

// func (r *ReservationRepository) Create(reservation Reservation) error {
// 	if _, exists := r.data[reservation.NationalID]; exists {
// 		return fmt.Errorf("Reservation with national id %d already exists",reservation.NationalID)
// 	}
// 	r.data[reservation.NationalID] = reservation
// 	return nil
// }

// func (r *ReservationRepository) Get() ([]Reservation,error) {
// 	if len(r.data) == 0 {
// 		return nil ,errors.New("no reservation found")
// 	}
// 	reservations := make([]Reservation,0)
// 	for _,res := range r.data {
// 		reservations = append(reservations, res)
// 	}
// 	return reservations, nil
// }

// func (r *ReservationRepository) FindById(id int) (Reservation,error) {
// 	reservation,exists := r.data[id]
// 	if !exists {
// 		return Reservation{} , fmt.Errorf("reservation with nation id %d not found",id)
// 	}
// 	return reservation,nil

// }

// func (r *ReservationRepository) Delete(id int) error {
// 	_,exists := r.data[id]
// 	if !exists {
// 		return fmt.Errorf("reservation with nation id %d not found",id)
// 	}
// 	delete(r.data,id)
// 	return nil
// }

// func (r *ReservationRepository) Update(reservation Reservation, id int) error {
// 	_,exists := r.data[id]
// 	if !exists {
// 		return fmt.Errorf("reservation with nation id %d not found",id)
// 	}
// 	r.data[id] = reservation
// 	return nil
// }


// second code

package cache

import "fmt"

type Reservation struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	NationalID  int    `json:"national_id"`
	TicketCount int    `json:"ticket_count"`
}

type ReservationRepository struct {
	Data map[int]Reservation
}

func NewReservationRepository() *ReservationRepository {
	return &ReservationRepository{
		Data: make(map[int]Reservation),
	}
}

func (r *ReservationRepository) Create(reservation Reservation) error {
	if _, exists := r.Data[reservation.NationalID]; exists {
		return fmt.Errorf("Reservation with national id %d already exists", reservation.NationalID)
	}
	r.Data[reservation.NationalID] = reservation
	fmt.Println(reservation,r.Data,r.Data[reservation.NationalID])
	return nil
}

func (r *ReservationRepository) Get() ([]Reservation, error) {
	fmt.Println("hello")
	if len(r.Data) == 0 {
		return nil, fmt.Errorf("no reservation found")
	}
	reservations := make([]Reservation, 0)
	for _, res := range r.Data {
		reservations = append(reservations, res)
	}
	return reservations, nil
}

func (r *ReservationRepository) FindById(id int) (Reservation, error) {
	reservation, exists := r.Data[id]
	if !exists {
		return Reservation{}, fmt.Errorf("reservation with national id %d not found", id)
	}
	return reservation, nil
}

func (r *ReservationRepository) Delete(id int) error {
	_, exists := r.Data[id]
	if !exists {
		return fmt.Errorf("reservation with national id %d not found", id)
	}
	delete(r.Data, id)
	return nil
}

func (r *ReservationRepository) Update(id int, reservation Reservation) error {
	_, exists := r.Data[id]
	if !exists {
		return fmt.Errorf("reservation with national id %d not found", id)
	}
	r.Data[id] = reservation
	return nil
}
