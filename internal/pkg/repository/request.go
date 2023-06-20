package repository

// import (
// 	"context"

// 	"github.com/JamshedJ/goHR/internal/models"
// 	"github.com/jackc/pgx/v5"
// )

// func (d *DB) CreateRequest(ctx context.Context, r models.Request) (id int, err error) {
// 	err = d.conn.QueryRow(ctx,
// 		`INSERT INTO requests (first_name, last_name, position, department, employment_date, salary) 
// 		VAlUES ($1, $2, $3, $4, $5, $6) RETURNING id;`,
// 		e.FirstName, e.LastName, e.PositionID, e.DepartmentID, e.EmploymentDate, e.Salary).Scan(&id)
// 	return
// }

// func (d *DB) GetEmployeeByID(ctx context.Context, id int) (e models.Employee, err error) {
// 	if err = d.conn.QueryRow(ctx,
// 		`SELECT
// 			id,
// 			first_name,
// 			last_name,
// 			position_id,
// 			department_id,
// 			employment_date,
// 			salary
// 		FROM employees
// 		WHERE id = $1;`, id).Scan(
// 		&e.ID,
// 		&e.FirstName,
// 		&e.LastName,
// 		&e.PositionID,
// 		&e.DepartmentID,
// 		&e.EmploymentDate,
// 		&e.Salary,
// 	); err == pgx.ErrNoRows {
// 		err = models.ErrNoRows
// 	}
// 	return
// }

// func (d *DB) GetEmployees(ctx context.Context) (employees []models.Employee, err error) {
// 	rows, err := d.conn.Query(ctx,
// 		`SELECT
// 			id,
// 			first_name,
// 			last_name,
// 			position_id,
// 			department_id,
// 			employment_date,
// 			salary
// 		FROM employees;`)
// 	if err != nil {
// 		return
// 	}
// 	defer rows.Close()

// 	employees = make([]models.Employee, 0)
// 	for rows.Next() {
// 		var e models.Employee
// 		if err = rows.Scan(&e.ID,
// 			&e.FirstName,
// 			&e.LastName,
// 			&e.PositionID,
// 			&e.DepartmentID,
// 			&e.EmploymentDate,
// 			&e.Salary); err != nil {
// 			return
// 		}
// 		employees = append(employees, e)
// 	}
// 	return
// }

// func (d *DB) UpdateEmployee(ctx context.Context, e models.Employee) (err error) {
// 	result, err := d.conn.Exec(ctx, `
// 		UPDATE employees SET
// 			first_name = $2,
// 			last_name = $3,
// 			position_id = $4,
// 			department_id = $5,
// 			employment_date = $6,
// 			salary = $7
// 		WHERE id = $1;`,
// 		e.ID,
// 		e.FirstName,
// 		e.LastName,
// 		e.PositionID,
// 		e.DepartmentID,
// 		e.EmploymentDate,
// 		e.Salary,
// 	)
// 	if err != nil {
// 		return
// 	}
// 	if result.RowsAffected() != 1 {
// 		return models.ErrNoRows
// 	}
// 	return
// }

// func (d *DB) DeleteEmployee(ctx context.Context, id int) (err error) {
// 	result, err := d.conn.Exec(ctx, `DELETE FROM employees WHERE id = $1;`, id)
// 	if err != nil {
// 		return
// 	}
// 	if result.RowsAffected() != 1 {
// 		return models.ErrNoRows
// 	}
// 	return
// }
