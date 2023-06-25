package service

import (
	"context"

	"github.com/JamshedJ/goHR/internal/log"
	"github.com/JamshedJ/goHR/internal/models"
)

func (s *Service) CreateEmployeeRequest(ctx context.Context, r models.EmployeeRequest) (id int, err error) {
	if !r.Validate() {
		return 0, models.ErrBadRequest
	}

	id, err = s.db.CreateEmployeeRequest(ctx, r)
	if err != nil {
		log.Error.Println("app CreateEmployeeRequest", err)
	}
	return
}

func (s *Service) GetEmployeeRequestByID(ctx context.Context, id int) (request models.EmployeeRequest, err error) {
	if id <= 0 {
		return request, models.ErrBadRequest
	}
	request, err = s.db.GetRequestByID(ctx, id)
	if err != nil {
		if err != models.ErrNoRows {
			log.Error.Println("app GetEmployeeRequestByID", err)
		}
	}
	return
}

func (s *Service) GetAllEmployeeRequests(ctx context.Context) (requests []models.EmployeeRequest, err error) {
	requests, err = s.db.GetAllRequests(ctx)
	if err != nil {
		if err != models.ErrNoRows {
			log.Error.Println("app GetRequests", err)
		}
	}
	return
}

func (s *Service) UpdateEmployeeRequest(ctx context.Context, r models.EmployeeRequest) (err error) {
	if !r.Validate() {
		return models.ErrBadRequest
	}
	err = s.db.UpdateEmployeeRequest(ctx, r)
	if err != nil {
		log.Error.Println("app UpdateEmployeeRequest", err)
	}
	return
}

func (s *Service) DeleteEmployeeRequest(ctx context.Context, id int) (err error) {
	if id <= 0 {
		return models.ErrBadRequest
	}

	err = s.db.DeleteRequest(ctx, id)
	if err != nil {
		log.Error.Println("app DeleteEmployeeRequest", err)
	}
	return
}
