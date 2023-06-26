package main

import (
	"context"

	"github.com/JamshedJ/goHR/internal/configs"
	"github.com/JamshedJ/goHR/internal/log"
	"github.com/JamshedJ/goHR/internal/pkg/repository"
	"github.com/JamshedJ/goHR/internal/pkg/server"
	"github.com/JamshedJ/goHR/internal/pkg/service"
)

func main() {
	if err := configs.Init(); err != nil {
		panic(err)
	}
	if err := log.Init(); err != nil {
		panic(err)
	}

	log.Info.Println("App started")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	db := repository.New(ctx)
	defer db.Close(ctx)

	if err := db.Up(ctx); err != nil {
		log.Error.Fatal("Error while migrating tables: ", err)
	}

	//if err := db.Down(ctx); err != nil {
	//	log.Error.Fatal("Error while dropping tables: ", err)
	//	return
	//}

	app := service.New(db)

	if err := server.Run(ctx, app, configs.App.URL); err != nil {
		log.Error.Fatal(err)
	}

	log.Info.Println("App exited")
}
