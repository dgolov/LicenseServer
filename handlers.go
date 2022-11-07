package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// Handlers

func checkHandler(w http.ResponseWriter, r *http.Request) {
	// Ручка проверки лицензии
	var payload Payload

	log.Println(r.Method, r.RequestURI)

	if r.Body == nil {
		http.Error(w, "Please send a request body", http.StatusBadRequest)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	message, statusCode := checkLicense(payload)
	if statusCode != 200 {
		http.Error(w, message, statusCode)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	// health check
	log.Println(r.Method, r.RequestURI)
	w.WriteHeader(http.StatusOK)
}
