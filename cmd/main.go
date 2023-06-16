package main

import (
	"context"
	"github.com/JamshedJ/goHR/internal/pkg/handler"
	"github.com/JamshedJ/goHR/internal/pkg/repository"
	"github.com/JamshedJ/goHR/internal/pkg/service"
	"log"
)

func main() {
	log.Println("App started")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	db := repository.New(ctx, "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
	defer db.Close(ctx)

	if err := automigration.Up(ctx); err != nil {
		log.Fatal("Error while migrating tables: ", err)
		return
	}

	// if err := db.Down(ctx); err != nil {
	// 	log.Fatal("Error while dropping tables: ", err)
	// 	return
	// }

	app := service.New(db)

	if err := handler.Run(ctx, app, ":8080"); err != nil {
		log.Fatal(err)
	}

	log.Println("App exited")
}
