package main

import (
	"database/sql"
	_ "errors"
	"fmt"
	_ "io"
	"log"
	"net/http"
	_ "strings"
	_ "github.com/lib/pq"
)

var database *sql.DB

func main() {

  database, dberror := sql.Open("postgres", "user=cj dbname=jp-zeiten sslmode=disable")
  if dberror != nil {
    log.Panic()
  }

  rows, qerror := database.Query("SELECT * FROM zeiten_100_200")
  if qerror != {
    log.Fatal()
  }
  for rows.Next() {
    var int id
    var kfz_variante int
    var zeituhr int
    var nass int
    var gemessene_zeit double
    var youtube_url string

    err = rows.Scan(&id, kfz_variante, &zeituhr, &nass, &gemessene_zeit, &youtube_url)
    if err != nil {
      log.Fatal(err)
    }
    fmt.Fprint(response, "%d", gemessene_zeit)
  }

	http.HandleFunc("/api/v1", apiv1)

	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func apiv1(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "Hey there!")
}
