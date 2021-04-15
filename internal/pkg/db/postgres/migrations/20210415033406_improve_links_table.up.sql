BEGIN
    ISOLATION LEVEL READ COMMITTED;

ALTER TABLE links
	ADD created_at TIMESTAMP DEFAULT current_timestamp;

ALTER TABLE links
	ADD updated_at TIMESTAMP DEFAULT current_timestamp;

COMMIT;
