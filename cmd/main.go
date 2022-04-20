package main

import (
	"log"

	"github.com/VMironiuk/offlinescore"
	"github.com/VMironiuk/offlinescore/pkg/handler"
	"github.com/VMironiuk/offlinescore/pkg/repository"
	"github.com/VMironiuk/offlinescore/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(offlinescore.Server)

	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("an error is occured when running HTTP server: %s\n", err.Error())
	}
}
