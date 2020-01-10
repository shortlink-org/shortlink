CREATE TABLE links
(
    id       serial       not null
             constraint links_pk
             primary key,
    url      varchar(255) not null,
    hash     varchar(255) not null,
    describe text
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

INSERT INTO links(url, hash, describe)
    VALUES ('https://batazor.ru', 'myHash1', 'My personal website');

INSERT INTO links(url, hash, describe)
VALUES ('https://github.com/batazor', 'myHash2', 'My accout of github');

INSERT INTO links(url, hash, describe)
VALUES ('https://vk.com/batazor', 'myHash3', 'My page on vk.com');

-- ROLLBACK TO tx_create_default_links;

COMMIT;
