CREATE TABLE IF NOT EXISTS messages (
    id bigserial NOT NULL,
    content character varying(255) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    sender_id bigint NOT NULL,
    CONSTRAINT messages_pk PRIMARY KEY (id)
)