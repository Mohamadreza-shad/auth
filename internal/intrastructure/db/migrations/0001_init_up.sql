create table if not exists users(
    id BIGINT AUTO_INCREMENT primary key,
    name varchar(50) Not null,
    phone Char(13) Not null,
    national_id char(10) Not null,
    password char(64) not null,
    is_blocked tinyint(1) not null default 0,
);

create index idx_users_phone phone on users;