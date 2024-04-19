ALTER TABLE link.link_view REPLICA IDENTITY FULL;

alter table link.link_view
	add constraint link_view_pk
		primary key (id);
