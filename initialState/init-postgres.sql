-- Create a table for links
CREATE TABLE links
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

CREATE UNIQUE INDEX links_id_uindex
    ON links (id);

CREATE UNIQUE INDEX links_hash_uindex
    ON links (hash);

-- INCLUDE-index
-- as example: SELECT id, url, hash FROM links WHERE id = 10;
CREATE UNIQUE INDEX links_list ON links USING btree (hash) INCLUDE (url);

-- TRANSACTION ---------------------------------------------------------------------------------------------------------
BEGIN
    ISOLATION LEVEL READ COMMITTED;

SAVEPOINT tx_create_default_links;

INSERT INTO links(url, hash, describe, json)
    VALUES ('https://batazor.ru', 'myHash1', 'My personal website', '{"url":"https://batazor.ru", "hash":"myHash1","describe":"My personal website"}');

INSERT INTO links(url, hash, describe, json)
    VALUES ('https://github.com/batazor', 'myHash2', 'My accout of github', '{"url":"https://github.com/batazor", "hash":"myHash2","describe":"My accout of github"}');

INSERT INTO links(url, hash, describe, json)
    VALUES ('https://vk.com/batazor', 'myHash3', 'My page on vk.com', '{"url":"https://vk.com/batazor", "hash":"myHash3","describe":"My page on vk.com"}');

-- ROLLBACK TO tx_create_default_links;

COMMIT;
