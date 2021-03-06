package main

import (
	_"database/sql"
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



func main() {
	
	// Initial check to see if the database is available at startup
	// before we add the routes
	if success := postgresOnline(); success != true {
		log.Fatal("Db is offline")
	}
	
	// Adding routes
	router := httprouter.New()
	
	// GETs
	router.GET("/api/v1/zeiten-100-200", apiv1GetZeiten100200)
	router.GET("/api/v1/zeiten-0-100", apiv1GetZeiten0100)
	router.GET("/api/v1/zeiten-50-150", apiv1GetZeiten50150)
	router.GET("/api/v1/kfz-hersteller", apiv1GetHersteller)
	router.GET("/api/v1/basis-kfz", apiv1GetBasisKFZ)
	
	// POSTs
	router.POST("/api/v1/kfz-hersteller", apiv1AddHersteller)
	//router.GET("/api/v1/test", testfunc)

	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))

}


/* curl localhost:8080/api/v1/zeiten-100-200 */
func apiv1GetZeiten100200(response http.ResponseWriter, request *http.Request, _ httprouter.Params) {

	rows, qerror := Database.Query("SELECT t1.kfz_variante, t1.gemessene_zeit, t2.id, t2.serien_kfz, t3.id, t3.kfz_name, t3.fabrikationsjahr FROM zeiten_100_200 AS t1 INNER JOIN kfz_variante AS t2 ON t1.kfz_variante = t2.id INNER JOIN basis_kfz AS t3 ON t2.serien_kfz = t3.id")
  	if qerror != nil {
    	log.Fatal(qerror)
  	}

	zeitenArray := make([]*Zeiten, 0)

  	for rows.Next() {
    	queriedTime := new(Zeiten)

    	if err := rows.Scan(&queriedTime.KFZVariante, &queriedTime.GemesseneZeit, &queriedTime.KFZVarianteID, &queriedTime.SerienKFZ, &queriedTime.SerienKFZID, &queriedTime.KFZName, &queriedTime.Fabrikationsjahr); err != nil {
      		log.Fatal(err)
		}
		zeitenArray = append(zeitenArray, queriedTime)
	}
	apijson, _ := json.Marshal(zeitenArray)
	response.Write(apijson)
}


/* curl localhost:8080/api/v1/zeiten-0-100 */
func apiv1GetZeiten0100(response http.ResponseWriter, request *http.Request, _ httprouter.Params) {

	rows, qerror := Database.Query("SELECT t1.kfz_variante, t1.gemessene_zeit, t2.id, t2.serien_kfz, t3.id, t3.kfz_name, t3.fabrikationsjahr FROM zeiten_0_100 AS t1 INNER JOIN kfz_variante AS t2 ON t1.kfz_variante = t2.id INNER JOIN basis_kfz AS t3 ON t2.serien_kfz = t3.id")
	if qerror != nil {
	   	log.Fatal(qerror)
	}

	zeitenArray := make([]*Zeiten, 0)

	for rows.Next() {
		queriedTime := new(Zeiten)

	    if err := rows.Scan(&queriedTime.KFZVariante, &queriedTime.GemesseneZeit, &queriedTime.KFZVarianteID, &queriedTime.SerienKFZ, &queriedTime.SerienKFZID, &queriedTime.KFZName, &queriedTime.Fabrikationsjahr); err != nil {
	      		log.Fatal(err)
	    }
		zeitenArray = append(zeitenArray, queriedTime)
	}
	apijson, _ := json.Marshal(zeitenArray)
	response.Write(apijson)
}


/* curl localhost:8080/api/v1/zeiten-50-150 */
func apiv1GetZeiten50150(response http.ResponseWriter, request *http.Request, _ httprouter.Params) {

	rows, qerror := Database.Query("SELECT t1.kfz_variante, t1.gemessene_zeit, t2.id, t2.serien_kfz, t3.id, t3.kfz_name, t3.fabrikationsjahr FROM zeiten_0_100 AS t1 INNER JOIN kfz_variante AS t2 ON t1.kfz_variante = t2.id INNER JOIN basis_kfz AS t3 ON t2.serien_kfz = t3.id")
	if qerror != nil {
	   	log.Fatal(qerror)
	}

	zeitenArray := make([]*Zeiten, 0)

  	for rows.Next() {
		queriedTime := new(Zeiten)

		if err := rows.Scan(&queriedTime.KFZVariante, &queriedTime.GemesseneZeit, &queriedTime.KFZVarianteID, &queriedTime.SerienKFZ, &queriedTime.SerienKFZID, &queriedTime.KFZName, &queriedTime.Fabrikationsjahr); err != nil {
			log.Fatal(err)
		}
		zeitenArray = append(zeitenArray, queriedTime)
	}
	apijson, _ := json.Marshal(zeitenArray)
	response.Write(apijson)
}


/* curl -X POST localhost:8080/api/v1/kfz-hersteller/ -d "{\"kfz_hersteller\":\"Herstller\"}" */
func apiv1AddHersteller(response http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	
	// Just in case we'd like to keep track of 
	var ada map[string]interface{}
	err := json.NewDecoder(request.Body).Decode(&ada)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(ada["kfz_hersteller"])
	
	var insertedRowId int	
	insertError := Database.QueryRow("INSERT INTO kfz_hersteller (hersteller_name) VALUES ($1) RETURNING id", ada["kfz_hersteller"]).Scan(&insertedRowId)
	if insertError != nil {
		log.Fatal(insertError)
	}
	log.Println(insertedRowId)
	
	response.WriteHeader(http.StatusCreated)
}


/* curl localhost:8080/api/v1/kfz-hersteller */
func apiv1GetHersteller(response http.ResponseWriter, request *http.Request, _ httprouter.Params) {
		
	rows, queryError := Database.Query("SELECT * FROM kfz_hersteller")
	if queryError != nil {
		log.Fatal(queryError)
	}
	
	herstellerArray := make([]*KFZHersteller, 0)
	
	for rows.Next() {
		hersteller := new(KFZHersteller)
		
		if err := rows.Scan(&hersteller.Id, &hersteller.HerstellerName); err != nil {
			log.Fatal(err)
		}
		// You solved this a second time in this project alone.
		// If it's yelling at you check if you use := or simply = 
		herstellerArray = append(herstellerArray, hersteller)	
	}
	apijson, _ := json.Marshal(herstellerArray)
	response.Write(apijson)
}


/* curl http://localhost:8080/api/v1/basis_kfz */
func apiv1GetBasisKFZ(response http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	
	rows, queryError := Database.Query("SELECT * FROM basis_kfz")
	if queryError != nil {
		log.Fatal(queryError)
	}
	
	kfzArray := make([]*BasisKFZ, 0)
	
	for rows.Next() {
		kfz := new(BasisKFZ)
		
		if err := rows.Scan(&kfz.Id, &kfz.Hersteller, &kfz.KFZName, &kfz.Fabrikationsjahr); err != nil {
			log.Fatal(err)
		}
		
		kfzArray = append(kfzArray, kfz)
	}
	apijson, _ := json.Marshal(kfzArray)
	response.Write(apijson)
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