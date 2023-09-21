create table public.Ordering(
  id serial not null primary key, 
  commit_date date, 
  user_id int not null, 
  foreign key (user_id) references public.webUser(user_id), 
  price int check (price > 0), 
  current_status text not null
);
