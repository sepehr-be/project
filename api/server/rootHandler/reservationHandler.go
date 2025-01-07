package rootHandler

import (
	"apiTest/server/rootHandler/handler"
	"net/http"
)

func ReservationHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		handler.PostHandler(w, r)

	case http.MethodGet:
		handler.GetHandler(w, r)

	case http.MethodDelete:
		handler.DeletHandler(w, r)

	case http.MethodPut:
		handler.UpdateHandler(w, r)

	}
}
