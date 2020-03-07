drop database if exists demo;
create database demo charset='utf8';

use demo;

drop table if exists todo;
create table todo (

    primary key(id),
    id int not null auto_increment,
    title varchar(256) not null default '待办事项',
    completed bool not null default 0,

    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp,
    deleted_at timestamp not null default current_timestamp
) Engine=Innodb;
