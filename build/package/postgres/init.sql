-- drop database if exists shop;
-- create database shop;

drop table if exists User cascade;
create table public.User(
    user_id serial not null primary key,
    user_login text not null,
    user_password text not null,
    user_role text not null
);

drop table if exists Item cascade;
create table public.Item(
    id serial not null primary key,
    brand text not null,
    category text not null,
    size text not null,
    price int not null check (price > 0),
    sex text not null,
    image_id int not null
);

drop table if exists Order cascade;
create table public.Order(
    id serial not null primary key,
    commit_date date,
    user_id int not null,
    foreign key (user_id) references public.User(user_id),
    price int not null check (price > 0),
    current_status text not null,
);

drop table if exists OrderItems cascade;
create table public.OrderItems(
    id serial not null primary key,
    order_id int not null,
    foreign key (order_id) references public.Order(id),
    item_id int not null,
    foreign key (item_id) references public.Item(id),
    amount int not null check (amount > 0),
);

drop table if exists Storage cascade;
create table public.Storage(
    id serial not null primary key,
    full_address text not null,
    staff_numbers int not null check (amount > 0),
    floor_space int not null check (amount > 0),
);

drop table if exists StorageItems cascade;
create table public.OrderItems(
    id serial not null primary key,
    storage_id int not null,
    foreign key (storage_id) references public.Storage(id),
    item_id int not null,
    foreign key (item_id) references public.Item(id),
    amount int not null check (amount > 0),
);

set datestyle to 'dmy';

create user "default_guest";
create user "default_user";
create user "default_admin";

alter role "default_guest" password '00000000';
alter role "default_user" password '11111111';
alter role "default_admin" password '22222222';

grant select on table Item to "default_guest";
grant select on table Storage to "default_guest";
grant select on table StorageItems to "default_guest";

grant select on table Item to "default_user";
grant select on table Storage to "default_user";
grant select on table StorageItems to "default_user";
grant select on table Order to "default_user";
grant select on table OrderItems to "default_user";

alter role "default_admin" superuser;
