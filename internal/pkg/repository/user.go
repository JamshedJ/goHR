package repository

import (
	"context"

	"github.com/JamshedJ/goHR/internal/models"
	"github.com/jackc/pgx/v5"
)

func (d *DB) GetUserById(ctx context.Context, id int) (user models.User, err error) {
	if err = d.conn.QueryRow(ctx, `SELECT id, username, role FROM users WHERE id = $1;`, id).
		Scan(&user.ID, &user.Username, &user.Role); err == pgx.ErrNoRows {
		err = models.ErrNoRows
	}
	return
}

func (d *DB) GetAllUsers(ctx context.Context) (users []models.User, err error) {
	rows, err := d.conn.Query(ctx,
		`SELECT id, username, role FROM users;`)
	if err != nil {
		return
	}
	defer rows.Close()

	users = make([]models.User, 0)
	for rows.Next() {
		var u models.User
		if err = rows.Scan(&u.ID, &u.Username, &u.Role); err != nil {
			return
		}
		users = append(users, u)
	}
	return
}

func (d *DB) CreateUser(ctx context.Context, u models.User) (err error) {
	result, err := d.conn.Exec(ctx, `INSERT INTO users (username, password) VALUES ($1, $2);`, u.Username, u.Password)
	if err != nil {
		return
	}
	if result.RowsAffected() != 1 {
		err = models.ErrDuplicate
	}
	return
}

func (d *DB) UpdateUser(ctx context.Context, user models.User) (err error) {
	result, err := d.conn.Exec(ctx, `UPDATE users SET username = $2, password = $3, role = $4 WHERE id = $1;`,
		user.ID, user.Username, user.Password, user.Role)
	if err != nil {
		return
	}
	if result.RowsAffected() != 1 {
		return models.ErrNoRows
	}
	return
}

func (d *DB) DeleteUser(ctx context.Context, id int) (err error) {
	result, err := d.conn.Exec(ctx, `DELETE FROM users WHERE id = $1;`, id)
	if err != nil {
		return
	}
	if result.RowsAffected() != 1 {
		return models.ErrNoRows
	}
	return
}

func (d *DB) AuthenticateUser(ctx context.Context, u *models.User) (err error) {
	if err = d.conn.QueryRow(ctx,
		`SELECT id, role FROM users WHERE username = $1 AND password = $2;`,
		u.Username, u.Password).Scan(&u.ID, &u.Role); err == pgx.ErrNoRows {
		err = models.ErrNoRows
	}
	u.Password = ""
	return
}
