package service

import (
	"context"

	"github.com/JamshedJ/goHR/internal/log"
	"github.com/JamshedJ/goHR/internal/models"
)

func (s *Service) GetUserById(ctx context.Context, id int) (user models.User, err error) {
	if id <= 0 {
		return user, models.ErrBadRequest
	}

	user, err = s.db.GetUserById(ctx, id)
	if err != nil {
		if err != models.ErrNoRows {
			log.Error.Println("app GetUser", err)
		}
	}
	return
}

func (s *Service) GetAllUsers(ctx context.Context) (users []models.User, err error) {
	users, err = s.db.GetAllUsers(ctx)
	if err != nil && err != models.ErrNoRows {
		log.Error.Println("app GetAllUsers", err)
	}
	return
}

func (s *Service) AddUser(ctx context.Context, u models.User) (err error) {
	if !u.Validate() {
		err = models.ErrBadRequest
		return
	}
	u.Password = generatePasswordHash(u.Password)
	err = s.db.CreateUser(ctx, u)
	if err != nil {
		log.Error.Println("app AddUser", err)
	}
	return
}

func (s *Service) DeleteUser(ctx context.Context, id int) (err error) {
	if id <= 0 {
		return models.ErrBadRequest
	}

	err = s.db.DeleteUser(ctx, id)
	if err != nil {
		log.Error.Println("app DeleteUser", err)
	}
	return
}

func (s *Service) UpdateUser(ctx context.Context, user models.User) (err error) {
	if !user.Validate() || user.ID <= 0 {
		return models.ErrBadRequest
	}

	err = s.db.UpdateUser(ctx, user)
	if err != nil {
		log.Error.Println("app UpdateUser", err)
	}
	return
}
