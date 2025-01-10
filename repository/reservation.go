package repository

import "apiTest/repository/cache"


type ReservationInterface interface {
	    Create(reservation cache.Reservation) error
	    FindById(id int) (cache.Reservation, error)
	    Get() ([]cache.Reservation, error)
	    Update(id int, reservation cache.Reservation) error
	    Delete(id int) error
	}

	
var repoInstance *cache.ReservationRepository


func GetReservationRepository() *cache.ReservationRepository {
    if repoInstance == nil {
        repoInstance = &cache.ReservationRepository{
            Data: make(map[int]cache.Reservation),
        }
    }
    return repoInstance
}
