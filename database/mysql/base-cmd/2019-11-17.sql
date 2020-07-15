-- create table tb_stu_info (
-- 	id int not null,
-- 	name char(50) default null,
-- 	dept_id int default null,
-- 	age int default null,
-- 	height int default null,
-- 	index(height)
-- );

-- show create table tb_stu_info;

-- create table tb_stu_info2 (
-- 	id int not null,
-- 	name char(50) default null,
-- 	age int default null,
-- 	height int default null,
-- 	unique index(height)
-- );

-- show create table tb_stu_info2;

-- show index from tb_stu_info2;

-- drop index height on tb_stu_info;

-- alter table tb_stu_info2 drop index height;

-- create user 'ljs1'@'localhost' identified by 'ljs1';

-- rename user ljs1@'localhost' to ljs2@'localhost';

-- rename user ljs2@'localhost' to ljs1@'localhost';

-- set password for 'ljs1'@'localhost'= password('ljs1');

-- drop user 'ljs1'@'localhost';