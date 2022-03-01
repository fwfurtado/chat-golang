CREATE TABLE IF NOT EXISTS users (
    id bigserial NOT NULL,
    name character varying(255) NOT NULL,
    CONSTRAINT users_pk PRIMARY KEY (id)
);