package service

import (
	"context"

	"github.com/JamshedJ/goHR/internal/log"
	"github.com/JamshedJ/goHR/internal/models"
)

func (s *Service) CreateEmployeeRequestType(ctx context.Context, e models.EmployeeRequestType) (id int, err error) {
	if err = e.Validate(); err != nil {
		log.Warning.Println("service CreateEmployeeRequestType", err)
		return
	}
	id, err = s.db.CreateEmployeeRequestType(ctx, e)
	if err != nil {
		log.Error.Println("service CreateEmployeeRequestType", err)
	}
	return
}

func (s *Service) GetEmployeeRequestTypeByID(ctx context.Context, id int) (erTypes models.EmployeeRequestType, err error) {
	if id <= 0 {
		err = models.NewErrorBadRequest("invalid id")
		log.Warning.Println("service GetEmploeeRequestTypeByID", err)
		return
	}
	erTypes, err = s.db.GetEmployeeRequestTypeByID(ctx, id)
	if err != nil {
		if err == models.ErrNoRows {
			log.Warning.Println("service GetEmployeeRequestTypeByID", err)
			return
		}
		log.Error.Println("service GetEmployeeRequestTypeByID", err)
	}
	return
}

func (s *Service) GetAllEmployeeRequestTypes(ctx context.Context) (erTypes []models.EmployeeRequestType, err error) {
	erTypes, err = s.db.GetAllEmployeeRequestTypes(ctx)
	if err != nil {
		if err == models.ErrNoRows {
			log.Warning.Println("service GetAllEmployeeRequestTypes", err)
			return
		}
		log.Error.Println("service GetAllEmployeeRequestTypes", err)
	}
	return
}

func (s *Service) UpdateEmployeeRequestType(ctx context.Context, e models.EmployeeRequestType) (err error) {
	if e.ID < 1 {
		err = models.NewErrorBadRequest("invalid id")
	} else {
		err = e.Validate()
	}
	if err != nil {
		log.Warning.Println("service UpdateEmployeeRequestType", err)
		return
	}
	err = s.db.UpdateEmployeeRequestType(ctx, e)
	if err != nil {
		log.Error.Println("service UpdateEmployeeRequestType", err)
	}
	return
}

func (s *Service) DeleteEmployeeRequestType(ctx context.Context, id int) (err error) {
	if id <= 0 {
		err = models.NewErrorBadRequest("invalid id")
		log.Warning.Println("service DeleteEmployeeRequestType", err)
		return
	}

	err = s.db.DeleteEmployeeRequestType(ctx, id)
	if err != nil {
		log.Error.Println("service DeleteEmployeeRequestType", err)
	}
	return
}
