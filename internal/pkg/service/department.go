package service

import (
	"context"

	"github.com/JamshedJ/goHR/internal/log"
	"github.com/JamshedJ/goHR/internal/models"
)

func (s *Service) CreateDepartment(ctx context.Context, d models.Department) (id int, err error) {
	if !d.Validate() {
		return 0, models.ErrBadRequest
	}

	id, err = s.db.CreateDepartment(ctx, d)
	if err != nil {
		log.Error.Println("app CreateDepartment", err)
	}
	return
}

func (s *Service) GetDepartmentByID(ctx context.Context, id int) (department models.Department, err error) {
	if id <= 0 {
		return department, models.ErrBadRequest
	}

	department, err = s.db.GetDepartmentByID(ctx, id)
	if err != nil && err != models.ErrNoRows {
		log.Error.Println("app GetDepartmentByID", err)
	}
	return
}

func (s *Service) GetAllDepartments(ctx context.Context) (departments []models.Department, err error) {
	departments, err = s.db.GetAllDepartments(ctx)
	if err != nil && err != models.ErrNoRows {
		log.Error.Println("app GetAllDepartments", err)
	}
	return
}

func (s *Service) UpdateDepartment(ctx context.Context, d models.Department) (err error) {
	if !d.Validate() {
		return models.ErrBadRequest
	}

	err = s.db.UpdateDepartment(ctx, d)
	if err != nil {
		log.Error.Println("app UpdateDepartment", err)
	}
	return
}

func (s *Service) DeleteDepartment(ctx context.Context, id int) (err error) {
	if id <= 0 {
		return models.ErrBadRequest
	}

	err = s.db.DeleteDepartment(ctx, id)
	if err != nil {
		log.Error.Println("app DeleteEmployee", err)
	}
	return
}
