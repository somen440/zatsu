-- 日付毎のグループ化
select
    count(*) as num
    -- @see https://www.w3schools.com/sql/func_mysql_date_format.asp
    , date_format(added, '%Y') as year
from authors
group by year;
