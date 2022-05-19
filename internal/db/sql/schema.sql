CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE dish (
    id uuid DEFAULT uuid_generate_v4(),
    name varchar NOT NULL,
    PRIMARY KEY (id)
);

CREATE TYPE priority AS ENUM ('UNSET', 'LOW', 'MEDIUM', 'HIGH');

CREATE TABLE tag (
    key varchar,
    name varchar NOT NULL,
    description varchar NOT NULL,
    short_name varchar,
    priority priority DEFAULT 'UNSET' NOT NULL,
    is_allergy boolean DEFAULT FALSE NOT NULL,
    PRIMARY KEY (key)
);

CREATE TABLE occurrence (
    id uuid DEFAULT uuid_generate_v4(),
    dish uuid NOT NULL,
    date date NOT NULL,
    kj float,
    kcal float,
    fat float,
    saturated_fat float,
    carbohydrates float,
    sugar float,
    fiber float,
    protein float,
    salt float,
    price_student integer NOT NULL,
    price_staff integer NOT NULL,
    price_guest integer NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE occurrence_side_dishes (
    occurrence_id uuid NOT NULL,
    dish_id uuid NOT NULL,
    CONSTRAINT fk_occurrence FOREIGN KEY(occurrence_id) REFERENCES occurrence(id),
    CONSTRAINT fk_dish FOREIGN KEY(dish_id) REFERENCES dish(id)
);

CREATE TABLE occurrence_tag (
    occurrence_id uuid NOT NULL,
    tag_key varchar NOT NULL,
    CONSTRAINT fk_occurrence FOREIGN KEY(occurrence_id) REFERENCES occurrence(id),
    CONSTRAINT fk_tag FOREIGN KEY(tag_key) REFERENCES tag(key)
);

CREATE TABLE review (
    id uuid DEFAULT uuid_generate_v4(),
    occurrence uuid NOT NULL,
    display_name varchar(32),
    stars integer NOT NULL,
    text text,
    up_votes integer DEFAULT 0 NOT NULL,
    down_votes integer DEFAULT 0 NOT NULL,
    created_at timestamptz DEFAULT NOW() NOT NULL,
    updated_at timestamptz DEFAULT NOW() NOT NULL,
    accepted_at timestamptz,
    PRIMARY KEY (id),
    CONSTRAINT fk_occurrence FOREIGN KEY(occurrence) REFERENCES occurrence(id)
);

CREATE TABLE image (
    id uuid DEFAULT uuid_generate_v4(),
    occurrence uuid NOT NULL,
    display_name varchar(32) NOT NULL,
    description text,
    up_votes integer DEFAULT 0 NOT NULL,
    down_votes integer DEFAULT 0 NOT NULL,
    created_at timestamptz DEFAULT NOW() NOT NULL,
    updated_at timestamptz DEFAULT NOW() NOT NULL,
    accepted_at timestamptz,
    PRIMARY KEY (id),
    CONSTRAINT fk_occurrence FOREIGN KEY(occurrence) REFERENCES occurrence(id)
);