# postgresql

切换到 postgres linux 系统用户
sudo su postgres

sudo passwd -d postgres  # 清空 linux 用户 postgres 的密码
sudo -u postgres passwd  # 设置密码 接下来按要求输入两次密码, 要与上面的密码相同

备份 plasma 到 ~/postgres.sql
-U 指定 postgres 用户, -f 指定导出的文件名, plasma 为数据库
pg_dump -U postgres -f ~/postgres.sql plasma
pg_dump -U postgres plasma > ~/database.sql
pg_dump -U postgres plasma -t table > ~/table.sql

2种方式恢复 ~/postgres.sql 到 plasma 数据库
psql -U postgres -f ~/postgres.sql plasma
psql -U postgres -d plasma < ~/postgres.sql

远程登录 pg
psql -h 172.18.1.196 -p 5432 -U postgres
psql -U dbuser -d exampledb -h 127.0.0.1 -p 5432
psql -U liangjisheng -d exampledb -h 127.0.0.1 -p 5432

从24机器登录23的pg 密码为123456
psql -U postgres -d postgres -h 172.20.101.23 -p 5432

备份数据库
pg_basebackup -h 172.20.101.23 -U postgres -D /var/lib/postgresql/13/data/ -X stream -P
pg_basebackup -h 172.20.101.23 -U liangjisheng -w -F p -P -X stream -R -D /var/lib/postgresql/13/main/

```shell
# 设置密码,2种方式都可以
\password
alter user postgres with password '123456';
# 列出所有数据库
\l
# 连接其他数据库
\c
# 查看所有用户
\du
# 列出当前数据库和连接的信息
\conninfo
# 退出
\q
# Expanded display is on.
\x on;
```

创建用户和数据库,方法1使用postgres用户登录pg,在上面进行操作
一个role相当于一个用户或者用户组
create user liangjisheng with password '123456';
create role liangjisheng login replication encrypted password '123456';
create database exampledb owner liangjisheng;

将exampledb数据库的所有权限赋予 liangjisheng, 否则 liangjisheng 只能登录psql, 没有任何数据库操作权限
grant all privileges on database exampledb to liangjisheng;

方法2, 安装PostgreSQL后提供了createuser和createdb命令行程序
/usr/lib/postgresql/10/bin/
首先创建数据库用户 liangjisheng1 并指定为超级用户
下面这个命令是在shell终端上执行, 而不是像上面一样需要登录pg

sudo -u postgres createuser --superuser liangjisheng1;

接着登录psql控制台设置其密码(123456)后退出
\password liangjisheng1;

然后创建数据库并指定所有者
sudo -u postgres createdb -O liangjisheng1 exampledb1;

删除用户的话要先删除所有属于这个用户的数据库
drop database exampledb1;
drop user liangjisheng1;

基本操作命令

```shell
# 创建新表
create table user_tbl(name varchar(20), signup_date date);
# 插入数据
insert into user_tbl(name, signup_date) values('alice', '2021-08-11');
# 选择记录
select * from user_tbl;
# 更新数据
update user_tbl set name = 'xxx' where name = 'alice';
# 删除记录
delete from user_tbl where name = 'xxx';
# 添加栏位
alter table user_tbl add email varchar(40);
# 更新结构
alter table user_tbl alter column signup_date set not null;
# 更名栏位
alter table user_tbl rename column signup_date to signup;
# 删除栏位
alter table user_tbl drop column email;
# 表格更名
alter table user_tbl rename to backup_tbl
# 删除表格
drop table if exists backup_tbl;
```

清空表并保留表结构

当表没有其他关系时
TRUNCATE TABLE tablename;
当表中有外键时，要用级联方式删所有关联的数据
TRUNCATE TABLE tablename CASCADE;

清空数据库还原数据库为新建时的状态
在postgresql中, 创建数据库时会自动创建public模式
一般我们把表都保存在该模式中, 因此直接删除该模式再重新创建该模式

删除public模式以及模式里面所有的对象
DROP SCHEMA public CASCADE;
创建public模式
CREATE SCHEMA public;

```sql
create table company (
    id int primary key not null,
    name text not null,
    age int not null,
    address char(50),
    salary real,
    join_date date
);

insert into company (id, name, age, address, salary, join_date) values (1, 'Paul', 32, 'California', 20000.00,'2001-07-13');
insert into company (id, name, age, address, salary, join_date) values (2, 'Allen', 25, 'Texas', 15000,'2002-07-13');
update company set join_date = '2021-11-21' where id = 1;
update company set salary = 75000 where id = 1;
update company set age = 55 where id = 2;
insert into company (id, name, age, address, salary, join_date) values (4, 'Mark', 25, 'Rich-Mond ', 65000.00, '2007-12-13' ), (5, 'David', 27, 'Texas', 85000.00, '2007-12-13');
insert into company (id, name, age, address, salary, join_date) values (6, 'Kim', 22, 'South-Hall ', 45000.00, '2003-12-13' ), (7, 'James', 24, 'Houston', 10000.00, '2004-12-13');

select * from company where salary != 25000;
select * from company where salary <> 25000;
select * from company where salary is not null;
select count(*) as "records" from company;
select * from company where name like 'Pa%';
select * from company where age in (25, 27);
select * from company where age not in (25, 27);
select * from company where age between 25 and 27;
select age from company where exists (select age from company where salary > 65000);
select * from company where age > (select max(age) from company where salary > 65000);
select * from company limit 3;
select * from company limit 3 offset 2;
select * from company order by age asc;
select * from company order by name, salary asc;
select * from company order by name desc;

group by 子句必须放在 where 子句中的条件之后, 必须放在 order by 子句之前
select name, sum(salary) from company group by name;
insert into company values (8, 'Paul', 24, 'Houston', 20000.00, '2005-12-13');
insert into company values (9, 'James', 44, 'Norway', 5000.00, '2006-12-13');
insert into company values (10, 'James', 45, 'Texas', 5000.00, '2007-12-13');
select name, sum(salary) from company group by name order by name;

create table company1 (
    id int primary key not null,
    name text not null,
    age int not null,
    address char(50),
    salary real,
    join_date date
);

with moved_rows as (
    delete from company
    where
        salary >= 30000
    returning *
)
insert into company1 (select * from moved_rows);
select * from company1;

having 子句可以让我们筛选分组后的各组数据
where 子句在所选列上设置条件, 而 having 子句则在由 group by 子句创建的分组上设置条件
having 子句必须放置于 group by 子句后面, order by 子句前面

select name from company group by name having count(name) < 2;
select name from company group by name having count(name) > 1;

select distinct name from company;

create table department (
    id int primary key not null,
    dept char(50) not null,
    emp_id int not null
);

\d 查看表格
\d tablename 查看表格信息

create schema myschema;
create table myschema.company(
   ID   INT              NOT NULL,
   NAME VARCHAR (20)     NOT NULL,
   AGE  INT              NOT NULL,
   ADDRESS  CHAR (25),
   SALARY   DECIMAL (18, 2),
   PRIMARY KEY (ID)
);
select * from myschema.company;

select 2+3;
select (17 + 6) as addition;
select 2*3;
select 10/5;
select 12%5;
select 2^3;
select |/ 25.0; 平方根
select ||/ 27.0; 立方根
select 5 !; 阶乘
select !!5; 阶乘(前缀操作符)

位运算
select 60 | 13;
select 60 & 13;
select (~60);
select (60 << 2);
select (60 >> 2);
select 60 # 13;

select current_timestamp;
```

约束

```sql
create table company6 (
    id int primary key not null,
    name text not null,
    age int not null unique,
    address char(50) not null,
    salary real default 5000 check(salary > 0)
);

create table department1(
    id int primary key not null,
    dept char(50) not null,
    emp_id int references company6(id)
);

/* exclusion 约束
需要为每个数据库执行一次 create extension btree_gist 命令，这将安装 btree_gist 扩展，它定义了对纯标量数据类型的 EXCLUDE 约束
*/

create table company7 (
    id int primary key not null,
    name text not null,
    age int not null,
    address char(50) not null,
    salary real default 5000 check(salary > 0),
    exclude using gist (name with =, age with <>)
);

/* 强制执行了年龄必须相同，让我们通过向表插入记录来查看这一点 */
insert into company7 values (1, 'Paul', 32, 'California', 20000.00);
insert into company7 values (2, 'Paul', 32, 'Texas', 20000.00);
insert into company7 values (3, 'Paul', 42, 'Texas', 20000.00);

删除约束
alter table table_name drop constraint some_name;
```

join

```sql
company表
https://static.runoob.com/download/company.sql

insert into company values (8, 'Paul', 24, 'Houston', 20000.00);
insert into company values (9, 'James', 44, 'Norway', 5000.00);
insert into company values (10, 'James', 45, 'Texas', 5000.00);

drop table if exists department;
create table department(
    id int primary key not null,
    dept char(50) not null,
    emp_id int not null
);
insert into department values (1, 'IT Billing', 1);
insert into department values (2, 'Engineering', 2);
insert into department values (3, 'Finance', 7);

交叉连接（CROSS JOIN）把第一个表的每一行与第二个表的每一行进行匹配。如果两个输入表分别有 x 和 y 行，则结果表有 x*y 行
由于交叉连接（CROSS JOIN）有可能产生非常大的表，使用时必须谨慎，只在适当的时候使用它们
select emp_id, name, dept from company cross join department;

内连接（INNER JOIN）根据连接谓词结合两个表（table1 和 table2）的列值来创建一个新的结果表。查询会把 table1 中的每一行与 table2 中的每一行进行比较，找到所有满足连接谓词的行的匹配对
当满足连接谓词时，A 和 B 行的每个匹配对的列值会合并成一个结果行
内连接（INNER JOIN）是最常见的连接类型，是默认的连接类型
select emp_id, name, dept from company inner join department on company.id = department.emp_id;

外部连接是内部连接的扩展。SQL 标准定义了三种类型的外部连接: LEFT、RIGHT 和 FULL, PostgreSQL 支持所有这些
对于左外连接，首先执行一个内连接。然后，对于表 T1 中不满足表 T2 中连接条件的每一行，其中 T2 的列中有 null 值也会添加一个连接行。因此，连接的表在 T1 中每一行至少有一行
select emp_id, name, dept from company left outer join department on company.id = department.emp_id;

右外连接,首先执行内部连接。然后，对于表T2中不满足表T1中连接条件的每一行，其中T1列中的值为空也会添加一个连接行。这与左联接相反;对于T2中的每一行，结果表总是有一行
select emp_id, name, dept from company right outer join department on company.id = department.emp_id;

外连接,首先执行内部连接。然后，对于表 T1 中不满足表 T2 中任何行连接条件的每一行，如果 T2 的列中有 null 值也会添加一个到结果中。此外，对于 T2 中不满足与 T1 中的任何行连接条件的每一行，将会添加 T1 列中包含 null 值的到结果中
select emp_id, name, dept from company full outer join department on company.id = department.emp_id;
```

union

UNION 操作符合并两个或多个 SELECT 语句的结果
UNION 操作符用于合并两个或多个 SELECT 语句的结果集
请注意，UNION 内部的每个 SELECT 语句必须拥有相同数量的列。列也必须拥有相似的数据类型。同时，每个 SELECT 语句中的列的顺序必须相同
UNION ALL 操作符可以连接两个有重复行的 SELECT 语句，默认地，UNION 操作符选取不同的值。如果允许重复的值，请使用 UNION ALL

```sql
select emp_id, name, dept from company inner join department on company.id = department.emp_id union select emp_id, name, dept from company left outer join department on company.id = department.emp_id;

select emp_id, name, dept from company inner join department on company.id = department.emp_id union all select emp_id, name, dept from company left outer join department on company.id = department.emp_id;
```

null 无法比较 NULL 和 0 因为它们是不等价的

```sql
update company set address = null, salary = null where id = 6;
select id, name, age, address, salary from company where salary is null;
select id, name, age, address, salary from company where salary is not null;
```

别名

```sql
select c.id, c.name, c.age, d.dept from company as c, department as d where c.id = d.emp_id;
select c.id as company_id, c.name as company_name, c.age, d.dept from company as c, department as d where c.id = d.emp_id;
```

触发器

```sql
create table audit (
    company_id int not null,
    entry_date text not null
);

创建pg过程函数
create or replace function auditlogfunc() returns trigger as $example_table$
  begin
    insert into audit values (new.id, current_timestamp);
    return new;
  end;
$example_table$ language plpgsql;

创建触发器
create trigger example_trigger after insert on company for each row execute procedure auditlogfunc();

insert into company values (11, 'James', 45, 'Texas', 5000.00);
select * from audit;

查看触发器
select * from pg_trigger;
select tgname from pg_trigger, pg_class where tgrelid = pg_class.oid and relname = 'company';

drop trigger example_trigger on company;
```

index

```sql
create index salary_index on company (salary);
\d company;
\di
```

alter

```sql
alter table table_name add column_name datatype;
alter table table_name drop column column_name;
alter table table_name alter column type datatype;
alter table table_name alter column_name datatype not null;
/* 添加唯一约束 */
alter table table_name add constraint myUniqueConstraint unique(column1, column2...);
alter table table_name add constraint myUniqueConstraint check (condition);
alter table table_name add constraint myPrimaryKey primary key (column1, column2...);
/* 删除约束 */
alter table table_name drop constraint myUniqueConstraint;
/* 删除主键 */
alter table table_name drop constraint myPrimaryKey;
/* mysql */
alter table table_name drop index myUniqueConstraint;
alter table table_name drop primary key;
```

view

```sql
create view company_view as select id, name, age from company;
select * from company_view;
drop view company_view;
```

transaction

```sql
begin;
delete from company where age = 25;
rollback;
/* rollback 后仍然有记录, 使用 commit 提交事务 */
```

子查询

```sql
select * from company where id in (select id from company where salary > 45000);

CREATE TABLE IF NOT EXISTS company_bkp (
    id integer primary key NOT NULL,
    name text COLLATE pg_catalog."default" NOT NULL,
    age integer NOT NULL,
    address character(50) COLLATE pg_catalog."default",
    salary real
);
insert into company_bkp select * from company where id in (select id from company);
select * from company_bkp;

update company set salary = salary * 0.5 where age in (select age from company_bkp where age > 27);
delete from company where age in (select age from company_bkp where age > 27);
```

smallserial 2bytes, serial 4bytes, bigserial 8bytes

```sql
create table if not exists company2(
    id serial primary key,
    name text not null,
    age int not null,
    address char(50),
    salary real
);
insert into company2 (name, age, address, salary) values ('Paul', 32, 'California', 20000.00);
insert into company2 (name, age, address, salary) values ('Allen', 25, 'Texas', 15000.00);
insert into company2 (name, age, address, salary) values ('Teddy', 23, 'Norway', 20000.00);
insert into company2 (name, age, address, salary) values ('Mark', 25, 'Rich-Mond ', 65000.00);
insert into company2 (name, age, address, salary) values ('David', 27, 'Texas', 85000.00);
insert into company2 (name, age, address, salary) values ('Kim', 22, 'South-Hall', 45000.00);
insert into company2 (name, age, address, salary) values ('James', 24, 'Houston', 10000.00);
select * from company2;
```

privilege

```sql
GRANT privilege [, ...]
ON object [, ...]
TO { PUBLIC | GROUP group | username }

privilege − 值可以为：SELECT，INSERT，UPDATE，DELETE， RULE，ALL。
object − 要授予访问权限的对象名称。可能的对象有： table， view，sequence。
PUBLIC − 表示所有用户。
GROUP group − 为用户组授予权限。
username − 要授予权限的用户名。PUBLIC 是代表所有用户的简短形式

REVOKE privilege [, ...]
ON object [, ...]
FROM { PUBLIC | GROUP groupname | username }
```
