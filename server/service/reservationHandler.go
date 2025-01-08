package service

import (
	"apiTest/server/handler"
	"net/http"
)

func ReservationHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		handler.Post(w, r)

	case http.MethodGet:
		handler.Get(w, r)

	case http.MethodDelete:
		handler.Delete(w, r)

	case http.MethodPut:
		handler.Update(w, r)

	}
}
