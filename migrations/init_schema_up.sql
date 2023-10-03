CREATE TABLE IF NOT EXISTS users (
    id bigserial not null primary key,
    hashed_pass varchar(64) not null ,
    hashed_refresh_token varchar(64) not null,
    email varchar(32) unique not null,
    nick varchar(12) unique not null,
    image bytea
);

CREATE TABLE IF NOT EXISTS posts (
    id bigserial not null,
    user_id bigint not null references users(id),
    title varchar(16) not null,
    description varchar(255) not null
);