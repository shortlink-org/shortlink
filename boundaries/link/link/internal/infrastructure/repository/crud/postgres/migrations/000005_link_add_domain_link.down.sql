ALTER TABLE link.links ALTER COLUMN json TYPE jsonb USING json::jsonb;

DROP DOMAIN link;
