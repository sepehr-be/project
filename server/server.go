package server

import (
 "apiTest/server/router"
 "context"
 "fmt"
 "log"
 "net/http"
 "os"
 "os/signal"
 "syscall"
 "time"
)

func Server(port string) {
 router.ReservationRoots()

 fmt.Printf("Server is running on http://localhost:%s\n", port)

 server := &http.Server{Addr: ":" + port}

 
 stop := make(chan os.Signal, 1)
 signal.Notify(stop, os.Interrupt, syscall.SIGTERM)


 go func() {
  if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
   log.Fatalf("Server error: %v", err)
  }
 }()


 <-stop
 fmt.Println("\nShutting down server gracefully...")


 ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
 defer cancel()


 if err := server.Shutdown(ctx); err != nil {
  log.Fatalf("Error shutting down server: %v", err)
 }

 fmt.Println("Server stopped gracefully.")
}