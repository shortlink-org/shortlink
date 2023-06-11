BEGIN
    ISOLATION LEVEL READ COMMITTED;

ALTER TABLE shortlink.links
	ADD created_at TIMESTAMP DEFAULT current_timestamp;

ALTER TABLE shortlink.links
	ADD updated_at TIMESTAMP DEFAULT current_timestamp;

COMMIT;
