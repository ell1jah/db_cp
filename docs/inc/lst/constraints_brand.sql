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
