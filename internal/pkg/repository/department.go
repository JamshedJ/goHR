package repository

import (
	"context"

	"github.com/JamshedJ/goHR/internal/models"
	"github.com/jackc/pgx/v5"
)

func (d *DB) CreateDepartment(ctx context.Context, dp models.Department) (id int, err error) {
	err = d.conn.QueryRow(ctx,
		`INSERT INTO departments (title, department_head)
		VALUES ($1, $2) RETURNING id;`,
		dp.Title, dp.DepartmentHead).Scan(&id)
	return
}

func (d *DB) GetDepartmentByID(ctx context.Context, id int) (dp models.Department, err error) {
	if err = d.conn.QueryRow(ctx, 
		`SELECT
			id,
			title,
			department_head
		FROM departments
		WHERE id = $1;`, id).Scan(
		&dp.ID,
		&dp.Title,
		&dp.DepartmentHead,
		); err == pgx.ErrNoRows {
			err = models.ErrNoRows
		}
	return
}

func (d *DB) GetAllDepartments(ctx context.Context) (departments []models.Department, err error) {
	rows, err := d.conn.Query(ctx, 
		`SELECT
			id,
			title,
			department_head
		FROM departments;`)
	if err != nil {
		return
	}
	defer rows.Close()

	departments = make([]models.Department, 0)
	for rows.Next() {
		var d models.Department
		if err = rows.Scan(
			&d.ID,
			&d.Title,
			&d.DepartmentHead,
		); err != nil {
			return
		}
		departments = append(departments, d)
	}
	return
}

func (d *DB) UpdateDepartment(ctx context.Context, dp models.Department) (err error) {
	result, err := d.conn.Exec(ctx, `
		UPDATE departments SET
			title = $2,
			department_head = $3
		WHERE id = $1;`,
		dp.ID,
		dp.Title,
		dp.DepartmentHead,
	)
	if err != nil {
		return
	}
	if result.RowsAffected() != 1 {
		return models.ErrNoRows
	}
	return
}

func (d *DB) DeleteDepartment(ctx context.Context, id int) (err error) {
	result, err := d.conn.Exec(ctx, `DELETE FROM departments WHERE id = $1;`, id)
	if err != nil {
		return
	}
	if result.RowsAffected() != 1 {
		return models.ErrNoRows
	}
	return
}
