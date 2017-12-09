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
	router.GET("/api/v1/zeiten_0_100", apiv1zeiten0100)
	router.GET("/api/v1/zeiten_50_150", apiv1zeiten50150)
	router.GET("/api/v1/test", testfunc)

	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))

}

type Test struct {
	BasisKFZID int64 `json:"id"`
	HerstellerID int64 `json:"hersteller_id"`
	KFZName string `json:"kfz_name"`
	Herstellungsjahr int64 `json:"herstellungsjahr"`
	hersteller `json:"hersteller"`
}

type hersteller struct {
	KFZID int64 `json:"id"`
	HerstellerName string `json:"hersteller_name"`
}

func testfunc (response http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	database, dberror := sql.Open("postgres", "user=cj dbname=jp-zeiten sslmode=disable")
	if dberror != nil {
		log.Fatal(dberror)
	}

	rows, qerror := database.Query("SELECT * FROM basis_kfz INNER JOIN kfz_hersteller ON basis_kfz.hersteller = kfz_hersteller.id")
	if qerror != nil {
		log.Fatal(qerror)
	}

	asmdk := make([]*Test, 0)

	for rows.Next() {
		newitem := new(Test)
		err := rows.Scan(&newitem.BasisKFZID, &newitem.HerstellerID, &newitem.KFZName, &newitem.Herstellungsjahr, &newitem.hersteller.KFZID, &newitem.hersteller.HerstellerName)

		if err != nil {
			log.Fatal(err)
		}

		asmdk = append(asmdk, newitem)
	}

	jsonify, _ := json.MarshalIndent(asmdk, "", " ")
	response.Write(jsonify)

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
	apijson, _ := json.Marshal(zeitenArray)
	response.Write(apijson)
}


func apiv1zeiten0100(response http.ResponseWriter, request *http.Request, _ httprouter.Params) {

}


func apiv1zeiten50150(response http.ResponseWriter, request *http.Request, _ httprouter.Params) {

}
