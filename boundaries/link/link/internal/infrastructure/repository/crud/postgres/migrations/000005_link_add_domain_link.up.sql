CREATE DOMAIN link AS jsonb CHECK (
    jsonb_typeof(value->'uri') = 'string' AND
    length(value->>'uri') > 0
);

AlTER TABLE link.links ALTER column json TYPE link USING json::link;
