package main

import "database/sql"

type zeiten struct {
	Id int64 `json:"id"`
	KfzVariante int64 `json:"kfz_variante"`
	Zeituhr int64 `json:"zeituhr"`
	Nass int64 `json:"nass"`
	GemesseneZeit float64 `json:"gemessene_zeit"`
	YoutubeURL string `json:"youtube_url"`
}

type Zeiten struct {
	KFZVariante int64 `json:"kfz_variante"`
	GemesseneZeit float64 `json:"gemessene_zeit"`
	KFZVarianteID int64 `json:"-"`
	SerienKFZ int64 `json:"serien_kfz"`
	SerienKFZID int64 `json:"-"`
	KFZName string  `json:"kfz_name"`
	Fabrikationsjahr int64 `json:"fabrikationsjahr"`
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

// type zeiten struct {
// 	Id int64 `json:"id"`
// 	KfzVariante int64 `json:"kfz_variante"`
// 	Nass int64 `json:"nass"`
// 	GemesseneZeit float64 `json:"gemessene_zeit"`
//
//
// }

type basisKFZ struct {
	Id int64 `json:"id"`
	Hersteller int64 `json:"hersteller"`
	KFZName string `json:"kfz_name"`
	Fabrikationsjahr int64 `json:"fabrikationsjahr"`
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
