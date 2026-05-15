package dbconfig

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func NewDB() (*sql.DB, error) {

	connStr := "postgres://postgres:Root@1234@localhost:5432/Test?sslmode=disable" // connStr is the connection string for the PostgreSQL database,
	// it includes the username, password, host, port, database name and sslmode
	return sql.Open("postgres", connStr) // sql.Open is a function from the database/sql package that
	// opens a connection to the database specified by the driver name and data source name (DSN).

}
