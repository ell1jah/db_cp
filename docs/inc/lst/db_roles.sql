
create user "default_guest";
create user "default_user";
create user "default_admin";
alter role "default_guest" password '00000000';
alter role "default_user" password '11111111';
alter role "default_admin" password '22222222';
grant 
select 
  on table Item to "default_guest";
