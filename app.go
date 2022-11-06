package main


import (
	"log"
	"net/http"
)

// Main

func main() {
	http.HandleFunc("/check", checkHandler)
	http.HandleFunc("/test", testHandler)

	log.Println("Start server on port 8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
