-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE getriebe (
	id serial PRIMARY KEY,
	getriebe_bezeichnung text NOT NULL,
	gang_anzahl int);

INSERT INTO getriebe (getriebe_bezeichnung, gang_anzahl) VALUES ('Manuell', 4), ('Manuell', 5), ('Manuell', 6), ('Automatik', 6), ('Doppelkupplungsgetriebe', 6), ('Doppelkupplungsgetriebe', 7), ('Doppelkupplungsgetriebe', 8), ('Doppelkupplungsgetriebe', 9);

ALTER TABLE kfz_variante ADD getriebe int REFERENCES getriebe(id);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
