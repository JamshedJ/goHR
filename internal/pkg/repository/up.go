package repository

import (
	"context"
	"fmt"
	"github.com/JamshedJ/goHR/internal/models"
	"github.com/JamshedJ/goHR/internal/configs"
)

const (
	createDepartmentsTable = `CREATE TABLE IF NOT EXISTS departments
	(
		id              SERIAL 			PRIMARY KEY,
    	title           VARCHAR(255) 	NOT NULL,
    	department_head VARCHAR(255)
	);`

	createPositionsTable = `CREATE TABLE IF NOT EXISTS positions
	(
		id              SERIAL 			PRIMARY KEY,
    	title           VARCHAR(255) 	NOT NULL,
    	salary          DECIMAL(10, 2) 	NOT NULL,
    	qualification   VARCHAR(255)
	);`

	createEmployeesTable = `CREATE TABLE IF NOT EXISTS employees
	(
		id              SERIAL 			PRIMARY KEY,
    	first_name      VARCHAR(255) 	NOT NULL,
    	last_name       VARCHAR(255) 	NOT NULL,
    	position_id     INT REFERENCES 	positions(id),
    	department_id   INT REFERENCES 	departments(id),
    	employment_date DATE 			NOT NULL,
    	salary          DECIMAL(10, 2) 	NOT NULL
	);`

	createEmployeeRequestTypesTable = `CREATE TABLE IF NOT EXISTS employee_request_types
	(
		id              SERIAL 			PRIMARY KEY,
    	title           VARCHAR(255)  --(sick_leave or vacation)
	);`

	createEmployeeRequestsTable = `CREATE TABLE IF NOT EXISTS employee_requests
	(
		id              			SERIAL 			PRIMARY KEY,
    	employee_id     			INT 			REFERENCES employees(id),
    	starts_at       			DATE 			NOT NULL,
    	ends_at         			DATE 			NOT NULL,
    	reason          			VARCHAR(255),
    	employee_request_type_id    INT 			NOT NULL REFERENCES employee_request_types(id)
	);`

	createUsersTable = `CREATE TABLE IF NOT EXISTS users (
			id              SERIAL 			PRIMARY KEY,
			username        VARCHAR(255) 	NOT NULL UNIQUE,
			password        VARCHAR(255) 	NOT NULL,
			role            VARCHAR(50) 	DEFAULT 'user' -- admin, user
		);`
)

func (d *DB) Up(ctx context.Context) error {
	var commandList = []map[string]string{
		{"create_table_departments": createDepartmentsTable},
		{"create_table_positions": createPositionsTable},
		{"create_table_employees": createEmployeesTable},
		{"create_table_employee_request_types": createEmployeeRequestTypesTable},
		{"create_table_employee_requests": createEmployeeRequestsTable},
		{"create_user_admin_default": createUsersTable},

		{"create_default_admin": fmt.Sprintf(
			`INSERT INTO users (username, password, role) 
			VALUES ('%s', '%s', 'admin')
			ON CONFLICT (username) DO NOTHING;`,
			configs.App.AdminUsername, models.GeneratePasswordHash(configs.App.AdminPassword))},
	}

	for _, command := range commandList {
		for name, query := range command {
			_, err := d.conn.Exec(ctx, query)
			if err != nil {
				return fmt.Errorf("error occurred while running %q: %w", name, err)
			}
		}
	}
	return nil
}
