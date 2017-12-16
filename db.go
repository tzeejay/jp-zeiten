package main

import (
	"database/sql"
	"log"
)

var Database *sql.DB

func postgresOnline() (sucess bool) {
	database, dberror := sql.Open("postgres", "user=cj dbname=jp-zeiten sslmode=disable")
	if dberror != nil {
		log.Fatal(dberror)
	}
	Database = database
	
	return true
}