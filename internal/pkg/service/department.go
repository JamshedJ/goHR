package service

import (
	"context"

	"github.com/JamshedJ/goHR/internal/models"
	"github.com/JamshedJ/goHR/internal/log"
)

func (a *App) CreateDepartment(ctx context.Context, d models.Department) (id int, err error) {
	if !d.Validate() {
		return 0, models.ErrBadRequest
	}

	id, err = a.db.CreateDepartment(ctx, d)
	if err != nil {
		log.Error.Println("app CreateDepartment", err)
	}
	return
}

func (a *App) GetDepartmentByID(ctx context.Context, id int) (department models.Department, err error) {
	if id <= 0 {
		return department, models.ErrBadRequest
	}

	department, err = a.db.GetDepartmentByID(ctx, id)
	if err != nil && err != models.ErrNoRows {
		log.Error.Println("app GetDepartmentByID", err)
	}
	return
}

func (a *App) GetAllDepartments(ctx context.Context) (departments []models.Department, err error) {
	departments, err = a.db.GetAllDepartments(ctx)
	if err != nil && err != models.ErrNoRows {
		log.Error.Println("app GetAllDepartments", err)
	}
	return
}

func (a *App) UpdateDepartment(ctx context.Context, d models.Department) (err error) {
	if !d.Validate() {
		return models.ErrBadRequest
	}

	err = a.db.UpdateDepartment(ctx, d)
	if err != nil {
		log.Error.Println("app UpdateDepartment", err)
	}
	return
}

func (a *App) DeleteDepartment(ctx context.Context, id int) (err error) {
	if id <= 0 {
		return models.ErrBadRequest
	}

	err = a.db.DeleteDepartment(ctx, id)
	if err != nil {
		log.Error.Println("app DeleteEmployee", err)
	}
	return
}
