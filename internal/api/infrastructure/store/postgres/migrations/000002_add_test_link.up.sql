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
