package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // Postgres driver — needed for database/sql to work with Postgres
)

// Connect opens a connection to the Postgres database and returns it.
// dbHost, dbPort, etc. would normally come from your .env file.
func Connect(dbHost, dbPort, dbUser, dbPassword, dbName string) *sql.DB {
	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName,
	)

	conn, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal("Could not open database connection:", err)
	}

	// Ping actually checks the connection works, not just that it was created
	if err := conn.Ping(); err != nil {
		log.Fatal("Could not connect to database:", err)
	}

	return conn
}
