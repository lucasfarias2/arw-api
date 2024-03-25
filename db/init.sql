drop table if exists roles cascade;
drop table if exists organizations cascade;
drop table if exists projects cascade;
drop table if exists users cascade;
drop table if exists user_organization cascade;
drop table if exists user_project cascade;
drop table if exists toolkits cascade;

create table users
(
    id       serial not null,
    email    text   not null unique,
    password text   not null,
    primary key (id)
);

insert into users(email, password)
values ('mocked@test.com', 'asdasd123');
