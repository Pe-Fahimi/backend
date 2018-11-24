-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE locations
(
    id        BIGSERIAL NOT NULL PRIMARY KEY,
    title     TEXT      NOT NULL,
    parent_id BIGINT REFERENCES locations (id) ON DELETE SET NULL ON UPDATE CASCADE
);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE locations CASCADE;
