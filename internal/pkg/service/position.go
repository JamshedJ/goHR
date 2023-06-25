package service

import (
	"context"

	"github.com/JamshedJ/goHR/internal/log"
	"github.com/JamshedJ/goHR/internal/models"
)

func (s *Service) CreatePosition(ctx context.Context, p models.Position) (id int, err error) {
	if err = p.Validate(); err != nil {
		log.Warning.Println("service CreatePosition", err)
		return
	}

	id, err = s.db.CreatePosition(ctx, p)
	if err != nil {
		log.Error.Println("service CreatePosition", err)
	}
	return
}

func (s *Service) GetPositionByID(ctx context.Context, id int, isAdmin bool) (position models.Position, err error) {
	if id <= 0 {
		err = models.NewErrorBadRequest("invalid id")
		log.Warning.Println("service GetPositionByID", err)
		return
	}

	position, err = s.db.GetPositionByID(ctx, id)
	if err != nil {
		if err == models.ErrNoRows {
			log.Warning.Println("service GetPositionByID", err)
			return
		}
		log.Error.Println("service GetPositionByID", err)
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
		if err == models.ErrNoRows {
			log.Warning.Println("service GetAllPositions", err)
			return
		}
		log.Error.Println("service GetAllPositions", err)
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
	if err = p.Validate(); err != nil {
		log.Warning.Println("service UpdatePosition", err)
		return
	}

	err = s.db.UpdatePosition(ctx, p)
	if err != nil {
		log.Error.Println("service UpdatePosition", err)
	}
	return
}

func (s *Service) DeletePosition(ctx context.Context, id int) (err error) {
	if id <= 0 {
		err = models.NewErrorBadRequest("invalid id")
		log.Warning.Println("service DeletePosition", err)
		return
	}

	err = s.db.DeletePosition(ctx, id)
	if err != nil {
		log.Error.Println("service DeletePosition", err)
	}
	return
}
