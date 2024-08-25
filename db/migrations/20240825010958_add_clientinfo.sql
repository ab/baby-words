-- migrate:up
create table client_info (
    id integer primary key,
    user_agent string not null,
    ip_address string not null,
    created_at text not null default CURRENT_TIMESTAMP,
    UNIQUE(ip_address, user_agent)
);

alter table babies add column client_info_id integer references client_info(id);
alter table words add column client_info_id integer references client_info(id);

-- migrate:down
alter table babies remove column client_info_id;
alter table words remove column client_info_id;

drop table client_info;
