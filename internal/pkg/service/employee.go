package service

import (
	"context"
	"github.com/JamshedJ/goHR/internal/models"
	"log"
)

func (a *App) CreateEmployee(ctx context.Context, e models.Employee) (id int, err error) {
	if !e.Validate() {
		return 0, models.ErrBadRequest
	}
	id, err = a.db.CreateEmployee(ctx, e)
	if err != nil {
		log.Println("app CreateEmployee", err)
	}
	return
}

func (a *App) GetEmployeeByID(ctx context.Context, id, positionID, departmentID int) (employee models.Employee, err error) {
	if id <= 0 {
		return employee, models.ErrBadRequest
	}
	employee, err = a.db.GetEmployeeByID(ctx, id, positionID, departmentID)
	if err != nil && err != models.ErrNoRows {
		log.Println("app GetEmployeeByID", err)
	}
	return
}
