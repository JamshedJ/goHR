package repository

import (
	"context"
	"fmt"
)

const (
	dropUsersTable                = `DROP TABLE IF EXISTS users;`
	dropEmployeeRequestsTable     = `DROP TABLE IF EXISTS employee_requests;`
	dropEmployeeRequestTypesTable = `DROP TABLE IF EXISTS employee_request_types;`
	dropEmployeesTable            = `DROP TABLE IF EXISTS employees;`
	dropPositionsTable            = `DROP TABLE IF EXISTS positions;`
	dropDepartmentsTable          = `DROP TABLE IF EXISTS departments;`
)

func (d *DB) Down(ctx context.Context) error {
	var dropTables = []map[string]string{
		{"users": dropUsersTable},
		{"employee_requests": dropEmployeeRequestsTable},
		{"employee_request_types": dropEmployeeRequestTypesTable},
		{"employees": dropEmployeesTable},
		{"positions": dropPositionsTable},
		{"departments": dropDepartmentsTable},
	}

	for _, drop := range dropTables {
		for name, query := range drop {
			_, err := d.conn.Exec(ctx, query)
			if err != nil {
				return fmt.Errorf("error occurred while dropping table %q: %w", name, err)
			}
		}
	}
	return nil
}
