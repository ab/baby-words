-- migrate:up
create table babies (
    id integer primary key,
    slug text unique not null,
    name text not null,
    birth_date text,
    created_at text not null default CURRENT_TIMESTAMP,
    timezone text
);

create table words (
    id integer primary key,
    baby_id integer not null,
    word text not null,
    number integer not null,
    learned_date text not null,
    created_at text not null default CURRENT_TIMESTAMP,
    UNIQUE(baby_id, word),
    FOREIGN KEY (baby_id) REFERENCES babies (id)
);


-- migrate:down

drop table words;
drop table babies;


