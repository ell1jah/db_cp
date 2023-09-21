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
