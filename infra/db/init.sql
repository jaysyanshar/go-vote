create database if not exists go_vote;

use go_vote;

create table if not exists users (
    id bigint auto_increment not null primary key,
    name varchar(50) null,
    email varchar(100) not null unique,
    password varchar(255) not null
);