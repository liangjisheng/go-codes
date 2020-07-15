package database

import (
	"database/sql"
	"fmt"

	// postgres
	_ "github.com/lib/pq"
)

// Connect connect postgres
func Connect(user, password, dbname, host, port string) (*sql.DB, error) {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s", user, password, dbname, host, port)
	return sql.Open("postgres", connStr)
}
