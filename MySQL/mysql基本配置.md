# MySQL-Community-Server

# 1.安装

```bash
wget https://dev.mysql.com/get/mysql80-community-release-el8-1.noarch.rpm

dnf install mysql80-community-release-el8-1.noarch.rpm

dnf install mysql-server

systemctl enable --now mysqld
```



```mysql
# 修改密码
alter user "root"@"localhost" identified with mysql_native_password by "666";

create user "root"@"10.0.0.%" identified with mysql_native_password by "666";

# 打开配置文件删除前面的注释(很多软件还不支持新版的密码规则,
# 这里继续使用原来的密码规则)
vim /etc/my.cnf
25 default-authentication-plugin=mysql_native_password

# 在log文件中找到临时密码
cat /var/log/mysqld.log
A temporary password is generated for root@localhost: .tFxoMw#Q53,

# 登录MySQL
    root@me ~$ mysql -u root -p
    Enter password:

# 把该密码复制进去登录MySQL

Welcome to the MySQL monitor.  Commands end with ; or \g.
Your MySQL connection id is 21
Server version: 8.0.18 MySQL Community Server - GPL

Copyright (c) 2000, 2019, Oracle and/or its affiliates. All rights reserved.

Oracle is a registered trademark of Oracle Corporation and/or its
affiliates. Other names may be trademarks of their respective
owners.

Type 'help;' or '\h' for help. Type '\c' to clear the current input statement.

mysql> 

出现该页面表示登录成功!


mysql> show variables like 'validate_password%';
+--------------------------------------+--------+
| Variable_name                        | Value  |
+--------------------------------------+--------+
| validate_password.check_user_name    | ON     |
| validate_password.dictionary_file    |        |
| validate_password.length             | 8      |
| validate_password.mixed_case_count   | 1      |
| validate_password.number_count       | 1      |
| validate_password.policy             | MEDIUM |
| validate_password.special_char_count | 1      |
+--------------------------------------+--------+


# 临时修改密码规格和长度

set global validate_password.policy=low;
set global validate_password.length=3; 

show variables like 'validate_password%';

# 默认长度最短为4个字符
mysql> show variables like 'validate_password%';
+--------------------------------------+-------+
| Variable_name                        | Value |
+--------------------------------------+-------+
| validate_password.check_user_name    | ON    |
| validate_password.dictionary_file    |       |
| validate_password.length             | 4     |
| validate_password.mixed_case_count   | 1     |
| validate_password.number_count       | 1     |
| validate_password.policy             | LOW   |
| validate_password.special_char_count | 1     |
+--------------------------------------+-------+


# 修改密码
alter user 'root'@'localhost' identified with mysql_native_password by '666';

# 查看用户权限
show grants for root@'localhost';

# 可以利用该语句授予所有权限给指定用户(比如root >^_^< )
grant all privileges on *.* to 'root'@'localhost';

# 查看用户
select user,host from mysql.user;

# 修改host名为%(任何ip都可以登录,这样是不太安全的!!)
update mysql.user set host='%' where user='root';

select user,host from mysql.user;

mysql> select host,user from mysql.user;
+-----------+------------------+
| host      | user             |
+-----------+------------------+
| %         | root             |
| localhost | mysql.infoschema |
| localhost | mysql.session    |
| localhost | mysql.sys        |
+-----------+------------------+


# 之后修改密就不能再@localhost了而是@%.
alter user 'root'@'%' identified with mysql_native_password by '6666';

# 查看权限也是
show grants for root@'%';


# 查看数据表
show databases;
mysql> show databases;
+--------------------+
| Database           |
+--------------------+
| information_schema |
| mysql              |
| performance_schema |
| sys                |
+--------------------+
4 rows in set (0.00 sec)

# 查看登录的用户
select user();

flush privileges;


# 新建数据表
create database wordpress;


# 使用数据表
use wordpress;

# 创建新用户
create user 'wordpress'@'localhost' identified by '6666';

# 查看该用户的权限
show grants for wordpress@localhost;

# 修改host名为%(任何ip都可以登录)
update mysql.user set host='%' where user='wordpress';

# 授所有wordpress表的权限该wordpress用户
grant all privileges on wordpress.* to 'wordpress'@'%';

# 确认权限信息
show grants for wordpress@'%';

show grants for wordpress@'%';
mysql> show grants for wordpress@'%';
+---------------------------------------+
| Grants for wordpress@%                |
+---------------------------------------+
| GRANT USAGE ON *.* TO `wordpress`@`%` |
+---------------------------------------+


select version();


# 登录后修改密码
alter user 'root'@'localhost' identified by '666';
set password for root@localhost='666';

# 未登录修改密码
mysqladmin -u root password '666'
mysqladmin -u root -p '555' password '666'
mysqladmin -u root -h web01 password '666'



```

