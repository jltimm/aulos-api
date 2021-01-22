package postgres

import (
	"database/sql"
	"fmt"

	"../../common"
	"../../secrets"
	"github.com/lib/pq"
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

// GetAllArtists returns all artists in the database
func GetAllArtists() []common.Artist {
	selectStatement := "SELECT id, name, popularity, recommended FROM artists;"
	rows, err := db.Query(selectStatement)
	if err != nil {
		panic(err)
	}
	var artists []common.Artist
	for rows.Next() {
		var artist common.Artist
		err := rows.Scan(&artist.ID, &artist.Name, &artist.Popularity, pq.Array(&artist.Recommended))
		if err != nil {
			panic(err)
		}
		artists = append(artists, artist)
	}
	return artists
}
