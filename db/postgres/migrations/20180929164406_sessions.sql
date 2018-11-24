-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE sessions
(
    id         BIGSERIAL   NOT NULL PRIMARY KEY,
    user_id    BIGINT      NOT NULL REFERENCES users (id) ON DELETE CASCADE ON UPDATE CASCADE,
    token      TEXT        NOT NULL UNIQUE,
    client_ip  TEXT,
    user_agent TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    expires_at TIMESTAMPTZ NOT NULL,
    deleted_at TIMESTAMPTZ
);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE sessions CASCADE;
