package repository

import (
	"context"

	"github.com/JamshedJ/goHR/internal/models"
	"github.com/jackc/pgx/v5"
)

func (d *DB) CreateEmployee(ctx context.Context, e models.Employee) (id int, err error) {
	err = d.conn.QueryRow(ctx,
		`INSERT INTO employees (first_name, last_name, position, department, employment_date, salary) 
			VAlUES ($1, $2, $3, $4, $5, $6) 
			RETURNING id;`, e.FirstName, e.LastName, e.PositionID, e.DepartmentID, e.EmploymentDate, e.Salary).Scan(&id)
	return
}

func (d *DB) GetEmployeeByID(ctx context.Context, id, PositionID, DepartmentID int) (employee models.Employee, err error) {
	if err = d.conn.QueryRow(ctx,
		`SELECT
			id,
			first_name,
			last_name,
			position_id,
			department_id,
			employment_date,
			salary
		FROM employees
		WHERE id = $1 AND position_id = $2 AND department_id = $3;`, id, PositionID, DepartmentID).Scan(
		&employee.ID,
		&employee.FirstName,
		&employee.LastName,
		&employee.PositionID,
		&employee.DepartmentID,
		&employee.EmploymentDate,
		&employee.Salary,
	); err == pgx.ErrNoRows {
		err = models.ErrNoRows
	}
	return
}
