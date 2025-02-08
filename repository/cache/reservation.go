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
	return nil
}

func (r *ReservationRepository) Get() ([]Reservation, error) {
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
