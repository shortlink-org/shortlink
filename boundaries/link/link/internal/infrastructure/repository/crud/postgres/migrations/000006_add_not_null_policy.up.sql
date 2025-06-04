alter table link.links alter column created_at set not null;
alter table link.links alter column updated_at set not null;

alter table link.link_view alter column created_at set not null;
alter table link.link_view alter column updated_at set not null;
