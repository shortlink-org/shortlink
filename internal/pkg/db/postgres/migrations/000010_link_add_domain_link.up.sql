CREATE DOMAIN link AS jsonb CHECK (
    CASE jsonb_typeof(value->'url')::text
        WHEN 'string' THEN
            length(to_json(value->'url')::text) > 0
        AND
            length(to_json(value->'hash')::text) > 0
        ELSE false
    END
);

AlTER TABLE shortlink.links ALTER column json TYPE link USING json::link;
