-- +goose Up
-- SQL in this section is executed when the migration is applied.
ALTER TABLE basis_kfz RENAME COLUMN "herstellungsjahr" TO "fabrikationsjahr";

ALTER TABLE tuning RENAME COLUMN "aubgasanlage" TO "abgasanlage";

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
