create table users
(
    id       integer not null
        constraint users_id_key
            unique,
    prev_msg varchar,
    role     integer default 0,
    username varchar
);

alter table users
    owner to postgres;

create table ips
(
    id            serial
        constraint ips_id_key
            unique,
    query         varchar(255) not null
        constraint ips_pk
            primary key,
    status        varchar(255) not null,
    country       varchar(255),
    "countryCode" varchar(255),
    region        varchar(255),
    "regionName"  varchar(255),
    city          varchar(255),
    zip           varchar(255),
    lat           double precision,
    lon           double precision,
    timezone      varchar(255),
    isp           varchar(255),
    org           varchar(255),
    "as"          varchar(255)
);

alter table ips
    owner to postgres;

create unique index ips_query_uindex
    on ips (query);

create table user_histories
(
    ips_id  integer not null,
    user_id integer not null,
    time    timestamp default now(),
    query   varchar(255)
);

alter table user_histories
    owner to postgres;


