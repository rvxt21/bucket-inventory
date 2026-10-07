-- +goose Up
CREATE SCHEMA IF NOT EXISTS files;

-- +goose Down
DROP SCHEMA IF EXISTS files;
