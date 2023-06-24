package service

import (
	"context"

	"github.com/JamshedJ/goHR/internal/models"
	"github.com/JamshedJ/goHR/internal/log"
)

func (a *App) CreatePosition(ctx context.Context, p models.Position) (id int, err error) {
	if !p.Validate() {
		return 0, models.ErrBadRequest
	}

	id, err = a.db.CreatePosition(ctx, p)
	if err != nil {
		log.Error.Println("app CreatePosition", err)
	}
	return
}

func (a *App) GetPositionByID(ctx context.Context, id int, isAdmin bool) (position models.Position, err error) {
	if id <= 0 {
		return position, models.ErrBadRequest
	}

	position, err = a.db.GetPositionByID(ctx, id)
	if err != nil {
		if err != models.ErrNoRows {
			log.Error.Println("app GetPositionByID", err)
		}
		return
	}
	if !isAdmin {
		position.Salary = 0
	}
	return
}

func (a *App) GetAllPositions(ctx context.Context, isAdmin bool) (positions []models.Position, err error) {
	positions, err = a.db.GetAllPositions(ctx)
	if err != nil {
		if err != models.ErrNoRows {
			log.Error.Println("app GetAllPositions", err)
		}
		return
	}
	if !isAdmin {
		for _, p := range positions {
			p.Salary = 0
		}
	}
	return
}

func (a *App) UpdatePosition(ctx context.Context, p models.Position) (err error) {
	if !p.Validate() {
		return models.ErrBadRequest
	}

	err = a.db.UpdatePosition(ctx, p)
	if err != nil {
		log.Error.Println("app UpdatePosition", err)
	}
	return
}

func (a *App) DeletePosition(ctx context.Context, id int) (err error) {
	if id <= 0 {
		return models.ErrBadRequest
	}

	err = a.db.DeletePosition(ctx, id)
	if err != nil {
		log.Error.Println("app DeletePosition", err)
	}
	return
}
