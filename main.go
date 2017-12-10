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
	//router.GET("/api/v1/test", testfunc)

	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))

}

/*func testfunc (response http.ResponseWriter, request *http.Request, _ httprouter.Params) {
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

}*/


func apiv1Zeiten100200(response http.ResponseWriter, request *http.Request, _ httprouter.Params) {

	database, dberror := sql.Open("postgres", "user=cj dbname=jp-zeiten sslmode=disable")
	if dberror != nil {
		log.Fatal(dberror)
  	}

	rows, qerror := database.Query("SELECT t1.kfz_variante, t1.gemessene_zeit, t2.id, t2.serien_kfz, t3.id, t3.kfz_name, t3.herstellungsjahr FROM zeiten_100_200 AS t1	INNER JOIN kfz_variante AS t2 ON t1.kfz_variante = t2.id	INNER JOIN basis_kfz AS t3 ON t2.serien_kfz = t3.id")
  if qerror != nil {
    log.Fatal(qerror)
  }

	zeitenArray := make([]*Zeiten, 0)

  for rows.Next() {
    queriedTime := new(Zeiten)

    err := rows.Scan(&queriedTime.KFZVariante, &queriedTime.GemesseneZeit, &queriedTime.KFZVarianteID, &queriedTime.SerienKFZ, &queriedTime.SerienKFZID, &queriedTime.KFZName, &queriedTime.Herstellungsjahr)
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
