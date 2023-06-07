alter table shortlink.links alter column created_at set not null;
alter table shortlink.links alter column updated_at set not null;

alter table shortlink.link_view alter column created_at set not null;
alter table shortlink.link_view alter column updated_at set not null;
