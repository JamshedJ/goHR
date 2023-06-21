package repository

import (
	"context"

	"github.com/JamshedJ/goHR/internal/models"
	"github.com/jackc/pgx/v5"
)

func (d *DB) CreateType(ctx context.Context, t models.Type) (id int, err error) {
	err = d.conn.QueryRow(ctx,
		`INSERT INTO types (title)
		VALUES ($1) RETURNING id;`,
		t.Title).Scan(&id)
	return
}

func (d *DB) GetTypeByID(ctx context.Context, id int) (t models.Type, err error) {
	if err = d.conn.QueryRow(ctx,
		`SELECT
			id,
			title
		FROM types
		WHERE id = $1;`, id).Scan(
		&t.ID,
		&t.Title,
	); err == pgx.ErrNoRows {
		err = models.ErrNoRows
	}
	return
}

func (d *DB) GetAllTypes(ctx context.Context) (types []models.Type, err error) {
	rows, err := d.conn.Query(ctx,
		`SELECT
			id,
			title
		FROM types;`)
	if err != nil {
		return
	}
	defer rows.Close()

	types = make([]models.Type, 0)
	for rows.Next() {
		var t models.Type
		if err = rows.Scan(
			&t.ID,
			&t.Title,
		); err != nil {
			return
		}
		types = append(types, t)
	}
	return
}

func (d *DB) UpdateType(ctx context.Context, t models.Type) (err error) {
	result, err := d.conn.Exec(ctx, `
		UPDATE types SET
			title = $2,
		WHERE id = $1;`,
		t.ID,
		t.Title,
	)
	if err != nil {
		return
	}
	if result.RowsAffected() != 1 {
		return models.ErrNoRows
	}
	return
}

func (d *DB) DeleteType(ctx context.Context, id int) (err error) {
	result, err := d.conn.Exec(ctx, `DELETE FROM types WHERE id = $1;`, id)
	if err != nil {
		return
	}
	if result.RowsAffected() != 1 {
		return models.ErrNoRows
	}
	return
}
