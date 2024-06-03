-- DOMAIN FOR PAYLOAD ==================================================================================================
CREATE DOMAIN price AS jsonb CHECK (
    CASE jsonb_typeof(value->'amount')
        -- only cast when it is safe to do so
        -- note: this will reject numeric values stored as text in the json object (eg '{"amount": "1"}')
        WHEN 'number' THEN (value->>'amount')::integer >= 0
        ELSE false
    END
);

-- VALIDATE CURRENT DATA AND ALTER COLUMN TYPE ========================================================================
-- Ensure all existing data adheres to the new domain constraints before altering the column type
UPDATE tariff
  SET payload = payload
  WHERE (payload->>'amount')::integer >= 0;

ALTER TABLE tariff
  ALTER COLUMN payload
  SET DATA TYPE price
  USING payload::jsonb;
