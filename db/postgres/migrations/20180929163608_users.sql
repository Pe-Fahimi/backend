-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE users
(
    id                   BIGSERIAL   NOT NULL PRIMARY KEY,
    email                TEXT        NOT NULL UNIQUE,
    password_hash        TEXT        NOT NULL,
    registered_at        TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE users CASCADE;
