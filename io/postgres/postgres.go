package postgres

import (
	"database/sql"
	"fmt"

	"../../secrets"
	// Import is blank because it currently is only used as a driver
	_ "github.com/lib/pq"
)

var db *sql.DB

// Close closes the database in the shutdown hook
func Close() {
	db.Close()
}

// Initialize postgres
func Initialize() {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		secrets.GetHost(), secrets.GetPort(), secrets.GetUser(), secrets.GetPassword(), secrets.GetDbname())
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
}
