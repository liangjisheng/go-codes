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

-- create user 'alice1'@'localhost' identified by 'alice1';

-- rename user alice1@'localhost' to alice2@'localhost';

-- rename user alice2@'localhost' to alice1@'localhost';

-- set password for 'alice1'@'localhost'= password('alice1');

-- drop user 'alice1'@'localhost';