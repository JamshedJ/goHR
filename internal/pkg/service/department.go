package service

import (
	"context"

	"github.com/JamshedJ/goHR/internal/log"
	"github.com/JamshedJ/goHR/internal/models"
)

func (s *Service) CreateDepartment(ctx context.Context, d models.Department) (id int, err error) {
	if err = d.Validate(); err != nil {
		log.Warning.Println("service CreateDepartment", err)
		return
	}

	id, err = s.db.CreateDepartment(ctx, d)
	if err != nil {
		log.Error.Println("service CreateDepartment", err)
	}
	return
}

func (s *Service) GetDepartmentByID(ctx context.Context, id int) (department models.Department, err error) {
	if id <= 0 {
		err = models.NewErrorBadRequest("invalid id")
		log.Warning.Println("service GetDepartmentByID", err)
		return
	}

	department, err = s.db.GetDepartmentByID(ctx, id)
	if err != nil {
		if err == models.ErrNoRows {
			log.Warning.Println("service GetDepartmentByID", err)
			return
		}
		log.Error.Println("service GetDepartmentByID", err)
	}
	return
}

func (s *Service) GetAllDepartments(ctx context.Context) (departments []models.Department, err error) {
	departments, err = s.db.GetAllDepartments(ctx)
	if err != nil {
		if err == models.ErrNoRows {
			log.Warning.Println("service GetAllDepartments", err)
			return
		}
		log.Error.Println("service GetAllDepartments", err)
	}
	return
}

func (s *Service) UpdateDepartment(ctx context.Context, d models.Department) (err error) {
	if err = d.Validate(); err != nil {
		log.Warning.Println("service UpdateDepartment", err)
		return
	}

	err = s.db.UpdateDepartment(ctx, d)
	if err != nil {
		log.Error.Println("service UpdateDepartment", err)
	}
	return
}

func (s *Service) DeleteDepartment(ctx context.Context, id int) (err error) {
	if id <= 0 {
		err = models.NewErrorBadRequest("invalid id")
		log.Warning.Println("service DeleteDepartment", err)
		return
	}

	err = s.db.DeleteDepartment(ctx, id)
	if err != nil {
		log.Error.Println("service DeleteEmployee", err)
	}
	return
}
