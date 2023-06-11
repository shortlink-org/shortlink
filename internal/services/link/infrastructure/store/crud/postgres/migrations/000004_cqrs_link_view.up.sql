-- CQRS for links ======================================================================================================
CREATE TABLE shortlink.link_view
(
  id UUID DEFAULT uuid_generate_v4() NOT NULL,
  url VARCHAR(2048) NOT NULL,
  hash VARCHAR(20) NOT NULL,
  describe TEXT,
  created_at TIMESTAMP DEFAULT current_timestamp,
  updated_at TIMESTAMP DEFAULT current_timestamp
) WITH (FILLFACTOR = 80);

COMMENT ON TABLE shortlink.link_view IS 'CQRS for links';

-- Creating an index concurrently to avoid locking the table
-- We can't use concurrent index because golang-migrate doesn't support it
CREATE INDEX link_view_hash_index
  ON shortlink.link_view (hash);
