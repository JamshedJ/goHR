package service

import (
	"github.com/JamshedJ/goHR/internal/log"
	"github.com/JamshedJ/goHR/internal/pkg/repository"
)

type Service struct {
	db *repository.DB
}

func New(db *repository.DB) *Service {
	log.Debug.Println("service instance created")
	return &Service{db: db}
}
