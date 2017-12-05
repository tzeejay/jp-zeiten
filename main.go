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

type zeiten struct {
	Id int64 `json:"id"`
	KfzVariante int64 `json:"kfz_variante"`
	Zeituhr int64 `json:"zeituhr"`
	Nass int64 `json:"nass"`
	GemesseneZeit float64 `json:"gemessene_zeit"`
	YoutubeURL string `json:"youtube_url"`
}

type basisKFZ struct {
	Id int64 `json:"id"`
	Hersteller int64 `json:"hersteller"`
	KFZName string `json:"kfz_name"`
	Herstellungsjahr int64 `json:"herstellungsjahr"`
}

type getriebe struct {
	Id int64 `json:"id"`
	GetriebeBezeichnung string `json:"getriebe_bezeichnung"`
	GanzAnzahl int64 `json:"gang_anzahl"`
}

type KFZHersteller struct {
	Id int64 `json:"id"`
	HerstellerName string `json:"hersteller_name"`
}

type KFZVariante struct {
	Id int64 `json:"id"`
	SerieAbWerk bool `json:"serie_ab_werk"`
	SerienKFZ int64 `json:"serien_kfz"`
	PS int64 `json:"ps"`
	PSGemessen bool `json:"ps_gemessen"`
	NM int64 `json:"nm"`
	NMGemessen bool `json:"nm_gemessen"`
	Gewicht int64 `json:"gewicht"`
	GewichtGemessen bool `json:"gewicht_gemessen"`
	Tuning int64 `json:"tuning"`
	Getriebe int64 `json:"getriebe"`
}

type tuning struct {
	Id int64 `json:"id"`
	SerienKFZ int64 `json:"serien_kfz"`
	TuningName string `json:"tuning_name"`
	TuningBeschreibung string `json:"tuning_beschreibung"`
	Ansaugung bool `json:"ansaugung"`
	Abgasanlage bool `json:"abgasanlage"`
	Ladedruck bool `json:"ladedruck"`
	Fahrwerk bool  `json:"fahrwerk"`
	Luftfahrwerk bool `json:"luftfahrwerk"`
	Bremsanlage bool `json:"bremsanlage"`
	Getriebe bool `json:"getriebe"`
	Software bool `json:"software"`
	GewichtsVerringerung bool `json:"gewicht_verringerung"`
	Reifen bool `json:"reifen"`
	Karosserie bool `json:"karosserie"`
	Aerodynamik bool `json:"aerodynamik"`
	Plus10PSProSticker bool `json:"plus_10_ps_pro_sticker"`
	StickerAnzahl int64 `json:"sticker_anzahl"`
}

func main() {

	router := httprouter.New()
	router.GET("/api/v1/zeiten_100_200", apiv1Zeiten100200)

	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))

}

func apiv1Zeiten100200(response http.ResponseWriter, request *http.Request, _ httprouter.Params) {

	database, dberror := sql.Open("postgres", "user=cj dbname=jp-zeiten sslmode=disable")
  if dberror != nil {
    log.Panic()
  }

	rows, qerror := database.Query("SELECT * FROM zeiten_100_200")
  if qerror != nil {
    log.Fatal(qerror)
  }

	zeitenArray := make([]*zeiten, 0)

  for rows.Next() {
    queriedTime := new(zeiten)

    err := rows.Scan(&queriedTime.Id, &queriedTime.KfzVariante, &queriedTime.Zeituhr, &queriedTime.Nass, &queriedTime.GemesseneZeit, &queriedTime.YoutubeURL)
    if err != nil {
      log.Fatal(err)
    }
		zeitenArray = append(zeitenArray, queriedTime)
  }
	json, _ := json.MarshalIndent(zeitenArray, "", " ")
	response.Write(json)
}
