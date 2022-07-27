CREATE TABLE IF NOT EXISTS authors
(
    id          uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    name        varchar(255)     NOT NULL,
    name_search tsvector GENERATED ALWAYS AS ( to_tsvector('indonesian', name)) STORED,
    created_at  timestamp        NOT NULL DEFAULT (NOW())
);

INSERT INTO authors(id, name)
VALUES ('927ef3a7-08be-41b8-a156-e75f5cfe199e', 'Arya Yunanta'),
       ('3604e5c7-6a63-4718-899b-8ff0ae96ac5f', 'Achmad Mario Sunjabar'),
       ('d3ad2722-675e-4d94-935a-a31fb5495dbf', 'Artamananda');