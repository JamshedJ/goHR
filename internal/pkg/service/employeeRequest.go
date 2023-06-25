package service

import (
	"context"

	"github.com/JamshedJ/goHR/internal/log"
	"github.com/JamshedJ/goHR/internal/models"
)

func (s *Service) CreateEmployeeRequest(ctx context.Context, r models.EmployeeRequest) (id int, err error) {
	if err = r.Validate(); err != nil {
		log.Warning.Println("service CreateEmployeeRequest", err)
		return
	}

	id, err = s.db.CreateEmployeeRequest(ctx, r)
	if err != nil {
		log.Error.Println("service CreateEmployeeRequest", err)
	}
	return
}

func (s *Service) GetEmployeeRequestByID(ctx context.Context, id int) (request models.EmployeeRequest, err error) {
	if id <= 0 {
		err = models.NewErrorBadRequest("invalid id")
		log.Warning.Println("service GetEmployeeRequestByID", err)
		return
	}
	request, err = s.db.GetRequestByID(ctx, id)
	if err != nil {
		if err == models.ErrNoRows {
			log.Warning.Println("service GetEmployeeRequestByID", err)
			return
		}
		log.Error.Println("service GetEmployeeRequestByID", err)
	}
	return
}

func (s *Service) GetAllEmployeeRequests(ctx context.Context) (requests []models.EmployeeRequest, err error) {
	requests, err = s.db.GetAllRequests(ctx)
	if err != nil {
		if err == models.ErrNoRows {
			log.Warning.Println("service GetAllEmployeeRequests", err)
			return
		}
		log.Error.Println("service GetAllEmployeeRequests", err)
	}
	return
}

func (s *Service) UpdateEmployeeRequest(ctx context.Context, r models.EmployeeRequest) (err error) {
	if err = r.Validate(); err != nil {
		log.Warning.Println("service UpdateEmployeeRequest", err)
		return
	}
	err = s.db.UpdateEmployeeRequest(ctx, r)
	if err != nil {
		log.Error.Println("service UpdateEmployeeRequest", err)
	}
	return
}

func (s *Service) DeleteEmployeeRequest(ctx context.Context, id int) (err error) {
	if id <= 0 {
		err = models.NewErrorBadRequest("invalid id")
		log.Warning.Println("service DeleteEmployeeRequest", err)
		return
	}

	err = s.db.DeleteRequest(ctx, id)
	if err != nil {
		log.Error.Println("service DeleteEmployeeRequest", err)
	}
	return
}
