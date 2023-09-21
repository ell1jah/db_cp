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
