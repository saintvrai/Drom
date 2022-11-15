package main

import (
	"github.com/saintvrai/Drom"
	"github.com/saintvrai/Drom/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(Drom.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
