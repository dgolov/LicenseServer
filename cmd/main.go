package main

import (
	"github.com/dgolov/LicenseServer"
	"github.com/dgolov/LicenseServer/pkg/handler"
	"log"
)

// Main

func main() {
	log.Println("Start app")
	handlers := new(handler.Handler)
	port := "8001"
	srv := new(LicenseServer.Server)

	if err := srv.Run(port, handlers.InitRoutes()); err != nil {
		log.Fatalf("Error while runing http server: %s", err.Error())
		return
	}

	log.Printf("Server start on port %s", port)
}
