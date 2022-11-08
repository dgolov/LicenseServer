package main

import (
	"github.com/dgolov/LicenseServer"
	"github.com/dgolov/LicenseServer/pkg/handler"
	"log"
)

// Main

func main() {
	handlers := new(handler.Handler)
	srv := new(LicenseServer.Server)

	if err := srv.Run("8001", handlers.InitRoutes()); err != nil {
		log.Fatalf("Error while runing http server: %s", err.Error())
		return
	}
}
