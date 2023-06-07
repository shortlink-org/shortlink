BEGIN
    ISOLATION LEVEL READ COMMITTED;

ALTER TABLE shortlink.links DROP COLUMN created_at;
ALTER TABLE shortlink.links DROP COLUMN updated_at;

COMMIT;
