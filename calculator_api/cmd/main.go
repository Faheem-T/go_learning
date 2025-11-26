package main

import (
	"log"
	"net/http"
)

func main() {
	router := http.NewServeMux()

	router.HandleFunc("POST /add", HandleAdd)
	router.HandleFunc("POST /subtract", HandleSubtract)
	router.HandleFunc("POST /multiply", HandleMultiply)
	router.HandleFunc("POST /divide", HandleDivide)

	server := http.Server{
		Addr:    ":3000",
		Handler: router,
	}

	log.Println("Listening on port :3000")

	server.ListenAndServe()
}
