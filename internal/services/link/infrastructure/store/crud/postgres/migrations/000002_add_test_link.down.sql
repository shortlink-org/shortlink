BEGIN
    ISOLATION LEVEL READ COMMITTED;

DELETE FROM shortlink.links
  WHERE hash IN ("myHash1", "myHash2", "myHash3")

COMMIT;
