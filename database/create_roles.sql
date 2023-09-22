create user "default_guest";
create user "default_user";
create user "default_admin";

alter role "default_guest" password '11111111';
alter role "default_user" password '12344321';
alter role "default_admin" password '12345678';

grant select on table grandprix to "default_guest";
grant select on race_results_view to "default_guest";
grant select on drivers_of_season to "default_guest";

grant select on table drivers to "default_user";
grant select on table grandprix to "default_user";
grant select on table qualificationresults to "default_user";
grant select on table raceresults to "default_user";
grant select on table season_standings to "default_user";
grant select on table teams to "default_user";
grant select on table teamsdrivers to "default_user";
grant select on table tracks to "default_user";
grant select on race_results_view to "default_user";
grant select on drivers_of_season to "default_user";

alter role "default_admin" superuser;

-- drop owned by "default_user";
-- drop owned by "default_admin";
-- drop user "default_user";
-- drop user "default_admin";