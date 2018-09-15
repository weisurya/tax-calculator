// Package common - query. This file contains several functions that query-related
package common

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// CreateInsertStatement function uses to create standarize sql statement for inserting purpose
func CreateInsertStatement(table, field, value string) string {
	return fmt.Sprintf("INSERT INTO %v(%v) VALUES(%v);", table, field, value)
}

// ExecuteSQL function uses to create standarize sql statement for executing purpose
func ExecuteSQL(database *sql.DB, sqlStatement string) (err error) {
	if _, err := database.Exec(sqlStatement); err != nil {
		return err
	}
	return nil
}
