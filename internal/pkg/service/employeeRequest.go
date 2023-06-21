package service

import (
	"context"
	"log"

	"github.com/JamshedJ/goHR/internal/models"
)

func (a *App) CreateEmployeeRequest(ctx context.Context, u models.User, r models.EmployeeRequest) (id int, err error) {
	if !u.IsAdmin() {
		return 0, models.ErrUnauthorized
	}
	if !r.Validate() {
		return 0, models.ErrBadRequest
	}

	id, err = a.db.CreateEmployeeRequest(ctx, r)
	if err != nil {
		log.Println("app CreateEmployeeRequest", err)
	}
	return
}

func (a *App) GetEmployeeRequestByID(ctx context.Context, id int) (request models.EmployeeRequest, err error) {
	if id <= 0 {
		return request, models.ErrBadRequest
	}
	request, err = a.db.GetRequestByID(ctx, id)
	if err != nil {
		if err != models.ErrNoRows {
			log.Println("app GetEmployeeRequestByID", err)
		}
	}
	return
}

func (a *App) GetAllEmployeeRequests(ctx context.Context, u models.User) (requests []models.EmployeeRequest, err error) {
	if !u.IsAdmin() {
		return nil, models.ErrUnauthorized
	}

	requests, err = a.db.GetAllRequests(ctx)
	if err != nil {
		if err != models.ErrNoRows {
			log.Println("app GetRequests", err)
		}
	}
	return
}

func (a *App) UpdateEmployeeRequest(ctx context.Context, u models.User, r models.EmployeeRequest) (err error) {
	if !u.IsAdmin() {
		return models.ErrUnauthorized
	}
	if !r.Validate() {
		return models.ErrBadRequest
	}
	err = a.db.UpdateEmployeeRequest(ctx, r)
	if err != nil {
		log.Println("app UpdateEmployeeRequest", err)
	}
	return
}

func (a *App) DeleteEmployeeRequest(ctx context.Context, u models.User, id int) (err error) {
	if !u.IsAdmin() {
		return models.ErrUnauthorized
	}
	if id <= 0 {
		return models.ErrBadRequest
	}

	err = a.db.DeleteRequest(ctx, id)
	if err != nil {
		log.Println("app DeleteEmployeeRequest", err)
	}
	return
}
