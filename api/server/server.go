package server

import (
	"fmt"
	"log"
	"net/http"
)

func Server(port string) {
	fmt.Printf("Server is running on http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
