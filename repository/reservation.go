package repository

import "goApi/repository/cache"

type ReservationInterface interface {
	Create(reservation cache.Reservation) error
	FindById(id int) (cache.Reservation, error)
	Get() ([]cache.Reservation, error)
	Update(id int, reservation cache.Reservation) error
	Delete(id int) error
}
