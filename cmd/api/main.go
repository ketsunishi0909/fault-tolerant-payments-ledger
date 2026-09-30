package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

type healthResponse struct {
	Status    string    `json:"status"`
	Service   string    `json:"service"`
	Timestamp time.Time `json:"timestamp"`
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(healthResponse{
			Status:    "ok",
			Service:   "payments-ledger",
			Timestamp: time.Now().UTC(),
		})
	})

	mux.HandleFunc("POST /v1/transfers", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "transfer service not implemented yet", http.StatusNotImplemented)
	})

	addr := os.Getenv("HTTP_ADDR")
	if addr == "" {
		addr = ":8080"
	}

	log.Printf("payments-ledger API listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
