CREATE MATERIALIZED VIEW billing.billing_account_balance
AS
    SELECT account.user_id, tariff.name
    FROM billing.account as account
    LEFT JOIN billing.tariff as tariff
        ON account.tariff_id = tariff.id
WITH NO DATA;
