package main

type zeiten struct {
	Id int64 `json:"id"`
	KfzVariante int64 `json:"kfz_variante"`
	Zeituhr int64 `json:"zeituhr"`
	Nass int64 `json:"nass"`
	GemesseneZeit float64 `json:"gemessene_zeit"`
	YoutubeURL string `json:"youtube_url"`
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
