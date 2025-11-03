CREATE TYPE role AS ENUM (
    'administrator',
    'teacher',
    'secretary'
);
CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    username text NOT NULL,
    name text NOT NULL,
    surname text NOT NULL,
    password_hash bytea,
    role role NOT NULL,
    created_at TIMESTAMP NOT NULL,
    deleted BOOLEAN NOT NULL
);
