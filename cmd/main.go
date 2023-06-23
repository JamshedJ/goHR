package main

import (
	"context"

	"github.com/JamshedJ/goHR/internal/configs"
	"github.com/JamshedJ/goHR/internal/log"
	"github.com/JamshedJ/goHR/internal/pkg/handler"
	"github.com/JamshedJ/goHR/internal/pkg/repository"
	"github.com/JamshedJ/goHR/internal/pkg/service"
)

func main() {
	configs.PutAdditionalSettings()

	if err := log.Init(); err != nil {
		panic(err)
	}

	log.Info.Println("App started")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	db := repository.New(ctx, "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
	defer db.Close(ctx)

	if err := db.Up(ctx); err != nil {
		log.Error.Fatal("Error while migrating tables: ", err)
	}

	// if err := db.Down(ctx); err != nil {
	// 	log.Fatal("Error while dropping tables: ", err)
	// 	return
	// }

	app := service.New(db)

	if err := handler.Run(ctx, app, ":8080"); err != nil {
		log.Error.Fatal(err)
	}

	log.Info.Println("App exited")
}
