package main

import (
	"github.com/dgolov/LicenseServer"
	"github.com/dgolov/LicenseServer/pkg/handler"
	"github.com/dgolov/LicenseServer/pkg/repository"
	"github.com/dgolov/LicenseServer/pkg/service"
	"github.com/spf13/viper"
	"log"
)

// Main

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Error initializing config file: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		log.Fatalf("Error initializing database: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

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
