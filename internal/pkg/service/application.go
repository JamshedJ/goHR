package service

import (
	"github.com/JamshedJ/goHR/internal/log"
	"github.com/JamshedJ/goHR/internal/pkg/repository"
)

type App struct {
	db *repository.DB
}

func New(db *repository.DB) *App {
	log.Debug.Println("service instance created")
	return &App{db: db}
}
