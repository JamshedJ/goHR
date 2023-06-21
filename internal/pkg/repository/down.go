package repository

import (
	"context"
	"fmt"
)

const (
	dropDepartmentsTable = `DROP TABLE IF EXIST departments`
	dropPositionsTable   = `DROP TABLE IF EXIST positions`
	dropEmployeesTable   = `DROP TABLE IF EXIST employees`
	dropUsersTable       = `DROP TABLE IF EXIST users`
	dropTypesTable       = `DROP TABLE IF EXIST types`
	dropRequestsTable    = `DROP TABLE IF EXIST requests`
)

var dropTables = []string{
	dropDepartmentsTable,
	dropPositionsTable,
	dropEmployeesTable,
	dropUsersTable,
	dropTypesTable,
	dropRequestsTable,
}

func (d *DB) Down(ctx context.Context) error {
	for i, table := range dropTables {
		_, err := d.conn.Exec(ctx, table)
		if err != nil {
			return fmt.Errorf("error occurred while dropping table №%d: %w", i, err)
		}
	}
	return nil
}
