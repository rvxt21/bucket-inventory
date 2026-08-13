-- +goose Up
CREATE TABLE IF NOT EXISTS files.file (
    id              TEXT PRIMARY KEY DEFAULT uuidv7(), 
    name            TEXT NOT NULL,
    storage_key     TEXT NOT NULL UNIQUE,
    content_type    TEXT,
    size            BIGINT,
    created_at      TIMESTAMPTZ DEFAULT now(),
    updated_at      TIMESTAMPTZ DEFAULT now()
);

-- +goose Down
DROP TABLE IF EXISTS files.file;
