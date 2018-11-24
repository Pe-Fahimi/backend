-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE items
(
    id          BIGSERIAL   NOT NULL PRIMARY KEY,
    title       TEXT        NOT NULL,
    content     TEXT        NOT NULL,
    author_id   BIGINT      NOT NULL REFERENCES users (id) ON DELETE CASCADE ON UPDATE CASCADE,
    location_id BIGINT      NOT NULL REFERENCES locations (id) ON DELETE CASCADE ON UPDATE CASCADE,
    category_id BIGINT      NOT NULL REFERENCES categories (id) ON DELETE CASCADE ON UPDATE CASCADE,
    status      TEXT        NOT NULL DEFAULT 'pending',
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at  TIMESTAMPTZ
);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE items CASCADE;
