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
