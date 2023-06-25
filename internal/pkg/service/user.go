package service

import (
	"context"

	"github.com/JamshedJ/goHR/internal/log"
	"github.com/JamshedJ/goHR/internal/models"
)

func (s *Service) GetUserById(ctx context.Context, id int) (user models.User, err error) {
	if id <= 0 {
		err = models.NewErrorBadRequest("invalid id")
		log.Warning.Println("service GetUserById", err)
		return
	}

	user, err = s.db.GetUserById(ctx, id)
	if err != nil {
		if err == models.ErrNoRows {
			log.Warning.Println("service GetUserById", err)
			return
		}
		log.Error.Println("service GetUserById", err)
	}
	return
}

func (s *Service) GetAllUsers(ctx context.Context) (users []models.User, err error) {
	users, err = s.db.GetAllUsers(ctx)
	if err != nil {
		if err == models.ErrNoRows {
			log.Warning.Println("service GetAllUsers", err)
			return
		}
		log.Error.Println("service GetAllUsers", err)
	}
	return
}

func (s *Service) AddUser(ctx context.Context, u models.User) (err error) {
	if err = u.Validate(); err != nil {
		log.Warning.Println("service AddUser", err)
		return
	}
	u.Password = generatePasswordHash(u.Password)
	err = s.db.CreateUser(ctx, u)
	if err != nil {
		log.Error.Println("service AddUser", err)
	}
	return
}

func (s *Service) DeleteUser(ctx context.Context, id int) (err error) {
	if id <= 0 {
		err = models.NewErrorBadRequest("invalid id")
		log.Warning.Println("service DeleteUser", err)
		return
	}

	err = s.db.DeleteUser(ctx, id)
	if err != nil {
		log.Error.Println("service DeleteUser", err)
	}
	return
}

func (s *Service) UpdateUser(ctx context.Context, user models.User) (err error) {
	if user.ID <= 0 {
		err = models.NewErrorBadRequest("invalid id")
	} else {
		err = user.Validate()
	}
	if err != nil {
		log.Warning.Println("service UpdateUser", err)
		return
	}

	err = s.db.UpdateUser(ctx, user)
	if err != nil {
		log.Error.Println("service UpdateUser", err)
	}
	return
}
