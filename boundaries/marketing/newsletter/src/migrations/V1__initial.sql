create table shortlink.newsletters
(
    id    UUID default gen_random_uuid() not null
        constraint newsletters_pk
            primary key,
    email text not null
);

create unique index newsletters_email_uindex
    on shortlink.newsletters (email);

create unique index newsletters_id_uindex
    on shortlink.newsletters (id);
