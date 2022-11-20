package main

import (
	_ "github.com/lib/pq"
	"github.com/saintvrai/Drom"
	"github.com/saintvrai/Drom/pkg/handler"
	"github.com/saintvrai/Drom/pkg/repository"
	"github.com/saintvrai/Drom/pkg/service"
	"github.com/spf13/viper"
	"github.com/subosito/gotenv"
	"log"
	"os"
)

func main() {

	if err := initConfig(); err != nil {
		log.Fatalf("error initialializing configs:  %s", err.Error())
	}
	if err := gotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(Drom.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}

}
func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
