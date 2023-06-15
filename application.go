package goHR

import "github.com/JamshedJ/goHR/internal/database"

type App struct {
	db *database.DB
}

func New(db *database.DB) *App {
	return &App{db: db}
}
