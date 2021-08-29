CREATE OR REPLACE FUNCTION make_tsvector_link_view(description TEXT, keywords TEXT)
   RETURNS tsvector AS $$
BEGIN
  RETURN (setweight(to_tsvector('simple', keywords),'A') ||
    setweight(to_tsvector('simple', description), 'B'));
END
$$ LANGUAGE 'plpgsql' IMMUTABLE;

CREATE INDEX idx_fts_link_view ON shortlink.link_view
  USING gin(make_tsvector_link_view(meta_description, meta_keywords));
