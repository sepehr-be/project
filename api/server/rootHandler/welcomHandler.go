package rootHandler

import (
	"fmt"
	"net/http"
)

func WelcomHandler(response http.ResponseWriter, request *http.Request) {

	fmt.Fprintf(response, "Please send data to '/reserve' for reservation.")
}
