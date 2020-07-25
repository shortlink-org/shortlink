-- Create a table for links
CREATE TABLE IF NOT EXISTS links
(
    id       serial       not null
             constraint links_pk
             primary key,
    url      varchar(255) not null,
    hash     varchar(255) not null,
    describe text,
    json     jsonb        not null
);

COMMENT ON TABLE links IS 'Link list';

ALTER TABLE links
    OWNER TO shortlink;

CREATE UNIQUE INDEX IF NOT EXISTS links_id_uindex
    ON links (id);

CREATE UNIQUE INDEX IF NOT EXISTS links_hash_uindex
    ON links (hash);

-- INCLUDE-index
-- as example: SELECT id, url, hash FROM links WHERE id = 10;
CREATE UNIQUE INDEX IF NOT EXISTS links_list ON links USING btree (hash) INCLUDE (url);
