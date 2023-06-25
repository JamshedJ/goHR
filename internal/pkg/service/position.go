package service

import (
	"context"

	"github.com/JamshedJ/goHR/internal/log"
	"github.com/JamshedJ/goHR/internal/models"
)

func (s *Service) CreatePosition(ctx context.Context, p models.Position) (id int, err error) {
	if !p.Validate() {
		return 0, models.ErrBadRequest
	}

	id, err = s.db.CreatePosition(ctx, p)
	if err != nil {
		log.Error.Println("app CreatePosition", err)
	}
	return
}

func (s *Service) GetPositionByID(ctx context.Context, id int, isAdmin bool) (position models.Position, err error) {
	if id <= 0 {
		return position, models.ErrBadRequest
	}

	position, err = s.db.GetPositionByID(ctx, id)
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

func (s *Service) GetAllPositions(ctx context.Context, isAdmin bool) (positions []models.Position, err error) {
	positions, err = s.db.GetAllPositions(ctx)
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

func (s *Service) UpdatePosition(ctx context.Context, p models.Position) (err error) {
	if !p.Validate() {
		return models.ErrBadRequest
	}

	err = s.db.UpdatePosition(ctx, p)
	if err != nil {
		log.Error.Println("app UpdatePosition", err)
	}
	return
}

func (s *Service) DeletePosition(ctx context.Context, id int) (err error) {
	if id <= 0 {
		return models.ErrBadRequest
	}

	err = s.db.DeletePosition(ctx, id)
	if err != nil {
		log.Error.Println("app DeletePosition", err)
	}
	return
}
