CREATE 
OR REPLACE FUNCTION ItemsInUsersBasket(webUser int) RETURNS TABLE (
  id int,
  category text,
  size text,
  price int, 
  sex text,
  image_id int,
  brand_id int, 
  is_available boolean,
  amount int
) AS $$ declare basket_id int;
BEGIN 
select 
  o.id into basket_id 
from 
  Ordering o 
where 
  o.user_id = webUser 
  and current_status = 'корзина';
