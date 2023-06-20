package service

import (
	"context"
	"log"

	"github.com/JamshedJ/goHR/internal/models"
)

func (a *App) CreateDepartment(ctx context.Context, d models.Department) (id int, err error) {
	id, err = a.db.CreateDepartment(ctx, d)
	if err != nil {
		log.Println("app CreateDepartment", err)
	}
	return
}

func (a *App) GetDepartmentByID(ctx context.Context, id int) (department models.Department, err error) {
	if id <= 0 {
		return department, models.ErrBadRequest
	}
	department, err = a.db.GetDepartmentByID(ctx, id)
	if err != nil {
		if err != models.ErrNoRows {
			log.Println("app GetDepartmentByID", err)
		}
		return
	}
	return
}

func (a *App) GetAllDepartments(ctx context.Context) (departments []models.Department, err error) {
	departments, err = a.db.GetAllDepartments(ctx)
	if err != nil {
		if err != models.ErrNoRows {
			log.Println("app GetAllDepartments", err)
		}
		return
	}
	return
}

func (a *App) UpdateDepartment(ctx context.Context, d models.Department) (err error) {
	err = a.db.UpdateDepartment(ctx, d)
	if err != nil {
		log.Println("app UpdateEmployee", err)
	}
	return
}

func (a *App) DeleteDepartment(ctx context.Context, id int) (err error) {
	if id <= 0 {
		return models.ErrBadRequest
	}

	err = a.db.DeleteDepartment(ctx, id)
	if err != nil {
		log.Println("app DeleteEmployee", err)
	}
	return
}