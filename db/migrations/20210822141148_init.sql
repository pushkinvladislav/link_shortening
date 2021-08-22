-- +goose Up
-- +goose StatementBegin
CREATE TABLE URLs (
    id serial not null unique,
    longURL varchar not null unique,
    shortURL varchar(10) not null unique
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE URLs
-- +goose StatementEnd
