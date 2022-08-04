package main

import (
	"log"
	"microservice-go/server"
	"net/http"
	"os"
)

const message = "Hello World"

var (
	CertFile    = os.Getenv("CERT_FILE")
	KeyFile     = os.Getenv("KEY_FILE")
	ServiceAddr = os.Getenv("SERVICE_ADDR")
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(message))
		if err != nil {
			log.Fatalf("error while sending response: %v", err)
		}
	})

	srv := server.New(mux, ServiceAddr)

	log.Println(srv.ListenAndServeTLS(CertFile, KeyFile))
}
