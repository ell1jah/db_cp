-- drop database if exists clothshop;
-- create database clothshop;
create table public.webUser(
  user_id serial not null primary key, 
  user_login text not null unique, user_password text not null, 
  user_name text not null, user_sex text not null, 
  user_role text not null
);
create table public.Brand(
  id serial not null primary key, 
  brand_name text not null, 
  founding_year int not null check (
    founding_year > 1500 
    and founding_year < 2024
  ), 
  logo_id int not null, 
  brand_owner text not null
);
create table public.Item(
  id serial not null primary key, 
  category text not null, 
  size text not null, 
  price int not null check (price > 0), 
  sex text not null, 
  image_id int not null, 
  brand_id int not null, 
  foreign key (brand_id) references public.Brand(id), 
  is_available boolean
);
create table public.Ordering(
  id serial not null primary key, 
  commit_date date, 
  user_id int not null, 
  foreign key (user_id) references public.webUser(user_id), 
  price int check (price > 0), 
  current_status text not null
);
create table public.OrderItems(
  id serial not null primary key, 
  order_id int not null, 
  foreign key (order_id) references public.Ordering(id), 
  item_id int not null, 
  foreign key (item_id) references public.Item(id), 
  amount int not null check (amount > 0)
);
set 
  datestyle to 'dmy';
create user "default_guest";
create user "default_user";
create user "default_admin";
alter role "default_guest" password '00000000';
alter role "default_user" password '11111111';
alter role "default_admin" password '22222222';
grant 
select 
  on table Item to "default_guest";
grant 
select 
  on table Brand to "default_guest";
grant 
select 
  on table Item to "default_user";
grant 
select 
  on table Ordering to "default_user";
grant 
select 
  on table Brand to "default_user";
grant 
select 
  on table OrderItems to "default_user";
alter role "default_admin" superuser;
CREATE 
OR REPLACE FUNCTION AddItemUsersBasket(
  addItem int, webUser int, addAmount int
) RETURNS boolean AS $$ declare basket_id int;
declare orderItem_id int;
BEGIN IF (
  select 
    is_available 
  from 
    Item 
  where 
    id = addItem
) != true THEN return false;
END IF;
select 
  id into basket_id 
from 
  Ordering 
where 
  user_id = webUser 
  and current_status = 'корзина';
select 
  id into orderItem_id 
from 
  OrderItems o 
where 
  o.order_id = basket_id 
  and o.item_id = addItem;
IF FOUND THEN 
UPDATE 
  OrderItems 
SET 
  amount = amount + addAmount 
WHERE 
  id = orderItem_id;
ELSE INSERT INTO OrderItems 
VALUES 
  (
    (
      select 
        max(id) 
      from 
        OrderItems
    ) + 1, 
    basket_id, 
    addItem, 
    addAmount
  );
END IF;
return true;
END $$ LANGUAGE plpgsql;
CREATE 
OR REPLACE FUNCTION DecItemUsersBasket(
  decItem int, webUser int, decAmount int
) RETURNS boolean AS $$ declare basket_id int;
declare orderItem_id int;
declare orderItemAmount int;
BEGIN 
select 
  id into basket_id 
from 
  Ordering 
where 
  user_id = webUser 
  and current_status = 'корзина';
select 
  id, 
  amount into orderItem_id, 
  orderItemAmount 
from 
  OrderItems o 
where 
  o.order_id = basket_id 
  and o.item_id = decItem;
IF orderItemAmount < decAmount THEN return false;
ELSIF orderItemAmount > decAmount THEN 
UPDATE 
  OrderItems 
SET 
  amount = amount - decAmount 
WHERE 
  id = orderItem_id;
ELSE 
DELETE FROM 
  OrderItems 
WHERE 
  id = orderItem_id;
END IF;
return true;
END $$ LANGUAGE plpgsql;
CREATE 
OR REPLACE FUNCTION ItemsInUsersBasket(webUser int) RETURNS TABLE (
  id int, category text, size text, price int, 
  sex text, image_id int, brand_id int, 
  is_available boolean, amount int
) AS $$ declare basket_id int;
BEGIN 
select 
  o.id into basket_id 
from 
  Ordering o 
where 
  o.user_id = webUser 
  and current_status = 'корзина';
return query 
SELECT 
  i.id, 
  i.category, 
  i.size, 
  i.price, 
  i.sex, 
  i.image_id, 
  i.brand_id, 
  i.is_available, 
  o.amount 
FROM 
  OrderItems o 
  JOIN Item i ON o.item_id = i.id 
where 
  o.order_id = basket_id;
END $$ LANGUAGE plpgsql;
CREATE 
OR REPLACE FUNCTION UserBasketPrice(webUser int) RETURNS int
AS $$ declare res int;
BEGIN

  select 
    sum(price * amount) into res
  from 
    ItemsInUsersBasket(webUser);

IF res is null THEN return 0;
else return res;
END IF;

return 0;
END $$ LANGUAGE plpgsql;
CREATE 
OR REPLACE FUNCTION CommitOrder(webUser int) RETURNS int AS $$ declare basket_id int;
BEGIN IF NOT exists (
  select 
    * 
  from 
    ItemsInUsersBasket(webUser)
) THEN return (
  select 
    1
);
END IF;
IF exists (
  select 
    * 
  from 
    ItemsInUsersBasket(webUser) 
  where 
    is_available = false
) THEN return (
  select 
    2
);
END IF;
select 
  o.id into basket_id 
from 
  Ordering o 
where 
  o.user_id = webUser 
  and current_status = 'корзина';
UPDATE 
  Ordering 
SET 
  commit_date = (
    SELECT 
      CURRENT_DATE
  ), 
  price = (
    select 
      UserBasketPrice(webUser)
  ), 
  current_status = 'оформлен' 
WHERE 
  id = basket_id;
INSERT INTO Ordering (id, user_id, current_status) 
VALUES 
  (
    (
      select 
        max(id) 
      from 
        Ordering
    ) + 1, 
    webUser, 
    'корзина'
  );
return (
  select 
    0
);
END $$ LANGUAGE plpgsql;


CREATE OR REPLACE FUNCTION NewUser(
  user_login text, user_password text, user_name text, user_sex text, user_role text
) RETURNS int AS $$ declare cr_id int;
BEGIN 
insert into webUser (user_id, user_login, user_password, user_name, user_sex, user_role)
			values ((select max(user_id) from webUser) + 1, user_login, user_password, user_name, user_sex, user_role)
			returning user_id into cr_id;
insert into ordering (id, user_id, current_status) values ((select max(id) from ordering) + 1, cr_id, 'корзина');
return cr_id;
END $$ LANGUAGE plpgsql;
