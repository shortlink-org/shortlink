-- Details of this decision
-- https://www.depesz.com/2010/03/02/charx-vs-varcharx-vs-varchar-vs-text/
-- https://www.postgresql.org/docs/9.1/datatype-character.html

alter table link.link_view alter column url type text using url::text;
alter table link.links alter column url type text using url::text;

CREATE DOMAIN hash AS TEXT CHECK (length(VALUE) > 0 AND length(VALUE) <= 9);

alter table link.link_view alter column hash type hash using hash::hash;
alter table link.links alter column hash type hash using hash::hash;

