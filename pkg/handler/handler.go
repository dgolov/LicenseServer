package handler

import (
	"encoding/json"
	"github.com/dgolov/LicenseServer/pkg/service"
	"log"
	"net/http"
)

type Payload struct {
	// Структура запроса
	LicenseUuid        string `json:"license_uuid"`
	HardwareParameters string `json:"hardware_parameters"`
}

type Handler struct {
}

// Handlers

func (h *Handler) CheckHandler(w http.ResponseWriter, r *http.Request) {
	// Ручка проверки лицензии
	var payload Payload

	check := new(service.Check)

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

	log.Println("Request: ", payload)
	message, statusCode := check.CheckLicense(payload.LicenseUuid, payload.HardwareParameters)
	if statusCode != 200 {
		http.Error(w, message, statusCode)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) TestHandler(w http.ResponseWriter, r *http.Request) {
	// health check
	log.Println(r.Method, r.RequestURI)
	w.WriteHeader(http.StatusOK)
}
