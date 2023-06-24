package service

import (
	"context"

	"github.com/JamshedJ/goHR/internal/models"
	"github.com/JamshedJ/goHR/internal/log"
)

func (a *App) CreateEmployeeRequest(ctx context.Context, r models.EmployeeRequest) (id int, err error) {
	if !r.Validate() {
		return 0, models.ErrBadRequest
	}

	id, err = a.db.CreateEmployeeRequest(ctx, r)
	if err != nil {
		log.Error.Println("app CreateEmployeeRequest", err)
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
			log.Error.Println("app GetEmployeeRequestByID", err)
		}
	}
	return
}

func (a *App) GetAllEmployeeRequests(ctx context.Context) (requests []models.EmployeeRequest, err error) {
	requests, err = a.db.GetAllRequests(ctx)
	if err != nil {
		if err != models.ErrNoRows {
			log.Error.Println("app GetRequests", err)
		}
	}
	return
}

func (a *App) UpdateEmployeeRequest(ctx context.Context, r models.EmployeeRequest) (err error) {
	if !r.Validate() {
		return models.ErrBadRequest
	}
	err = a.db.UpdateEmployeeRequest(ctx, r)
	if err != nil {
		log.Error.Println("app UpdateEmployeeRequest", err)
	}
	return
}

func (a *App) DeleteEmployeeRequest(ctx context.Context, id int) (err error) {
	if id <= 0 {
		return models.ErrBadRequest
	}

	err = a.db.DeleteRequest(ctx, id)
	if err != nil {
		log.Error.Println("app DeleteEmployeeRequest", err)
	}
	return
}
