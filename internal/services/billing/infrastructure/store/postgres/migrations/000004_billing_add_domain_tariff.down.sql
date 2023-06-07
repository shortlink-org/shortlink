ALTER TABLE billing.tariff ALTER COLUMN payload TYPE jsonb USING payload::jsonb;

DROP DOMAIN tariff;
