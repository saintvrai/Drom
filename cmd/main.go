package main

import (
	"github.com/saintvrai/Drom"
	"log"
)

func main() {
	srv := new(Drom.Server)
	if err := srv.Run("8000"); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
