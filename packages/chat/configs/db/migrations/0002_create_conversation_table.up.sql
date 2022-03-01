CREATE TABLE IF NOT EXISTS conversations (
    id bigserial NOT NULL,
    name character varying(255) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    owner_id bigint NOT NULL,
    CONSTRAINT conversations_pk PRIMARY KEY (id),
    CONSTRAINT conversations_owner_id_fk FOREIGN KEY (owner_id)
        REFERENCES users(id) MATCH SIMPLE
        ON UPDATE NO ACTION ON DELETE NO ACTION
);