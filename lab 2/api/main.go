package main

import (
	"log"
	"net/http"
	"time"

	"myapp/internal/handlers"
	"myapp/internal/middleware"
)

func main() {
	mux := http.NewServeMux()

	// single path, handle methods inside
	mux.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.GetUser(w, r)
		case http.MethodPost:
			handlers.CreateUser(w, r)
		default:
			http.Error(w, "", http.StatusMethodNotAllowed)
		}
	})

	handler := middleware.AuthMiddleware(mux)

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      handler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Printf("listening on %s", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}
