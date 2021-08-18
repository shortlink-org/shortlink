-- BILLING SCHEMA ======================================================================================================

CREATE SCHEMA shortlink;

COMMENT ON SCHEMA shortlink IS 'Shortlink schema';

-- CQRS for links ======================================================================================================
create table shortlink.link_view
(
	id UUID default uuid_generate_v4() not null,
	url varchar(255) not null,
	hash varchar(20) not null,
	describe text,
	created_at TIMESTAMP default current_timestamp,
	updated_at TIMESTAMP default current_timestamp
) WITH (fillfactor = 100);

comment on table shortlink.link_view is 'CQRS for links';

create index link_view_hash_index
	on shortlink.link_view (hash);
