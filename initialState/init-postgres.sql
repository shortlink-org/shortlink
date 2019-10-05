create table links
(
    id       serial       not null
        constraint links_pk
            primary key,
    url      varchar(255) not null,
    hash     varchar(255) not null,
    describe text
);

comment on table links is 'Link list';

alter table links
    owner to shortlink;

create unique index links_id_uindex
    on links (id);

create unique index links_hash_uindex
    on links (hash);

