package service

import (
	"context"
	"log"

	"github.com/JamshedJ/goHR/internal/models"
)

func (a *App) GetUserById(ctx context.Context, u models.User, id int) (user models.User, err error) {
	if id <= 0 {
		return user, models.ErrBadRequest
	}

	user, err = a.db.GetUserById(ctx, id)
	if err != nil {
		if err != models.ErrNoRows {
			log.Println("app GetUser", err)
		}
		return
	}
	if id != u.ID || !u.IsAdmin() { // remove sensitive data if user is not self or admin
		user.ID = 0
		user.Role = ""
	}
	return
}

func (a *App) GetAllUsers(ctx context.Context, u models.User) (users []models.User, err error) {
	users, err = a.db.GetAllUsers(ctx)
	if err != nil {
		if err != models.ErrNoRows {
			log.Println("app GetAllUsers", err)
		}
		return
	}
	if !u.IsAdmin() { // remove sensitive data for non-admins
		for _, user := range users {
			if u.ID != user.ID { // don't remove for self
				user.ID = 0
				user.Role = ""
			}
		}
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

func (a *App) DeleteUser(ctx context.Context, u models.User, id int) (err error) {
	if !u.IsAdmin() {
		return models.ErrUnauthorized
	}
	if id <= 0 {
		return models.ErrBadRequest
	}

	err = a.db.DeleteUser(ctx, id)
	if err != nil {
		log.Println("app DeleteUser", err)
	}
	return
}

func (a *App) UpdateUser(ctx context.Context, u, user models.User) (err error) {
	if !user.Validate() || user.ID <= 0 {
		return models.ErrBadRequest
	}
	if !u.IsAdmin() { // admins can update all users
		if user.ID != u.ID { // users can update only themselves
			return models.ErrUnauthorized
		}
		if user.Role == models.RoleAdmin { // users can't give themselves admin roles
			return models.ErrUnauthorized
		}
	}

	err = a.db.UpdateUser(ctx, user)
	if err != nil {
		log.Println("app UpdateUser", err)
	}
	return
}
