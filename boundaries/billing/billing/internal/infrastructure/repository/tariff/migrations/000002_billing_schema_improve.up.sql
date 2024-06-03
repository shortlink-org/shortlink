alter table tariff alter column created_at set default now();
alter table tariff alter column updated_at set default now();
alter table tariff alter column id set default gen_random_uuid();

alter table snapshots alter column created_at set default now();
alter table snapshots alter column updated_at set default now();

alter table aggregates alter column created_at set default now();
alter table aggregates alter column updated_at set default now();

alter table events alter column id set default gen_random_uuid();
alter table events alter column created_at set default now();
