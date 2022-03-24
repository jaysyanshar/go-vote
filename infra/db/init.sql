create database if not exists go_vote;

use go_vote;

create table if not exists users (
    id bigint auto_increment not null primary key,
    name varchar(50) null,
    email varchar(100) not null unique,
    password varchar(255) not null
);

create table if not exists sessions (
    id bigint auto_increment not null primary key,
    userId bigint not null references users(id),
    ipAddress varchar(25) not null,
    createdAt datetime not null default now(),
    expiredAt datetime not null,
    isRevoked bool not null default false
);