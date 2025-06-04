-- Revert the domain constraint back to support up to 9 characters
ALTER DOMAIN hash DROP CONSTRAINT hash_check;
ALTER DOMAIN hash ADD CONSTRAINT hash_check CHECK (length(VALUE) > 0 AND length(VALUE) <= 9);
