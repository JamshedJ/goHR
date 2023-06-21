package service

import (
	"context"
	"log"

	"github.com/JamshedJ/goHR/internal/models"
)

func (a *App) CreateType(ctx context.Context, t models.Type) (id int, err error) {
	id, err = a.db.CreateType(ctx, t)
	if err != nil {
		log.Println("app CreateType", err)
	}
	return
}

func (a *App) GetTypeByID(ctx context.Context, id int) (types models.Type, err error) {
	if id <= 0 {
		return types, models.ErrBadRequest
	}
	types, err = a.db.GetTypeByID(ctx, id)
	if err != nil {
		if err != models.ErrNoRows {
			log.Println("app GetTypeByID", err)
		}
		return
	}
	return
}

func (a *App) GetAllTypes(ctx context.Context) (types []models.Type, err error) {
	types, err = a.db.GetAllTypes(ctx)
	if err != nil {
		if err != models.ErrNoRows {
			log.Println("app GetAllTypes", err)
		}
		return
	}
	return
}

func (a *App) UpdateType(ctx context.Context, t models.Type) (err error) {
	err = a.db.UpdateType(ctx, t)
	if err != nil {
		log.Println("app UpdateType", err)
	}
	return
}

func (a *App) DeleteType(ctx context.Context, id int) (err error) {
	if id <= 0 {
		return models.ErrBadRequest
	}

	err = a.db.DeleteType(ctx, id)
	if err != nil {
		log.Println("app DeleteType", err)
	}
	return
}
