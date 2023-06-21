package repository

import (
	"context"
	"fmt"
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

	createUsersTable = `CREATE TABLE IF NOT EXISTS users
	(
		id              SERIAL 			PRIMARY KEY,
    	username        VARCHAR(255) 	NOT NULL,
    	password        VARCHAR(255) 	NOT NULL,
    	role            VARCHAR(50) 	DEFAULT 'user' -- admin, user
	);`

	createEmployeeRequestTypesTable = `CREATE TABLE IF NOT EXISTS employee_request_types
	(
		id              SERIAL 			PRIMARY KEY,
    	title           VARCHAR(255)  --(sick_leave or vacation)
	);`

	createEmployeeRequestsTable = `CREATE TABLE IF NOT EXISTS employee_requests
	(
		id              SERIAL 			PRIMARY KEY,
    	employee_id     INT 			REFERENCES employees(id),
    	starts_at      DATE 			NOT NULL,
    	ends_at        DATE 			NOT NULL,
    	reason          VARCHAR(255),
    	type_id         INT 			NOT NULL REFERENCES employee_request_types(id)
	);`
)

var createTable = map[string]string{
	"departments":            createDepartmentsTable,
	"positions":              createPositionsTable,
	"employees":              createEmployeesTable,
	"users":                  createUsersTable,
	"employee_request_types": createEmployeeRequestTypesTable,
	"employee_requests":      createEmployeeRequestsTable,
}

func (d *DB) Up(ctx context.Context) error {
	for tableName, query := range createTable {
		_, err := d.conn.Exec(ctx, query)
		if err != nil {
			return fmt.Errorf("error occurred while creating table %q: %w", tableName, err)
		}
	}
	return nil
}
