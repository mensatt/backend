-- migrate:up
CREATE TABLE users (
    id uuid DEFAULT uuid_generate_v4(),
    email varchar UNIQUE NOT NULL,
    password_hash varchar NOT NULL,
    created_at timestamptz DEFAULT NOW() NOT NULL,
    updated_at timestamptz DEFAULT NOW() NOT NULL,
    PRIMARY KEY (id)
);