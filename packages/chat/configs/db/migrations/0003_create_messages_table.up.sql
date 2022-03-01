CREATE TABLE IF NOT EXISTS messages (
    id bigserial NOT NULL,
    conversation_id bigint NOT NULL,
    content character varying(255) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    sender_id bigint NOT NULL,
    state varchar(255) NOT NULL,
    CONSTRAINT messages_pk PRIMARY KEY (id),
    CONSTRAINT messages_sender_id_fk FOREIGN KEY (sender_id)
        REFERENCES users (id) MATCH SIMPLE
        ON UPDATE NO ACTION ON DELETE NO ACTION,
    CONSTRAINT messages_conversation_id_fk FOREIGN KEY (conversation_id)
        REFERENCES conversations (id) MATCH SIMPLE
        ON UPDATE NO ACTION ON DELETE NO ACTION
);