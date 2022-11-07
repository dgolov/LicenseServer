package main

import (
	"github.com/dgolov/LicenseServer/pkg/handler"
	"log"
	"net/http"
)

// Main

func main() {
	handlers := new(handler.Handler)
	http.HandleFunc("/check", handlers.CheckHandler)
	http.HandleFunc("/test", handlers.TestHandler)

	log.Println("Start server on port 8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
