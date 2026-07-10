package main

import (
	"log"
	"net/http"
	"time"

	"github.com/eapra/eapra/request-plane/ai-gateway/internal/server"
)

func main() {
	srv := &server.Server{}
	log.Println("EAPRA gateway listening on :8080")
	httpSrv := &http.Server{
		Addr:              ":8080",
		Handler:           srv.Routes(),
		ReadHeaderTimeout: 5 * time.Second,
	}
	log.Fatal(httpSrv.ListenAndServe())
}