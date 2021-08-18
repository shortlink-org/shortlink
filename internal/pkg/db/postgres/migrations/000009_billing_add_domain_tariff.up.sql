CREATE DOMAIN tariff AS jsonb CHECK (
    CASE jsonb_typeof(value->'amount')
        -- only cast when it is safe to do so
        -- note: this will reject numeric values stored as text in the json object (eg '{"amount": "1"}')
        WHEN 'number' THEN (value->>'amount')::integer >= 0
            AND value->>'amount' IS NOT NULL
        ELSE false
    END
);

AlTER TABLE billing.tariff ALTER column payload TYPE tariff USING payload::tariff;
