package main

import (
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/jackc/pgx"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/saintvrai/Drom"
	"github.com/saintvrai/Drom/pkg/handler"
	"github.com/saintvrai/Drom/pkg/logging"
	"github.com/saintvrai/Drom/pkg/repository"
	"github.com/saintvrai/Drom/pkg/service"
	"github.com/spf13/viper"
	"github.com/subosito/gotenv"
	"golang.org/x/net/context"
	"os"
	"os/signal"
	"syscall"
)

// @title Drom App Api
// @version 1.0
// @description API Server for DromApp Application
// @host localhost:8000
// @BasePath /
func main() {

	log := logging.GetLogger()
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
	migrateDB(db, viper.GetString("db.dbname"))
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(car.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			log.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	log.Print("DromApp Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Print("DromApp Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Fatalf("error occured on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		log.Fatalf("error occured on db connection close: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
func migrateDB(db *sqlx.DB, dbname string) {
	log := logging.GetLogger()
	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		log.Fatalf("couldn't get database instance for running migrations; %s", err.Error())
	}

	m, err := migrate.NewWithDatabaseInstance("file://db/migrations", dbname, driver)
	if err != nil {
		log.Fatalf("couldn't create migrate instance; %s", err.Error())
	}

	if err := m.Up(); err != nil {
		log.Printf("couldn't run database migrations; %s", err.Error())
	} else {
		log.Println("database migration was run successfully")
	}
}
