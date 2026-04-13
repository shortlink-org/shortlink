CREATE DOMAIN link AS jsonb CHECK (
    (
        jsonb_typeof(VALUE->'url') = 'string'
        AND length(VALUE->>'url') > 0
    )
    OR (
        jsonb_typeof(VALUE->'uri') = 'string'
        AND length(VALUE->>'uri') > 0
    )
);

ALTER TABLE link.links ALTER COLUMN json TYPE link USING json::link;
