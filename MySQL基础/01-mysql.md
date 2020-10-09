```mysql
 CREATE TABLE test (
    user_name VARCHAR(30)
);
create table t1(
	a int,
    b int,
    c int
);
insert into t1 values (1,2,3),(2,3,4);
```

## 创建表

```mysql
create table table_name (
	列名1 数据类型,
    列名2 数据类型,
    列名3 数据类型,
)
```



##### 特殊变量(关键字)使用反引号转义一下: ``,

```mysql
-- create database test;
create database if not exists test;
use test;
drop database if exists test;

create table num_test (
	num_a int,
    num_b int zerofill,
    num_c int(3),
    num_d int(3) zerofill,
    num_e int(11),
    num_f int(11) zerofill,
    num_g int(100) zerofill
);

insert into num_test values (1,1,1,1,1,1,1);
-- show warnings;

-- select * from num_test;

CREATE TABLE products (
    PRIMARY KEY (product_id),
    product_id INT(11) NOT NULL AUTO_INCREMENT,
    name VARCHAR(100) DEFAULT NULL,
    image_url VARCHAR(255) DEFAULT NULL,
    sales_status ENUM('NEW', 'DEPRECATED') DEFAULT 'NEW',
    price DECIMAL(9 , 2 ) DEFAULT NULL,
    create_date DATETIME
);
-- desc products;
insert into products values
(1,"海苔岩烧大鸡腿饭A2","","NEW",37,"2019-12-12 10:10:00"),
(2,"夏日缤纷桶","","NEW",88,"2019-12-11 10:10:00"),
(3,"新奥尔良烤鸡腿堡人气套餐B","","NEW",88,"2019-12-10 10:10:00");

insert into products (name, price, create_date) values
("超级塔克午饭套餐",27,now());

alter table products add column detail varchar(100) not null after name;
alter table products add column title varchar(100) not null first;
alter table products add column other varchar(100) not null;

-- 删除列
alter table products drop column title;
select * from products; 

select * from products where price > 60;
-- 不等于的符号<>
select * from products where price <> 88;
select name from products where price between 20 and 40;

select * from products where price in (36,88);

select * from products where price=27 or price=88;

select * from products where image_url is null;

select * from products where image_url is not null;

-- 模糊匹配
select * from products where name like "%鸡腿%";
-- 以...开头
select * from products where name like "香辣%";
-- 以... 结尾
select * from products where name like "%桶";

-- 相对模糊匹配 _下划线表示单个字符
select * from products where name like "_____";

select * from products where price not between 20 and 60;

-- 按照某个字段排序 升序 asc(可以忽略不写,默认升序)
select * from products order by price;
-- 按照某个字段排序 降序
select * from products order by price desc;

-- 排序且限制column
select * from products order by price desc limit 3;


select * from products where create_date > "2019-12-11" order by create_date, price desc;
select * from products where create_date > "2019-12-11" order by create_date asc, price desc limit 3;

-- 别名
select `name`, price as org_price, price+10 as adjusted_price from products;

show variables like "%secure%";
set global require_secure_transport = "on";

-- 如果图片链接为空,则显示n/a,否则直接显示图片链接,结果用别名重新定义作为展示 
select product_id, `name`, if(image_url is null, "N/A", image_url) as image_url_detail from products;
select product_id, `name`,image_url, if(image_url is null, "N/A", image_url) as image_url_detail from products;

SELECT 
    `name`,
    price,
    (CASE
        WHEN price < 20 THEN '0-20'
        WHEN price >= 20 AND price < 30 THEN '20 - 30'
        WHEN price >= 30 AND price < 40 THEN '30 - 40'
        WHEN price >= 40 AND price < 50 THEN '40 - 50'
        ELSE '50+'
    END) AS price_group
FROM
    products;



```



```mysql
-- 大体上SQL语句查询是这样的
select column1, column2, ...
from table1, table2, ...
(where) condition1 and condition2 or condition3 ...
(group by) column name
(having) condition
(order by) sorting
(limit)
```

