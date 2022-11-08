package main

import (
	"github.com/dgolov/LicenseServer"
	"github.com/dgolov/LicenseServer/pkg/handler"
	"github.com/spf13/viper"
	"log"
)

// Main

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Error initializing config file: %s", err.Error())
	}

	handlers := new(handler.Handler)
	srv := new(LicenseServer.Server)

	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("Error while runing http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
