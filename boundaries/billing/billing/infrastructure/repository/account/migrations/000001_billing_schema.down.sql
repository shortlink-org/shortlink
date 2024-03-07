-- DROP FOREIGN KEY CONSTRAINTS ========================================================================================
ALTER TABLE billing.account DROP CONSTRAINT IF EXISTS "account_tariff_id_foreign";

-- DROP TABLE =========================================================================================================
DROP TABLE IF EXISTS billing.account;
