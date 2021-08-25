CREATE EXTENSION pg_prewarm;
ALTER SYSTEM SET shared_preload_libraries = 'pg_prewarm';
