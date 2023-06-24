package repository

import (
	"context"
	"fmt"

	// "github.com/JamshedJ/goHR/internal/configs"
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
		id              			SERIAL 			PRIMARY KEY,
    	employee_id     			INT 			REFERENCES employees(id),
    	starts_at       			DATE 			NOT NULL,
    	ends_at         			DATE 			NOT NULL,
    	reason          			VARCHAR(255),
    	employee_request_type_id    INT 			NOT NULL REFERENCES employee_request_types(id)
	);`
)

// var (
// 	defaultAdminUsername = configs.AppSettings.AppParams.Login
// 	defaultAdminPassword = configs.AppSettings.AppParams.SecretKey

// 	createDefaultAdmin = fmt.Sprintf(`INSERT INTO users (username, password, role) VALUES ('%s', '%s', "admin");`,
// 		defaultAdminUsername, defaultAdminPassword)
// )

var createTable = map[string]string{
	"create_table_departments":            createDepartmentsTable,
	"create_table_positions":              createPositionsTable,
	"create_table_employees":              createEmployeesTable,
	"create_table_users":                  createUsersTable,
	"create_table_employee_request_types": createEmployeeRequestTypesTable,
	"create_table_employee_requests":      createEmployeeRequestsTable,
	// "create_user_admin_default":           createDefaultAdmin,
}

func (d *DB) Up(ctx context.Context) error {
	for name, query := range createTable {
		_, err := d.conn.Exec(ctx, query)
		if err != nil {
			return fmt.Errorf("error occurred while running %q: %w", name, err)
		}
	}
	return nil
}
