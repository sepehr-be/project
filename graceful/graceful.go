package graceful

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"
	"time"
)

const ShutdownTimeuot = time.Second * 5

func GracefulShutdown(server *http.Server, db *sql.DB) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop
	log.Println("Shuting down server")

	ctx, cancel := context.WithTimeout(context.Background(), ShutdownTimeuot)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown:%v", err)
	}
	log.Println("Server exited cleanly")
	if err := db.Close(); err != nil {
		log.Fatalf("databaase forced to shutdown:%v", err)
	}
	log.Println("database exited cleanly")
}

func RecoverMidleware(server *http.Server, db *sql.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Recovered from http panic: %v", err)
				debug.PrintStack()
				go func() {
					time.Sleep(ShutdownTimeuot)

					ctx, cancel := context.WithTimeout(context.Background(), ShutdownTimeuot)
					defer cancel()

					if err := server.Shutdown(ctx); err != nil {
						log.Fatalf("Server forced to shutdown: %v", err)
					}
					log.Println("Server exited cleanly after http panic")
					if err := db.Close(); err != nil {
						log.Fatalf("database forced to shutdown: %v", err)
					}
					log.Println("database exited cleanly after http panic")
				}()
				http.Error(res, "Internal server error", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(res, req)
	})
}