package service

import (
	"context"
	"github.com/JamshedJ/goHR/internal/models"
	"log"
)

func (a *App) GetUserById(ctx context.Context, id, userID int) (user models.User, err error) {
	if id <= 0 {
		return user, models.ErrBadRequest
	}
	if id != userID {
		return user, models.ErrUnauthorized
	}

	user, err = a.db.GetUserById(ctx, id)
	if err != nil && err != models.ErrNoRows {
		log.Println("app GetUser", err)
	}
	return
}

func (a *App) GetAllUsers(ctx context.Context) (users []models.User, err error) {
	users, err = a.db.GetAllUsers(ctx)
	if err != nil && err != models.ErrNoRows {
		log.Println("app GetAllUsers", err)
	}
	return
}

func (a *App) AddUser(ctx context.Context, u models.User) (err error) {
	if !u.Validate() {
		err = models.ErrBadRequest
		return
	}
	u.Password = generatePasswordHash(u.Password)
	err = a.db.CreateUser(ctx, u)
	if err != nil {
		log.Println("app AddUser", err)
	}
	return
}

func (a *App) DeleteUser(ctx context.Context, id, userID int) (err error) {
	if id <= 0 {
		return models.ErrBadRequest
	}
	if id != userID {
		return models.ErrUnauthorized
	}
	err = a.db.DeleteUser(ctx, id)
	if err != nil {
		log.Println("app DeleteUser", err)
	}
	return
}

func (a *App) UpdateUser(ctx context.Context, id, userID int, u models.User) (err error) {
	if id <= 0 || !u.Validate() {
		return models.ErrBadRequest
	}
	if id != userID {
		return models.ErrUnauthorized
	}
	err = a.db.UpdateUser(ctx, id, u)
	if err != nil {
		log.Println("app UpdateUser", err)
	}
	return
}
