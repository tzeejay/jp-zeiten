-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE kfz_hersteller (
	id serial PRIMARY KEY,
	hersteller_name text UNIQUE);
	
CREATE TABLE kfz (
	id serial PRIMARY KEY,
	hersteller_name int REFERENCES kfz_hersteller(id),
	kfz_name text UNIQUE NOT NULL,
	serie_ab_werk boolean NOT NULL,
	tuning int REFERENCES tuning(id),
	ps int NOT NULL,
	ps_gemessen boolean NOT NULL,
	nm int NOT NULL,
	nm_gemessen boolean NOT NULL,
	herstellungsjahr int NOT NULL,
	gewicht int NOT NULL,
	gewicht_gemessen boolean);
	
CREATE TABLE tuning (
	id serial PRIMARY KEY,
	basis_kfz int REFERENCES kfz(id),
	tuning_bezeichnung text,
	ansaugung boolean NOT NULL,
	aubgasanlage boolean NOT NULL,
	ladedruck boolean NOT NULL,
	fahrwerk boolean NOT NULL,
	luftfahrwerk boolean NOT NULL,
	bremsanlage boolean NOT NULL,
	getriebe boolean NOT NULL,
	software boolean NOT NULL,
	gewichts_verringerung boolean NOT NULL,
	reifen boolean NOT NULL,
	karosserie boolean NOT NULL,
	aerodynamik boolean NOT NULL,
	plus_10_ps_pro_sticker boolean,
	sticker_anzahl int);
	
CREATE TABLE zeiten_0_100 (
	id serial PRIMARY KEY,
	kfz int REFERENCES kfz(id),
	zeituhr int REFERENCES zeituhr(id),
	nass int NOT NULL);

CREATE TABLE zeiten_50_150 (
	id serial PRIMARY KEY,
	kfz int REFERENCES kfz(id),
	zeituhr int REFERENCES zeituhr(id),
	nass int NOT NULL);
	
CREATE TABLE zeiten_100_200 (
	id serial PRIMARY KEY,
	kfz int REFERENCES kfz(id),
	zeituhr int REFERENCES zeituhr(id),
	nass int NOT NULL);

CREATE TABLE zeituhr (
	id serial PRIMARY KEY,
	hersteller_name text NOT NULL,
	name text NOT NULL);
	
-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
