SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: uuid-ossp; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;


--
-- Name: EXTENSION "uuid-ossp"; Type: COMMENT; Schema: -; Owner: -
--

COMMENT ON EXTENSION "uuid-ossp" IS 'generate universally unique identifiers (UUIDs)';


--
-- Name: priority; Type: TYPE; Schema: public; Owner: -
--

CREATE TYPE public.priority AS ENUM (
    'UNSET',
    'LOW',
    'MEDIUM',
    'HIGH'
);


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: dish; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.dish (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    name character varying NOT NULL
);


--
-- Name: image; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.image (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    occurrence uuid NOT NULL,
    display_name character varying(32) NOT NULL,
    description text,
    up_votes integer DEFAULT 0 NOT NULL,
    down_votes integer DEFAULT 0 NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    accepted_at timestamp with time zone
);


--
-- Name: occurrence; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.occurrence (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    dish uuid NOT NULL,
    date date NOT NULL,
    price_student integer NOT NULL,
    price_staff integer NOT NULL,
    price_guest integer NOT NULL
);


--
-- Name: occurrence_side_dishes; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.occurrence_side_dishes (
    occurrence_id uuid NOT NULL,
    dish_id uuid NOT NULL
);


--
-- Name: occurrence_tag; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.occurrence_tag (
    occurrence_id uuid NOT NULL,
    tag_key character varying NOT NULL
);


--
-- Name: review; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.review (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    occurrence uuid NOT NULL,
    display_name character varying(32),
    stars integer NOT NULL,
    text text,
    up_votes integer DEFAULT 0 NOT NULL,
    down_votes integer DEFAULT 0 NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    accepted_at timestamp with time zone
);


--
-- Name: schema_migrations; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.schema_migrations (
    version character varying(255) NOT NULL
);


--
-- Name: tag; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.tag (
    key character varying NOT NULL,
    name character varying NOT NULL,
    description character varying NOT NULL,
    short_name character varying,
    priority public.priority DEFAULT 'UNSET'::public.priority NOT NULL,
    is_allergy boolean DEFAULT false NOT NULL
);


--
-- Name: dish dish_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.dish
    ADD CONSTRAINT dish_pkey PRIMARY KEY (id);


--
-- Name: image image_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.image
    ADD CONSTRAINT image_pkey PRIMARY KEY (id);


--
-- Name: occurrence occurrence_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.occurrence
    ADD CONSTRAINT occurrence_pkey PRIMARY KEY (id);


--
-- Name: review review_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.review
    ADD CONSTRAINT review_pkey PRIMARY KEY (id);


--
-- Name: schema_migrations schema_migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.schema_migrations
    ADD CONSTRAINT schema_migrations_pkey PRIMARY KEY (version);


--
-- Name: tag tag_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.tag
    ADD CONSTRAINT tag_pkey PRIMARY KEY (key);


--
-- Name: occurrence_side_dishes fk_dish; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.occurrence_side_dishes
    ADD CONSTRAINT fk_dish FOREIGN KEY (dish_id) REFERENCES public.dish(id);


--
-- Name: occurrence_side_dishes fk_occurrence; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.occurrence_side_dishes
    ADD CONSTRAINT fk_occurrence FOREIGN KEY (occurrence_id) REFERENCES public.occurrence(id);


--
-- Name: occurrence_tag fk_occurrence; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.occurrence_tag
    ADD CONSTRAINT fk_occurrence FOREIGN KEY (occurrence_id) REFERENCES public.occurrence(id);


--
-- Name: review fk_occurrence; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.review
    ADD CONSTRAINT fk_occurrence FOREIGN KEY (occurrence) REFERENCES public.occurrence(id);


--
-- Name: image fk_occurrence; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.image
    ADD CONSTRAINT fk_occurrence FOREIGN KEY (occurrence) REFERENCES public.occurrence(id);


--
-- Name: occurrence_tag fk_tag; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.occurrence_tag
    ADD CONSTRAINT fk_tag FOREIGN KEY (tag_key) REFERENCES public.tag(key);


--
-- PostgreSQL database dump complete
--


--
-- Dbmate schema migrations
--

INSERT INTO public.schema_migrations (version) VALUES
    ('20220226235659'),
    ('20220227001217');
