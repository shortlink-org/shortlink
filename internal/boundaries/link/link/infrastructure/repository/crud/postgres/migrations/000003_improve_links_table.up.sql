BEGIN
    ISOLATION LEVEL READ COMMITTED;

ALTER TABLE link.links
	ADD created_at TIMESTAMP DEFAULT current_timestamp;

ALTER TABLE link.links
	ADD updated_at TIMESTAMP DEFAULT current_timestamp;

COMMIT;
