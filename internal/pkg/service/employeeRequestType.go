package service

import (
	"context"

	"github.com/JamshedJ/goHR/internal/log"
	"github.com/JamshedJ/goHR/internal/models"
)

func (s *Service) CreateEmployeeRequestType(ctx context.Context, e models.EmployeeRequestType) (id int, err error) {
	if !e.Validate() {
		return 0, models.ErrBadRequest
	}
	id, err = s.db.CreateEmployeeRequestType(ctx, e)
	if err != nil {
		log.Error.Println("app CreateEmployeeRequestType", err)
	}
	return
}

func (s *Service) GetEmployeeRequestTypeByID(ctx context.Context, id int) (erTypes models.EmployeeRequestType, err error) {
	if id <= 0 {
		return erTypes, models.ErrBadRequest
	}
	erTypes, err = s.db.GetEmployeeRequestTypeByID(ctx, id)
	if err != nil {
		if err != models.ErrNoRows {
			log.Error.Println("app GetEmployeeRequestTypeByID", err)
		}
	}
	return
}

func (s *Service) GetAllEmployeeRequestTypes(ctx context.Context) (erTypes []models.EmployeeRequestType, err error) {
	erTypes, err = s.db.GetAllEmployeeRequestTypes(ctx)
	if err != nil {
		if err != models.ErrNoRows {
			log.Error.Println("app GetAllEmployeeRequestTypes", err)
		}
	}
	return
}

func (s *Service) UpdateEmployeeRequestType(ctx context.Context, e models.EmployeeRequestType) (err error) {
	if !e.Validate() || e.ID < 1 {
		return models.ErrBadRequest
	}
	err = s.db.UpdateEmployeeRequestType(ctx, e)
	if err != nil {
		log.Error.Println("app UpdateEmployeeRequestType", err)
	}
	return
}

func (s *Service) DeleteEmployeeRequestType(ctx context.Context, id int) (err error) {
	if id <= 0 {
		return models.ErrBadRequest
	}

	err = s.db.DeleteEmployeeRequestType(ctx, id)
	if err != nil {
		log.Error.Println("app DeleteEmployeeRequestType", err)
	}
	return
}
