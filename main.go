package main

import (
	"auth-rate-limit-api/internal/auth"
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/register", auth.Register)

	mux.HandleFunc("/", auth.HomePage)

	log.Println("server running on :8081")
	log.Fatal(http.ListenAndServe(":8081", mux))
}
