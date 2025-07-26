-- +goose Up
-- +goose StatementBegin
CREATE TABLE categories (
    id BIGINT PRIMARY KEY DEFAULT unique_rowid (),
    name STRING NOT NULL,
    slug STRING NOT NULL,
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now(),
    UNIQUE (slug)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE categories;
-- +goose StatementEnd