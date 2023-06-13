package service

import "github.com/JamshedJ/goHR/pkg/repository"

type Authorization interface {
}

type Employee interface {
}

type Department interface {
}

type Vacancy interface {
}

type User interface {
}

type Service struct {
	Authorization
	Employee
	Department
	Vacancy
	User
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
