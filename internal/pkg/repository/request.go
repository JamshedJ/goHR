package repository

import (
	"context"

	"github.com/JamshedJ/goHR/internal/models"
	"github.com/jackc/pgx/v5"
)

func (d *DB) CreateEmployeeRequest(ctx context.Context, r models.EmployeeRequest) (id int, err error) {
	err = d.conn.QueryRow(ctx,
		`INSERT INTO employee_requests (employee_id, starts_at, ends_at, reason, employee_request_type_id) 
		VAlUES ($1, $2, $3, $4, $5) RETURNING id;`,
		r.EmployeeID, r.StartsAt, r.EndsAt, r.Reason, r.EmployeeRequestTypeID).Scan(&id)
	return
}

func (d *DB) GetRequestByID(ctx context.Context, id int) (e models.EmployeeRequest, err error) {
	if err = d.conn.QueryRow(ctx,
		`SELECT
			id,
			employee_id,
			TO_CHAR(starts_at, 'YYYY-MM-DD HH24:MI:SS'),
			TO_CHAR(ends_at, 'YYYY-MM-DD HH24:MI:SS'),
			reason,
			employee_request_type_id
		FROM employee_requests
		WHERE id = $1;`, id).Scan(
		&e.ID,
		&e.EmployeeID,
		&e.StartsAt,
		&e.EndsAt,
		&e.Reason,
		&e.EmployeeRequestTypeID,
	); err == pgx.ErrNoRows {
		err = models.ErrNoRows
	}
	return
}

func (d *DB) GetAllRequests(ctx context.Context) (requests []models.EmployeeRequest, err error) {
	rows, err := d.conn.Query(ctx,
		`SELECT
			id,
			employee_id,
			TO_CHAR(starts_at, 'YYYY-MM-DD HH24:MI:SS'),
			TO_CHAR(ends_at, 'YYYY-MM-DD HH24:MI:SS'),
			reason,
			employee_request_type_id
		FROM employee_requests;`)
	if err != nil {
		return
	}
	defer rows.Close()

	requests = make([]models.EmployeeRequest, 0)
	for rows.Next() {
		var e models.EmployeeRequest
		if err = rows.Scan(
			&e.ID,
			&e.EmployeeID,
			&e.StartsAt,
			&e.EndsAt,
			&e.Reason,
			&e.EmployeeRequestTypeID,
		); err != nil {
			return
		}
		requests = append(requests, e)
	}
	return
}

func (d *DB) UpdateEmployeeRequest(ctx context.Context, e models.EmployeeRequest) (err error) {
	result, err := d.conn.Exec(ctx, `
		UPDATE employee_requests SET
			employee_id = $2,
			starts_at = $3,
			ends_at = $4,
			reason = $5,
			employee_request_type_id = $6
		WHERE id = $1;`,
		e.ID,
		e.EmployeeID,
		e.StartsAt,
		e.EndsAt,
		e.Reason,
		e.EmployeeRequestTypeID,
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
	result, err := d.conn.Exec(ctx, `DELETE FROM employee_requests WHERE id = $1;`, id)
	if err != nil {
		return
	}
	if result.RowsAffected() != 1 {
		return models.ErrNoRows
	}
	return
}
