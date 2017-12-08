package main

import (
	"database/sql"
	_ "errors"
	_"fmt"
	_ "io"
	"log"
	"net/http"
	"encoding/json"
	_ "strings"
	_ "github.com/lib/pq"
	"github.com/julienschmidt/httprouter"
)

var database *sql.DB

func main() {

	router := httprouter.New()
	router.GET("/api/v1/zeiten_100_200", apiv1Zeiten100200)
	//router.GET("/api/v1/tuning/:id", lazyLoadTuning)
	//router.Get("/api/v1/")

	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))

}

type KfzZeiten struct {
  ZeitenId sql.NullInt64 `json:"zeiten-id"`
	KfzVariante sql.NullInt64 `json:"kfz_variante"`
	Nass sql.NullInt64 `json:"nass"`
	GemesseneZeit sql.NullFloat64 `json:"gemessene_zeit"`
	YoutubeURL sql.NullString `json:"youtube_url"`
  VarianteId sql.NullInt64 `json:"kfz_variante_id"`
  SerieAbWerk sql.NullBool `json:"serie_ab_werk"`
  PS sql.NullInt64 `json:"ps"`
  NM sql.NullInt64 `json:"nm"`
  Tuning sql.NullInt64 `json:"tuning"`
  BasisKfzId sql.NullInt64 `json:"basis_kfz"`
  KfzName sql.NullString `json:"kfz_name"`
  TuningId sql.NullInt64 `json:"tuning"`
  SerienKfz sql.NullInt64 `json:"serien_kfz"`
  TuningName sql.NullString `json:"tuning_name"`
}

func apiv1Zeiten100200(response http.ResponseWriter, request *http.Request, _ httprouter.Params) {

	database, dberror := sql.Open("postgres", "user=cj dbname=jp-zeiten sslmode=disable")
  if dberror != nil {
    log.Panic()
  }

	rows, qerror := database.Query("SELECT zeiten_100_200.id, zeiten_100_200.kfz_variante, zeiten_100_200.nass, zeiten_100_200.gemessene_zeit, zeiten_100_200.youtube_url, kfz_variante.id, kfz_variante.serie_ab_werk, kfz_variante.ps, kfz_variante.nm, kfz_variante.tuning, basis_kfz.id, basis_kfz.kfz_name, tuning.id, tuning.serien_kfz, tuning.tuning_name FROM zeiten_100_200 INNER JOIN kfz_variante ON zeiten_100_200.kfz_variante =  kfz_variante.id INNER JOIN basis_kfz ON kfz_variante.serien_kfz = basis_kfz.id LEFT OUTER JOIN tuning ON kfz_variante.tuning = tuning.id")
  if qerror != nil {
    log.Fatal(qerror)
  }

	zeitenArray := make([]*KfzZeiten, 0)

  for rows.Next() {
    queriedTime := new(KfzZeiten)

    err := rows.Scan(&queriedTime.ZeitenId, &queriedTime.KfzVariante, &queriedTime.Nass, &queriedTime.GemesseneZeit, &queriedTime.YoutubeURL, &queriedTime.VarianteId, &queriedTime.SerieAbWerk, &queriedTime.PS, &queriedTime.NM, &queriedTime.Tuning, &queriedTime.BasisKfzId, &queriedTime.KfzName, &queriedTime.TuningId, &queriedTime.SerienKfz, &queriedTime.TuningName)
    if err != nil {
      log.Fatal(err)
    }
		zeitenArray = append(zeitenArray, queriedTime)
  }
	json, _ := json.Marshal(zeitenArray)
	response.Write(json)
}
