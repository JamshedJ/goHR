package main

import (
	"context"
	"github.com/JamshedJ/goHR"
	"github.com/JamshedJ/goHR/pkg/handler"
	"github.com/JamshedJ/goHR/pkg/repository"
	"github.com/JamshedJ/goHR/pkg/service"
	"log"
)

func main() {
	log.Println("App started")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	repos := repository.NewRepository(cfg)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(goHR.Server)
	if err := srv.Run(ctx, "8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error ocured while running http server: %s", err.Error())
	}
}
