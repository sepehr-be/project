// package repository

// import "apiTest/repository/cache"

// type ReservationInterface interface {
// 	Create(reservation cache.Reservation) error
// 	FindById(id int) (cache.Reservation, error)
// 	Get() ([]cache.Reservation, error)
// 	Update(id int, reservation cache.Reservation) error
// 	Delete(id int) error
// }



// package repository

// import "apiTest/repository/cache"

// type ReservationInterface interface {
//     Create(reservation cache.Reservation) error
//     FindById(id int) (cache.Reservation, error)
//     Get() ([]cache.Reservation, error)
//     Update(id int, reservation cache.Reservation) error
//     Delete(id int) error
// }

// // NewReservationRepository ایجاد یک نمونه از ReservationRepository و برگرداندن آن
// func NewReservationRepository() ReservationInterface {
//     return &cache.ReservationRepository{
//         Data: make(map[int]cache.Reservation),
//     }
// }



package repository

import "apiTest/repository/cache"


type ReservationInterface interface {
	    Create(reservation cache.Reservation) error
	    FindById(id int) (cache.Reservation, error)
	    Get() ([]cache.Reservation, error)
	    Update(id int, reservation cache.Reservation) error
	    Delete(id int) error
	}

// تعریف یک متغیر برای نگه‌داری یک نمونه ثابت از repository
var repoInstance *cache.ReservationRepository

// متد برای دریافت یک نمونه از ReservationRepository
func GetReservationRepository() *cache.ReservationRepository {
    if repoInstance == nil {
        repoInstance = &cache.ReservationRepository{
            Data: make(map[int]cache.Reservation),
        }
    }
    return repoInstance
}
