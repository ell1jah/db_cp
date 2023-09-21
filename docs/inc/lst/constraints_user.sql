create table public.webUser(
  user_id serial not null primary key, 
  user_login text not null unique,
  user_password text not null, 
  user_name text not null,
  user_sex text not null, 
  user_role text not null
);
