// package service

// import (
// 	"apiTest/repository/cache"
// 	"apiTest/server/handler"
// 	"fmt"
// 	"net/http"
// )

// func ReservationHandler(w http.ResponseWriter, r *http.Request) {

	
// 	repo := cache.NewReservationRepository()
// 	handlers := handler.NewHandlers(repo)

// 	switch r.Method {
// 	case http.MethodPost:
// 		 handlers.CreateReservation(w , r)
// 		// // handler.CreateReservation(w, r)

// 	case http.MethodGet:
// 		handlers.GetReservations(w , r)
// 		// handler.GetReservations(w, r)

// 	case http.MethodDelete:
// 		handlers.DeleteReservation(w , r)
// 		// handler.DeleteReservation(w, r)

// 	case http.MethodPut:
// 		handlers.UpdateReservation(w , r)
// 		// handler.UpdateReservation(w, r)

// 	}
// }

// func WelcomHandler(response http.ResponseWriter, request *http.Request) {

// 	fmt.Fprintf(response, "Please send data to '/reserve' for reservation.")
// }




// package service

// import (
//     "apiTest/repository"
//     "apiTest/server/handler"
//     "fmt"
//     "net/http"
// )

// func ReservationHandler(w http.ResponseWriter, r *http.Request) {
//     repo := repository.NewReservationRepository()  // از اینجا درست استفاده کنید
//     handlers := handler.NewHandlers(repo)

//     switch r.Method {
//     case http.MethodPost:
//         handlers.CreateReservation(w, r)
//     case http.MethodGet:
//         handlers.GetReservations(w, r)
//     case http.MethodDelete:
//         handlers.DeleteReservation(w, r)
//     case http.MethodPut:
//         handlers.UpdateReservation(w, r)
//     default:
//         http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
//     }
// }

// func WelcomHandler(response http.ResponseWriter, request *http.Request) {
//     fmt.Fprintf(response, "Please send data to '/reserve' for reservation.")
// }


package service

import (
    "apiTest/repository"
    "apiTest/server/handler"
    "fmt"
    "net/http"
)

func ReservationHandler(w http.ResponseWriter, r *http.Request) {
    repo := repository.GetReservationRepository()  // از اینجا استفاده کنید
    handlers := handler.NewHandlers(repo)

    switch r.Method {
    case http.MethodPost:
        handlers.CreateReservation(w, r)
    case http.MethodGet:
        handlers.GetReservations(w, r)
    case http.MethodDelete:
        handlers.DeleteReservation(w, r)
    case http.MethodPut:
        handlers.UpdateReservation(w, r)
    default:
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
    }
}

func WelcomHandler(response http.ResponseWriter, request *http.Request) {
    fmt.Fprintf(response, "Please send data to '/reserve' for reservation.")
}

func SingleReservation(w http.ResponseWriter, r *http.Request) {
    repo := repository.GetReservationRepository()  // از اینجا استفاده کنید
    handlers := handler.NewHandlers(repo)

    switch r.Method {
    case http.MethodGet:
        handlers.FindById(w, r)
    default:
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
    }
}
