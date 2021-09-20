ALTER TABLE shortlink.link_view REPLICA IDENTITY FULL;

alter table shortlink.link_view
	add constraint link_view_pk
		primary key (id);
