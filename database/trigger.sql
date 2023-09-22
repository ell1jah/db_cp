-- Получение количества очков по гоночному результату гонщика
create or replace function GetScore(place int)
returns int
as $$
begin
    if place = 1 then
        return 25;
    elsif place = 2 then
        return 18;
    elsif place = 3 then
        return 15;
    elsif place = 4 then
        return 12;
    elsif place = 5 then
        return 10;
    elsif place = 6 then
        return 8;
    elsif place = 7 then
        return 6;
    elsif place = 8 then
        return 4;
    elsif place = 9 then
        return 2;
    elsif place = 10 then
        return 1;
    else return 0;
    end if;
end
$$ language plpgsql;

-- Получение сезона по id гонки
create or replace function GetSeason(id int)
returns int
as $$
declare res int;
begin
    select gp_season
    into res
    from raceresults r
    join grandprix g on g.gp_id = r.gp_id
    where race_id = id;
    return res;
end
$$ language plpgsql;


-- Функция триггера
create or replace function UpdateTrigger()
returns trigger
as $$
begin
    raise notice 'New =  %, season = %, driver = %', new, GetSeason(new.race_id), new.driver_id;
    if GetSeason(new.race_id) = 2022 then
        update season_standings
        set score = score + GetScore(new.race_driver_place)
        where driver_id = new.driver_id;
    end if;
    return new;
end
$$ language plpgsql;

-- Триггер
create trigger update_season_standing
after insert on raceresults
for each row
execute procedure UpdateTrigger();


create or replace function DeleteTrigger()
returns trigger
as $$
begin
    if GetSeason(old.race_id) = 2022 then
        update season_standings
        set score = score - GetScore(old.race_driver_place)
        where driver_id = old.driver_id;
    end if;
    return old;
end
$$ language plpgsql;

create trigger delete_season_standing
before delete on raceresults
for each row
execute procedure DeleteTrigger();