# 1. view语句

```mysql
# 创建视图
create view v1 as select * from t1 order by id desc;
# 查看视图,此时的order by优先级更高
select * from v1 order by id;
drop view v1;

# 视图类似于创建了一张虚拟表,视图仅在原表没做修改之前才有效
select * from (select * from t1 order by id desc) as t order by id;
```

# 2. index语句

```mysql
# 创建索引
create index idx1 on t1(name);
alter table t1 add index idx1(name);
# 删除索引
drop index idx1 on t1;
alter table t1 drop index idx1;

# 修改表名 rename操作同时也会把外键关联的数据修改为当前表名
rename table t1 to t1_ori;

# 删除数据表;truncate要比delete的性能高,但是删除后数据无法rollback回滚
truncate table t2;
```

# 3. 常用函数

```mysql
select lower("LAKJSDLKF");
select upper("lasdjf");

select greatest(1,2,5,9,7) as bigger;

# 查null值
insert into t1 values(4,null);
select * from t1 where name is null;
select * from t1 where name is not null;

select * from t1 where id between 2 and 4;

select * from t1 where id = 4;
# 不等于
select * from t1 where id <> 4;
select * from t1 where id != 4;

# null不等于null
select null=null;
+-----------+
| null=null |
+-----------+
|      NULL |
+-----------+

select null is null;
+---------------+
| null  is null |
+---------------+
|             1 |
+---------------+

select null is not null;
+------------------+
| null is not null |
+------------------+
|                0 |
+------------------+

# 查询范围内任意值
select * from t1 where id in (1,3,6,9,2);

# 这只变量值
select @var1 := 1;
set @var2 = 2;

# 取变量值
select @var1;
select @var1+@var2;

# char_length()
mysql> set @var1="aslkjdf";
mysql> select char_length(@var1);
+--------------------+
| char_length(@var1) |
+--------------------+
|                  7 |
+--------------------+

select name,char_length(name) as length from student;

# concat()
mysql> select concat("aa","bb");
+-------------------+
| concat("aa","bb") |
+-------------------+
| aabb              |
+-------------------+

select concat("id: ",id,"  ","name: ", name) as id_name from student;
# 结合ifnull去除null
select concat("id: ",id,"  ","name: ",ifnull(name,"")) as id2name from stuudent;

# concat_ws 字符串拼接
select concat_ws(",","a"," b"," c");

select insert("12345678",3,4,1234);

# % 任意字符
# 以a开头,后面跟任意字符
select * from student where name like "a%";

# _ 单个字符
# 以a开头后面必须跟一个
select * from student where name like "a_";

# \ 转义字符

select 5 div 2;
select 5 % 2;
select abs(-10);

mysql> select ceiling(1.2);
+--------------+
| ceiling(1.2) |
+--------------+
|            2 |
+--------------+

mysql> select floor(1.2);
+------------+
| floor(1.2) |
+------------+
|          1 |
+------------+

# 返回0-1之间的小数
select rand();

# rand()
select * from student order by rand();
select *, rand() as r from student order by r;

# 取10-20之间的小数
select 10 + rand()*10;

# 取10-20之间的整数
select floor(10 + rand()*10);

select curdate();
select now();
select sysdate();
select curtime(); <=> select current_time();
select unix_timestamp();


select date_format(now(),"%Y-%m-%d %h:%i:%s");
select now() as 北京时间;

alter table student add tstamp datetime default now();
select tstamp,date_format(tstamp,"%Y-%m-%d") from student;


select * from dept where exists (select * from student where student.dept_id=dept.id);
select * from dept where id in (select dept_id from student);

explain select * from student;

mysql> select count(0);
+----------+
| count(0) |
+----------+
|        1 |
+----------+

mysql> select sum(0);
+--------+
| sum(0) |
+--------+
|      0 |
+--------+

mysql> select count("");
+-----------+
| count("") |
+-----------+
|         1 |
+-----------+

mysql> select count(null);
+-------------+
| count(null) |
+-------------+
|           0 |
+-------------+

mysql> select sum(null);
+-----------+
| sum(null) |
+-----------+
|      NULL |
+-----------+

# limit 的偏移量是0, 从第四行还是,向后取5行
select * from student order by id limit 3,5;

union和union all的区别是什么?
union是最后的合并结果提出重复行
union all是显示所有的结果

select id,name from student union select id,name from dept;
select id,name from student union all select id,name from dept;
```

### 流程控制语句

```mysql
# case
# nullif()
select id, name, case gender when 1 then "male" when 0 then "femle" else "unknow" end gender, dept_id from student;

# if()
select if(1>2,"yes","no");

select if(1>0,"true","false");

select if(gender=1,"male","female") from student;
select name, if(gender=1,"male","female") as gender from student;

# ifnull()
select name,ifnull(name,"unknow") as name1 from student;

mysql> select ifnull(1/0,"yes") as result;
+--------+
| result |
+--------+
| yes    |
+--------+

ltrim()
rtrim()
trim()
repeat()
replace()
substring()
locate()

select name, substring(name,1,locate(" ",name)-1) as first_name,substring(name,locate(" ",name)+1,10) as last_name from t2;


```

