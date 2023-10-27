CREATE TABLE links
(
  id         CHAR(36) NOT NULL,
  url        TEXT NOT NULL,
  hash       CHAR(40) NOT NULL,  -- Adjusting for a potential SHA-1 hash
  `describe` TEXT,
  json       JSON NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY(id),
  INDEX hash_index (hash),
  INDEX url_index (url(255))    -- Index for url with prefix length
) ENGINE=InnoDB
  DEFAULT CHARSET=utf8mb4
  COLLATE=utf8mb4_unicode_ci
  COMMENT 'Link list';
