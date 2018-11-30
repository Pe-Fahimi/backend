-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied

ALTER TABLE items ADD COLUMN image_url TEXT NULL;

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back

ALTER TABLE items DROP COLUMN image_url;
