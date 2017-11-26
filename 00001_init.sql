-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE kfz_hersteller (
	id serial PRIMARY KEY,
	hersteller_name text UNIQUE);

CREATE TABLE basis_kfz (
	id serial PRIMARY KEY,
	hersteller int REFERENCES kfz_hersteller(id),
	kfz_name text UNIQUE NOT NULL,
	herstellungsjahr int NOT NULL);

CREATE TABLE kfz_variante (
	id serial PRIMARY KEY,
	serie_ab_werk boolean DEFAULT true,
	serien_kfz int REFERENCES basis_kfz(id),
	ps int DEFAULT 0,
	ps_gemessen boolean DEFAULT false,
	nm int DEFAULT 0,
	nm_gemessen boolean DEFAULT false,
	gewicht int DEFAULT 0,
	gewicht_gemessen boolean DEFAULT false);

CREATE TABLE tuning (
	id serial PRIMARY KEY,
	serien_kfz int REFERENCES basis_kfz(id),
	tuning_name text NOT NULL,
	tuning_beschreibung text,
	ansaugung boolean DEFAULT false,
	aubgasanlage boolean DEFAULT false,
	ladedruck boolean DEFAULT false,
	fahrwerk boolean DEFAULT false,
	luftfahrwerk boolean DEFAULT false,
	bremsanlage boolean DEFAULT false,
	getriebe boolean DEFAULT false,
	software boolean DEFAULT false,
	gewichts_verringerung boolean DEFAULT false,
	reifen boolean DEFAULT false,
	karosserie boolean DEFAULT false,
	aerodynamik boolean DEFAULT false,
	plus_10_ps_pro_sticker boolean DEFAULT false,
	sticker_anzahl int DEFAULT 0);
	
ALTER TABLE kfz_variante ADD tuning int REFERENCES tuning(id) DEFAULT null;

CREATE TABLE zeituhr (
	id serial PRIMARY KEY,
	hersteller_name text NOT NULL,
	name text NOT NULL);

CREATE TABLE zeiten_0_100 (
	id serial PRIMARY KEY,
	kfz_variante int REFERENCES kfz_variante(id),
	zeituhr int REFERENCES zeituhr(id),
	nass int NOT NULL,
	gemessene_zeit decimal NOT NULL,
	youtube_url text);

CREATE TABLE zeiten_50_150 (
	id serial PRIMARY KEY,
	kfz_variante int REFERENCES kfz_variante(id),
	zeituhr int REFERENCES zeituhr(id),
	nass int NOT NULL,
	gemessene_zeit decimal NOT NULL,
	youtube_url text);
	
CREATE TABLE zeiten_100_200 (
	id serial PRIMARY KEY,
	kfz_variante int REFERENCES kfz_variante(id),
	zeituhr int REFERENCES zeituhr(id),
	nass int NOT NULL,
	gemessene_zeit decimal NOT NULL,
	youtube_url text);
	
INSERT INTO zeituhr (hersteller_name, name) VALUES ('Qstarz', 'LT-6000S');
INSERT INTO kfz_hersteller (hersteller_name) VALUES ('Porsche'), ('Mercedes-Benz'), ('BMW'), ('Audi'), ('VW'), ('Lexus'), ('Mini'), ('Ford'), ('Nissan'), ('Ferrari'), ('Lamborghini'), ('Toyota');
INSERT INTO basis_kfz (hersteller, kfz_name, herstellungsjahr) VALUES (5, 'Golf 7 R', 2016);
INSERT INTO tuning (serien_kfz, tuning_name, software) VALUES (1, 'Stage 1', true);
INSERT INTO kfz_variante (serie_ab_werk, ps, ps_gemessen, nm, nm_gemessen, tuning) VALUES (true, 300, true, 400, true, null), (false, 350, true, 480, true, 1);
INSERT INTO zeiten_100_200 (kfz_variante, zeituhr, nass, gemessene_zeit, youtube_url) VALUES (1, 1, 0, 14.347, 'https://www.youtube.com/watch?v=La9jsFkVs9o'), (2, 1, 0, 12.490, 'https://www.youtube.com/watch?v=La9jsFkVs9o');

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
