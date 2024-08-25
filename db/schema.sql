CREATE TABLE IF NOT EXISTS "schema_migrations" (version varchar(128) primary key);
CREATE TABLE babies (
    id integer primary key,
    slug text unique not null,
    name text not null,
    birth_date text,
    created_at text not null default CURRENT_TIMESTAMP,
    timezone text
, client_info_id integer references client_info(id));
CREATE TABLE words (
    id integer primary key,
    baby_id integer not null,
    word text not null,
    number integer not null,
    learned_date text not null,
    created_at text not null default CURRENT_TIMESTAMP, client_info_id integer references client_info(id),
    UNIQUE(baby_id, word),
    FOREIGN KEY (baby_id) REFERENCES babies (id)
);
CREATE TABLE client_info (
    id integer primary key,
    user_agent string not null,
    ip_address string not null,
    created_at text not null default CURRENT_TIMESTAMP,
    UNIQUE(ip_address, user_agent)
);
-- Dbmate schema migrations
INSERT INTO "schema_migrations" (version) VALUES
  ('20240325015319'),
  ('20240825010958');
