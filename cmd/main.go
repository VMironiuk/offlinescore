package main

import (
	"log"

	"github.com/VMironiuk/offlinescore"
	"github.com/VMironiuk/offlinescore/pkg/handler"
)

func main() {
	handlers := handler.NewHandler()
	srv := new(offlinescore.Server)

	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("an error is occured when running HTTP server: %s\n", err.Error())
	}
}
