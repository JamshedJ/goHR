package service

import (
	"context"
	"log"

	"github.com/JamshedJ/goHR/internal/models"
)

func (a *App) CreatePosition(ctx context.Context, u models.User, p models.Position) (id int, err error) {
	if !u.IsAdmin() {
		return 0, models.ErrUnauthorized
	}
	if !p.Validate() {
		return 0, models.ErrBadRequest
	}

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

func (a *App) GetAllPositions(ctx context.Context, u models.User) (positions []models.Position, err error) {
	if !u.IsAdmin() {
		return nil, models.ErrUnauthorized
	}

	positions, err = a.db.GetAllPositions(ctx)
	if err != nil {
		if err != models.ErrNoRows {
			log.Println("app GetAllPositions", err)
		}
		return
	}
	return
}

func (a *App) UpdatePosition(ctx context.Context, u models.User, p models.Position) (err error) {
	if !u.IsAdmin() {
		return models.ErrUnauthorized
	}
	if !p.Validate() {
		return models.ErrBadRequest
	}

	err = a.db.UpdatePosition(ctx, p)
	if err != nil {
		log.Println("app UpdatePosition", err)
	}
	return
}

func (a *App) DeletePosition(ctx context.Context, u models.User, id int) (err error) {
	if id <= 0 {
		return models.ErrBadRequest
	}
	if !u.IsAdmin() {
		return models.ErrUnauthorized
	}

	err = a.db.DeletePosition(ctx, id)
	if err != nil {
		log.Println("app DeletePosition", err)
	}
	return
}
