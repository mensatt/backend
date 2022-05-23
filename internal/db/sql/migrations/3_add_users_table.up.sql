-- migrate:up
CREATE TABLE users (
    id uuid DEFAULT uuid_generate_v4(),
    email varchar UNIQUE NOT NULL,
    password_hash varchar NOT NULL,
    created_at timestamptz DEFAULT NOW() NOT NULL,
    updated_at timestamptz DEFAULT NOW() NOT NULL,
    PRIMARY KEY (id)
);

INSERT INTO users (email, password_hash) VALUES ('admin@mensatt.de', '$2a$10$pdvY6v8k2McSYbFk3HRDl.h8QfMjOxfpm2CywkDDzfOzlYDZV8NUm');
