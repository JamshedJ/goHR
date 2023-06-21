package service

import (
	"context"
	"log"

	"github.com/JamshedJ/goHR/internal/models"
)

func (a *App) CreateRequest(ctx context.Context, r models.Request) (id int, err error) {
	if !r.Validate() {
		return 0, models.ErrBadRequest
	}

	id, err = a.db.CreateRequest(ctx, r)
	if err != nil {
		log.Println("app CreateRequest", err)
	}
	return
}

func (a *App) GetRequestByID(ctx context.Context, id int) (request models.Request, err error) {
	if id <= 0 {
		return request, models.ErrBadRequest
	}
	request, err = a.db.GetRequestByID(ctx, id)
	if err != nil {
		if err != models.ErrNoRows {
			log.Println("app GetRequestByID", err)
		}
		return
	}
	return
}

func (a *App) GetAllRequests(ctx context.Context) (requests []models.Request, err error) {
	requests, err = a.db.GetAllRequests(ctx)
	if err != nil {
		if err != models.ErrNoRows {
			log.Println("app GetRequests", err)
		}
		return
	}
	return
}

func (a *App) UpdateRequest(ctx context.Context, r models.Request) (err error) {
	err = a.db.UpdateRequest(ctx, r)
	if err != nil {
		log.Println("app UpdateRequest", err)
	}
	return
}

func (a *App) DeleteRequest(ctx context.Context, id int) (err error) {
	if id <= 0 {
		return models.ErrBadRequest
	}

	err = a.db.DeleteRequest(ctx, id)
	if err != nil {
		log.Println("app DeleteRequest", err)
	}
	return
}
