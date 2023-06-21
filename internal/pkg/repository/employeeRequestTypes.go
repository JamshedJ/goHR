package repository

import (
	"context"

	"github.com/JamshedJ/goHR/internal/models"
	"github.com/jackc/pgx/v5"
)

func (d *DB) CreateEmployeeRequestType(ctx context.Context, e models.EmployeeRequestType) (id int, err error) {
	err = d.conn.QueryRow(ctx, `INSERT INTO employee_request_types (title) VALUES ($1) RETURNING id;`, e.Title).Scan(&id)
	return
}

func (d *DB) GetEmployeeRequestTypeByID(ctx context.Context, id int) (e models.EmployeeRequestType, err error) {
	if err = d.conn.QueryRow(ctx, `SELECT id, title FROM employee_request_types WHERE id = $1;`, id).Scan(
		&e.ID,
		&e.Title,
	); err == pgx.ErrNoRows {
		err = models.ErrNoRows
	}
	return
}

func (d *DB) GetAllEmployeeRequestTypes(ctx context.Context) (erTypes []models.EmployeeRequestType, err error) {
	rows, err := d.conn.Query(ctx, `SELECT id, title FROM employee_request_types;`)
	if err != nil {
		return
	}
	defer rows.Close()

	erTypes = make([]models.EmployeeRequestType, 0)
	for rows.Next() {
		var t models.EmployeeRequestType
		if err = rows.Scan(
			&t.ID,
			&t.Title,
		); err != nil {
			return
		}
		erTypes = append(erTypes, t)
	}
	return
}

func (d *DB) UpdateEmployeeRequestType(ctx context.Context, e models.EmployeeRequestType) (err error) {
	result, err := d.conn.Exec(ctx, `UPDATE employee_request_types SET title = $2 WHERE id = $1;`, e.ID, e.Title)
	if err != nil {
		return
	}
	if result.RowsAffected() != 1 {
		return models.ErrNoRows
	}
	return
}

func (d *DB) DeleteEmployeeRequestType(ctx context.Context, id int) (err error) {
	result, err := d.conn.Exec(ctx, `DELETE FROM employee_request_types WHERE id = $1;`, id)
	if err != nil {
		return
	}
	if result.RowsAffected() != 1 {
		return models.ErrNoRows
	}
	return
}
