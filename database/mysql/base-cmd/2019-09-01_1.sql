-- use test_db;
-- create table tb_emp3(
--     id int(11) not null primary key,
--     name varchar(25),
--     deptID int(11),
--     salary float
-- );
-- desc tb_emp3;

-- create table if not exists tb_emp4(
--     id int(11) not null,
--     name varchar(25),
--     deptID int(11),
--     salary float,
--     primary key (id)
-- );
-- desc tb_emp4;

-- create table if not exists tb_emp5(
--     name varchar(25),
--     deptID int(11),
--     salary float,
--     primary key (name, deptID)
-- );