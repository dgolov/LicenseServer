package main

import (
	"log"
)

type Payload struct {
	// Структура запроса
	LicenseUuid        string `json:"license_uuid"`
	HardwareParameters string `json:"hardware_parameters"`
}

func checkLicense(payload Payload) (string, int) {
	// Проверка лицензии
	log.Println("Request: ", payload)

	if len(payload.LicenseUuid) == 0 || len(payload.HardwareParameters) == 0 {
		return "Error input data. License uuid or hardware params is not found.", 409
	}

	return "Ok", 200
}
