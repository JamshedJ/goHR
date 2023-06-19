package service

import (
	"context"
	"log"

	"github.com/JamshedJ/goHR/internal/models"
)

func (a *App) CreateEmployee(ctx context.Context, u models.User, e models.Employee) (id int, err error) {
	if !u.IsAdmin() {
		return 0, models.ErrUnauthorized
	}
	if !e.Validate() {
		return 0, models.ErrBadRequest
	}
	id, err = a.db.CreateEmployee(ctx, e)
	if err != nil {
		log.Println("app CreateEmployee", err)
	}
	return
}

func (a *App) GetEmployeeByID(ctx context.Context, u models.User, id int) (employee models.Employee, err error) {
	if id <= 0 {
		return employee, models.ErrBadRequest
	}
	employee, err = a.db.GetEmployeeByID(ctx, id)
	if err != nil {
		if err != models.ErrNoRows {
			log.Println("app GetEmployeeByID", err)
		}
		return
	}
	if !u.IsAdmin() {
		employee.ID = 0
		employee.Salary = 0
	}
	return
}

func (a *App) GetEmployees(ctx context.Context, u models.User) (employees []models.Employee, err error) {
	employees, err = a.db.GetEmployees(ctx)
	if err != nil {
		if err != models.ErrNoRows {
			log.Println("app GetEmployees", err)
		}
		return
	}
	if !u.IsAdmin() {
		for _, e := range employees {
			e.ID = 0
			e.Salary = 0
		}
	}
	return
}

func (a *App) UpdateEmployee(ctx context.Context, u models.User, e models.Employee) (err error) {
	if !u.IsAdmin() {
		return models.ErrUnauthorized
	}
	if !e.Validate() || e.ID <= 0 {
		return models.ErrBadRequest
	}

	err = a.db.UpdateEmployee(ctx, e)
	if err != nil {
		log.Println("app UpdateEmployee", err)
	}
	return
}

func (a *App) DeleteEmployee(ctx context.Context, u models.User, id int) (err error) {
	if !u.IsAdmin() {
		return models.ErrUnauthorized
	}
	if id <= 0 {
		return models.ErrBadRequest
	}

	err = a.db.DeleteEmployee(ctx, id)
	if err != nil {
		log.Println("app DeleteEmployee", err)
	}
	return
}
