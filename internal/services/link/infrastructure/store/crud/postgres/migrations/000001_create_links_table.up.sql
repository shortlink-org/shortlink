-- ShortLink Schema ====================================================================================================

-- for local development
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
-- CREATE EXTENSION IF NOT EXISTS pg_prewarm;
-- ALTER SYSTEM SET shared_preload_libraries = 'pg_prewarm';
CREATE SCHEMA IF NOT EXISTS shortlink;

COMMENT ON SCHEMA shortlink IS 'Shortlink schema';

-- Create a table for links
CREATE TABLE shortlink.links
(
    id       UUID NOT NULL DEFAULT uuid_generate_v4(),
             CONSTRAINT id_links PRIMARY KEY(id),
    url      varchar(255) not null,
    hash     varchar(20)  not null,
    describe text,
    json     jsonb        not null
) WITH (fillfactor = 100);

COMMENT ON TABLE shortlink.links IS 'Link list';

CREATE UNIQUE INDEX links_id_uindex
    ON shortlink.links (id);

CREATE UNIQUE INDEX links_hash_uindex
    ON shortlink.links (hash);

-- INCLUDE-index
-- as example: SELECT id, url, hash FROM links WHERE id = 10;
CREATE UNIQUE INDEX links_list ON shortlink.links USING btree (hash) INCLUDE (url);
