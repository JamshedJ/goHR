package repository

import (
	"context"

	"github.com/JamshedJ/goHR/internal/models"
	"github.com/jackc/pgx/v5"
)

func (d *DB) CreateRequest(ctx context.Context, r models.Request) (id int, err error) {
	err = d.conn.QueryRow(ctx,
		`INSERT INTO requests (employee_id, start_date, end_date, reason, type_id) 
		VAlUES ($1, $2, $3, $4, $5) RETURNING id;`,
		r.EmployeeID, r.StartDate, r.EndDate, r.Reason, r.TypeID).Scan(&id)
	return
}

func (d *DB) GetRequestByID(ctx context.Context, id int) (r models.Request, err error) {
	if err = d.conn.QueryRow(ctx,
		`SELECT
			id,
			employee_id,
			start_date,
			end_date,
			reason,
			type_id
		FROM requests
		WHERE id = $1;`, id).Scan(
		&r.ID,
		&r.EmployeeID,
		&r.StartDate,
		&r.EndDate,
		&r.Reason,
		&r.TypeID,
	); err == pgx.ErrNoRows {
		err = models.ErrNoRows
	}
	return
}

func (d *DB) GetAllRequests(ctx context.Context) (requests []models.Request, err error) {
	rows, err := d.conn.Query(ctx,
		`SELECT
			id,
			employee_id,
			start_date,
			end_date,
			reason,
			type_id
		FROM requests;`)
	if err != nil {
		return
	}
	defer rows.Close()

	requests = make([]models.Request, 0)
	for rows.Next() {
		var r models.Request
		if err = rows.Scan(
			&r.ID,
			&r.EmployeeID,
			&r.StartDate,
			&r.EndDate,
			&r.Reason,
			&r.TypeID,
		); err != nil {
			return
		}
		requests = append(requests, r)
	}
	return
}

func (d *DB) UpdateRequest(ctx context.Context, r models.Request) (err error) {
	result, err := d.conn.Exec(ctx, `
		UPDATE requests SET
			employee_id = $2,
			start_date = $3,
			end_date = $4,
			reason = $5,
			type_id = $6
		WHERE id = $1;`,
		r.ID,
		r.EmployeeID,
		r.StartDate,
		r.EndDate,
		r.Reason,
		r.TypeID,
	)
	if err != nil {
		return
	}
	if result.RowsAffected() != 1 {
		return models.ErrNoRows
	}
	return
}

func (d *DB) DeleteRequest(ctx context.Context, id int) (err error) {
	result, err := d.conn.Exec(ctx, `DELETE FROM requests WHERE id = $1;`, id)
	if err != nil {
		return
	}
	if result.RowsAffected() != 1 {
		return models.ErrNoRows
	}
	return
}
