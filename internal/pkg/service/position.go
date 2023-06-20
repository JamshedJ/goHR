package service

import (
	"context"
	"log"

	"github.com/JamshedJ/goHR/internal/models"
)

func (a *App) CreatePosition(ctx context.Context, p models.Position) (id int, err error) {
	id, err = a.db.CreatePosition(ctx, p)
	if err != nil {
		log.Println("app CreatePosition", err)
	}
	return
}

func (a *App) GetPositionByID(ctx context.Context, id int) (position models.Position, err error) {
	if id <= 0 {
		return position, models.ErrBadRequest
	}
	position, err = a.db.GetPositionByID(ctx, id)
	if err != nil {
		if err != models.ErrNoRows {
			log.Println("app GetPositionByID", err)
		}
		return
	}
	return
}

func (a *App) GetAllPositions(ctx context.Context) (positions []models.Position, err error) {
	positions, err = a.db.GetAllPositions(ctx)
	if err != nil {
		if err != models.ErrNoRows {
			log.Println("app GetAllPositions", err)
		}
		return
	}
	return
}

func (a *App) UpdatePosition(ctx context.Context, p models.Position) (err error) {
	err = a.db.UpdatePosition(ctx, p)
	if err != nil {
		log.Println("app UpdatePosition", err)
	}
	return
}

func (a *App) DeletePosition(ctx context.Context, id int) (err error) {
	if id <= 0 {
		return models.ErrBadRequest
	}

	err = a.db.DeletePosition(ctx, id)
	if err != nil {
		log.Println("app DeletePosition", err)
	}
	return
}