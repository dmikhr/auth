-- +goose Up
-- +goose StatementBegin
CREATE TYPE role AS ENUM ('UNKNOWN', 'USER', 'ADMIN');

create table users (
    id serial primary key,
    name text not null,
    email text not null,
    password text not null,
    role role,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
