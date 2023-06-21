package service

import (
	"context"
	"log"

	"github.com/JamshedJ/goHR/internal/models"
)

func (a *App) CreateDepartment(ctx context.Context, u models.User, d models.Department) (id int, err error) {
	if !u.IsAdmin() {
		return 0, models.ErrUnauthorized
	}
	if !d.Validate() {
		return 0, models.ErrBadRequest
	}

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

func (a *App) GetAllDepartments(ctx context.Context, u models.User) (departments []models.Department, err error) {
	if !u.IsAdmin() {
		return nil, models.ErrUnauthorized
	}

	departments, err = a.db.GetAllDepartments(ctx)
	if err != nil {
		if err != models.ErrNoRows {
			log.Println("app GetAllDepartments", err)
		}
		return
	}
	return
}

func (a *App) UpdateDepartment(ctx context.Context, u models.User, d models.Department) (err error) {
	if !u.IsAdmin() {
		return models.ErrUnauthorized
	}
	if !d.Validate() {
		return models.ErrBadRequest
	}

	err = a.db.UpdateDepartment(ctx, d)
	if err != nil {
		log.Println("app UpdateDepartment", err)
	}
	return
}

func (a *App) DeleteDepartment(ctx context.Context, u models.User, id int) (err error) {
	if id <= 0 {
		return models.ErrBadRequest
	}
	if !u.IsAdmin() {
		return models.ErrUnauthorized
	}

	err = a.db.DeleteDepartment(ctx, id)
	if err != nil {
		log.Println("app DeleteEmployee", err)
	}
	return
}
