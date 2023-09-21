create table public.OrderItems(
  id serial not null primary key, 
  order_id int not null, 
  foreign key (order_id) references public.Ordering(id), 
  item_id int not null, 
  foreign key (item_id) references public.Item(id), 
  amount int not null check (amount > 0)
);
