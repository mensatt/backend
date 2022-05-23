CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE dish (
    id uuid DEFAULT uuid_generate_v4(),
    name varchar UNIQUE NOT NULL,
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

CREATE TYPE review_status AS ENUM('APPROVED', 'AWAITING_APPROVAL', 'UPDATED');

CREATE TABLE occurrence (
    id uuid DEFAULT uuid_generate_v4(),
    dish uuid NOT NULL,
    date date NOT NULL,
    review_status review_status DEFAULT 'AWAITING_APPROVAL' NOT NULL,
    kj integer,
    kcal integer,
    fat integer,
    saturated_fat integer,
    carbohydrates integer,
    sugar integer,
    fiber integer,
    protein integer,
    salt integer,
    price_student integer NOT NULL,
    price_staff integer NOT NULL,
    price_guest integer NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY(dish) REFERENCES dish(id)
);

CREATE TABLE occurrence_side_dishes (
    occurrence uuid NOT NULL,
    dish uuid NOT NULL,
    FOREIGN KEY(occurrence) REFERENCES occurrence(id),
    FOREIGN KEY(dish) REFERENCES dish(id)
);

CREATE TABLE occurrence_tag (
    occurrence uuid NOT NULL,
    tag varchar NOT NULL,
    FOREIGN KEY(occurrence) REFERENCES occurrence(id),
    FOREIGN KEY(tag) REFERENCES tag(key)
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
    FOREIGN KEY(occurrence) REFERENCES occurrence(id)
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
    FOREIGN KEY(occurrence) REFERENCES occurrence(id)
);

CREATE TABLE users (
    id uuid DEFAULT uuid_generate_v4(),
    email varchar UNIQUE NOT NULL,
    password_hash varchar NOT NULL,
    created_at timestamptz DEFAULT NOW() NOT NULL,
    updated_at timestamptz DEFAULT NOW() NOT NULL,
    PRIMARY KEY (id)
);