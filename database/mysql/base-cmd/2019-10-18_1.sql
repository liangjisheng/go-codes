-- use test_db;
-- create table tb_dept1 (
-- 	id int(11) primary key,
--     name varchar(50) not null,
--     location varchar(100)
-- );

-- create table tb_emp6 (
-- 	id int(11) primary key,
--     name varchar(50),
--     deptID int(11),
--     salary float,
--     constraint fk_emp_dept1
--     foreign key(deptID) references tb_dept1(id)
-- );

-- desc tb_emp2; 
-- alter table tb_emp2 add constraint fk_tb_dept1 foreign key(deptID) references tb_dept1(id);
-- show create table tb_emp2;
-- 删除外键
-- alter table tb_emp2 drop foreign key fk_tb_dept1;

-- 唯一约束
-- create table tb_dept2 (
-- 	id int(11) primary key,
--     name varchar(22) unique,
--     location varchar(50)
-- );

-- alter table tb_dept1 add constraint unique_name unique(name);
-- desc tb_dept1;
-- alter table tb_dept1 drop index unique_name;