package main

import (
	"github.com/saintvrai/Drom"
	"github.com/saintvrai/Drom/pkg/handler"
	"github.com/saintvrai/Drom/pkg/repository"
	"github.com/saintvrai/Drom/pkg/service"
	"github.com/spf13/viper"
	"log"
)

func main() {

	if err := initConfig(); err != nil {
		log.Fatalf("error initialializing configs:  %s", err.Error())
	}
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     "localhost",
		Port:     "5436",
		Username: "postgres",
		Password: "qwerty",
		DBName:   "postgres",
		SSLMode:  "disable",
	})
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(Drom.Server)
	if err := srv.Run(viper.GetString("8000"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}

}
func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
