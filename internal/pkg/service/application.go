package service

import (
	"github.com/JamshedJ/goHR/internal/pkg/repository"
)

type App struct {
	db *repository.DB
}

func New(db *repository.DB) *App {
	return &App{db: db}
}
