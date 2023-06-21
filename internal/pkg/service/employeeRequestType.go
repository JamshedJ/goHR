package service

import (
	"context"
	"log"

	"github.com/JamshedJ/goHR/internal/models"
)

func (a *App) CreateEmployeeRequestType(ctx context.Context, u models.User, e models.EmployeeRequestType) (id int, err error) {
	if !u.IsAdmin() {
		return 0, models.ErrUnauthorized
	}
	if !e.Validate() {
		return 0, models.ErrBadRequest
	}
	id, err = a.db.CreateEmployeeRequestType(ctx, e)
	if err != nil {
		log.Println("app CreateEmployeeRequestType", err)
	}
	return
}

func (a *App) GetEmployeeRequestTypeByID(ctx context.Context, id int) (erTypes models.EmployeeRequestType, err error) {
	if id <= 0 {
		return erTypes, models.ErrBadRequest
	}
	erTypes, err = a.db.GetEmployeeRequestTypeByID(ctx, id)
	if err != nil {
		if err != models.ErrNoRows {
			log.Println("app GetEmployeeRequestTypeByID", err)
		}
	}
	return
}

func (a *App) GetAllEmployeeRequestTypes(ctx context.Context) (erTypes []models.EmployeeRequestType, err error) {
	erTypes, err = a.db.GetAllEmployeeRequestTypes(ctx)
	if err != nil {
		if err != models.ErrNoRows {
			log.Println("app GetAllEmployeeRequestTypes", err)
		}
	}
	return
}

func (a *App) UpdateEmployeeRequestType(ctx context.Context, u models.User, e models.EmployeeRequestType) (err error) {
	if !u.IsAdmin() {
		return models.ErrUnauthorized
	}
	if !e.Validate() || e.ID < 1 {
		return models.ErrBadRequest
	}
	err = a.db.UpdateEmployeeRequestType(ctx, e)
	if err != nil {
		log.Println("app UpdateEmployeeRequestType", err)
	}
	return
}

func (a *App) DeleteEmployeeRequestType(ctx context.Context, u models.User, id int) (err error) {
	if !u.IsAdmin() {
		return models.ErrUnauthorized
	}
	if id <= 0 {
		return models.ErrBadRequest
	}

	err = a.db.DeleteEmployeeRequestType(ctx, id)
	if err != nil {
		log.Println("app DeleteEmployeeRequestType", err)
	}
	return
}
