CREATE TABLE IF NOT EXISTS links (
    id          binary(16)   NOT NULL DEFAULT (UUID_TO_BIN(UUID(), TRUE)),
    url         varchar(255) NOT NULL,
    hash        varchar(255) NOT NULL,
    description text         NULL,
    created_at  timestamp    NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  timestamp    NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
