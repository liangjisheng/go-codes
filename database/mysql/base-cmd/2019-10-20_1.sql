-- create table tb_emp7 (
-- 	id int(11) primary key,
--     name varchar(25),
--     deptID int(11),
--     salary float,
--     check (salary>0 and salary<100),
--     foreign key(deptID) references tb_dept1(id)
-- );

-- alter table tb_emp7 add constraint check_id check(id>0);
-- show create table tb_emp7; 

-- create table tb_dept3 (
-- 	id int(11) primary key,
--     name varchar(25),
--     location varchar(50) default 'Beijing'
-- );

-- alter table tb_dept3 change column location location varchar(50) default 'Shanghai';
-- 删除默认值约束
-- alter table tb_dept3 change column location location varchar(50) default null;
-- desc tb_dept3;

-- create table tb_dept4(
-- 	id int(11) primary key,
--     name varchar(22) not null,
--     location varchar(50)
-- );
-- alter table tb_dept4 change column location location varchar(50) not null;
-- 删除非空约束
-- alter table tb_dept4 change column location location varchar(50) null;
-- desc tb_dept4;
-- show create table tb_dept4;

-- create table tb_emp8(
-- 	id int (11) primary key,
--     name varchar(22) unique,
--     deptID int(11) not null,
--     salary float default 0,
--     check(salary>0),
--     foreign key(deptID) references tb_dept1(id)
-- );
-- show create table tb_emp8;
