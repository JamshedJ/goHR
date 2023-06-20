package repository

import (
	"context"

	"github.com/JamshedJ/goHR/internal/models"
	"github.com/jackc/pgx/v5"
)

func (d *DB) CreatePosition(ctx context.Context, p models.Position) (id int, err error) {
	err = d.conn.QueryRow(ctx,
		`INSERT INTO positions (title, salary, qualification)
		VALUES ($1, $2, $3) RETURNING id;`,
		p.Title, p.Salary, p.Qualification).Scan(&id)
	return
}

func (d *DB) GetPositionByID(ctx context.Context, id int) (p models.Position, err error) {
	if err = d.conn.QueryRow(ctx,
		`SELECT
			id,
			title,
			salary,
			qualification
		FROM positions
		WHERE id = $1;`, id).Scan(
		&p.ID,
		&p.Title,
		&p.Salary,
		&p.Qualification,
	); err == pgx.ErrNoRows {
		err = models.ErrNoRows
	}
	return
}

func (d *DB) GetAllPositions(ctx context.Context) (positions []models.Position, err error) {
	rows, err := d.conn.Query(ctx, 
		`SELECT
			id,
			title,
			salary,
			qualification
		FROM positions;`)
	if err != nil {
		return
	}
	defer rows.Close()

	positions = make([]models.Position, 0)
	for rows.Next() {
		var p models.Position
		if err = rows.Scan(
			&p.ID,
			&p.Title,
			&p.Salary,
			&p.Qualification,
		); err != nil {
			return
		}
		positions = append(positions, p)
	}
	return
}

func (d *DB) UpdatePosition(ctx context.Context, p models.Position) (err error) {
	result, err := d.conn.Exec(ctx, `
		UPDATE positions SET
			title = $2,
			salary = $3,
			qualification = $4
		WHERE id = $1;`,
		p.ID,
		p.Title,
		p.Salary,
		p.Qualification,
	)
	if err != nil {
		return
	}
	if result.RowsAffected() != 1 {
		return models.ErrNoRows
	}
	return
}

func (d *DB) DeletePosition(ctx context.Context, id int) (err error) {
	result, err := d.conn.Exec(ctx, `DELETE FROM positions WHERE id = $1;`, id)
	if err != nil {
		return
	}
	if result.RowsAffected() != 1 {
		return models.ErrNoRows
	}
	return
}
