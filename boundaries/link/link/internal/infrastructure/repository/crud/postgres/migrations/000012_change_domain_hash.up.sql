ALTER DOMAIN hash DROP CONSTRAINT hash_check;
ALTER DOMAIN hash ADD CONSTRAINT hash_check CHECK (length(VALUE) > 0 AND length(VALUE) <= 15);
