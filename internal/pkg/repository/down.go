package repository

import (
	"context"
	"fmt"
)

const (
	dropDepartmentsTable          = `DROP TABLE IF EXIST departments;`
	dropPositionsTable            = `DROP TABLE IF EXIST positions;`
	dropEmployeesTable            = `DROP TABLE IF EXIST employees;`
	dropUsersTable                = `DROP TABLE IF EXIST users;`
	dropEmployeeRequestTypesTable = `DROP TABLE IF EXIST employee_request_types;`
	dropEmployeeRequestsTable     = `DROP TABLE IF EXIST employee_requests;`
)

var dropTables = map[string]string{
	"departments":            dropDepartmentsTable,
	"positions":              dropPositionsTable,
	"employees":              dropEmployeesTable,
	"users":                  dropUsersTable,
	"employee_request_types": dropEmployeeRequestTypesTable,
	"employee_requests":      dropEmployeeRequestsTable,
}

func (d *DB) Down(ctx context.Context) error {
	for tableName, query := range dropTables {
		_, err := d.conn.Exec(ctx, query)
		if err != nil {
			return fmt.Errorf("error occurred while dropping table %q: %w", tableName, err)
		}
	}
	return nil
}
