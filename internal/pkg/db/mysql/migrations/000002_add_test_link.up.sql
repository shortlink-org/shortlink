-- TRANSACTION ---------------------------------------------------------------------------------------------------------
BEGIN;

INSERT INTO links(url, hash, description)
    VALUES ('https://batazor.ru', 'myHash1', 'My personal website');

INSERT INTO links(url, hash, description)
    VALUES ('https://github.com/batazor', 'myHash2', 'My accout of github');

INSERT INTO links(url, hash, description)
    VALUES ('https://vk.com/batazor', 'myHash3', 'My page on vk.com');

-- ROLLBACK TO tx_create_default_links;

COMMIT;
