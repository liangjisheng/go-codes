-- show create database test_db;
-- create database test_db_del;
-- show databases;
-- drop database test_db_del;
-- show engines;
-- use test_db;
-- create table tb_emp1 (
--     id int(11) comment '员工编号',
--     name varchar(25) comment '员工名称',
--     deptID int(11) comment '所在部门编号',
--     salary float comment '工资'
-- );
-- show tables;
-- describe tb_emp1;
-- desc tb_emp1;
-- show create table tb_emp1;

-- 增加字段
-- alter table tb_emp1 add column col1 int first;
-- desc tb_emp1;
-- alter table tb_emp1 add column col2 int after name;
-- desc tb_emp1;

-- 修改字段数据类型
-- alter table tb_emp1 modify name varchar(30) not null;
-- desc tb_emp1;

-- 删除字段
-- alter table tb_emp1 drop col2;
-- desc tb_emp1;

-- 修改字段名称
-- alter table tb_emp1 change col1 col3 char(30);
-- desc tb_emp1;

-- 修改表名
-- alter table tb_empl rename tb_emp2;
-- show tables;

-- 删除表
-- create table tb_emp3(
--     id int(11),
--     name varchar(25),
--     deptID int(11),
--     salary float
-- );
-- show tables;
-- drop table tb_emp3;
-- show tables;
