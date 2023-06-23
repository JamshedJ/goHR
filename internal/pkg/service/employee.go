package service

import (
	"context"

	"github.com/JamshedJ/goHR/internal/log"

	"github.com/JamshedJ/goHR/internal/models"
)

func (a *App) CreateEmployee(ctx context.Context, e models.Employee) (id int, err error) {
	if !e.Validate() {
		return 0, models.ErrBadRequest
	}
	id, err = a.db.CreateEmployee(ctx, e)
	if err != nil {
		log.Error.Println("app CreateEmployee", err)
	}
	return
}

func (a *App) GetEmployeeByID(ctx context.Context, id int, isAdmin bool) (employee models.Employee, err error) {
	if id <= 0 {
		return employee, models.ErrBadRequest
	}
	employee, err = a.db.GetEmployeeByID(ctx, id)
	if err != nil {
		if err != models.ErrNoRows {
			log.Error.Println("app GetEmployeeByID", err)
		}
		return
	}
	if !isAdmin {
		employee.Salary = 0
	}
	return
}

func (a *App) GetEmployees(ctx context.Context, isAdmin bool) (employees []models.Employee, err error) {
	employees, err = a.db.GetEmployees(ctx)
	if err != nil {
		if err != models.ErrNoRows {
			log.Error.Println("app GetEmployees", err)
		}
		return
	}
	if !isAdmin {
		for _, e := range employees {
			e.Salary = 0
		}
	}
	return
}

func (a *App) UpdateEmployee(ctx context.Context, e models.Employee) (err error) {
	if !e.Validate() || e.ID <= 0 {
		return models.ErrBadRequest
	}

	err = a.db.UpdateEmployee(ctx, e)
	if err != nil {
		log.Error.Println("app UpdateEmployee", err)
	}
	return
}

func (a *App) DeleteEmployee(ctx context.Context, id int) (err error) {
	if id <= 0 {
		return models.ErrBadRequest
	}

	err = a.db.DeleteEmployee(ctx, id)
	if err != nil {
		log.Error.Println("app DeleteEmployee", err)
	}
	return
}
