-- create table tb_courses (
-- 	course_id int not null auto_increment,
--     course_name char(40) not null,
--     course_grade float not null,
--     course_info char(100) null,
--     primary key(course_id)
-- );

-- insert into tb_courses (course_id, course_name, course_grade, course_info)
-- values (1, 'Network', 3, 'Computer Network');

-- create table tb_courses_new (
-- 	course_id int not null auto_increment,
--     course_name char(40) not null,
--     course_grade float not null,
--     course_info char(100) null,
--     primary key(course_id)
-- );

-- insert into tb_courses_new (course_id, course_name, course_grade, course_info)
-- select course_id, course_name, course_grade, course_info from tb_courses;

-- update tb_courses_new set course_grade=4 where course_id=1;

-- create view view_tb_courses as select * from tb_courses;
-- select * from view_tb_courses;

-- create view v_tb_courses (s_id, s_name, s_grade, s_info)
-- as select course_id, course_name, course_grade, course_info from tb_courses;

-- select * from v_tb_courses;

-- describe v_tb_courses;
-- desc v_tb_courses;

-- alter view view_tb_courses as select course_id, course_name, course_grade from tb_courses;
-- desc view_tb_courses;
-- select * from view_tb_courses;
-- update view_tb_courses set course_grade=4 where course_id=1;
-- select * from view_tb_courses;
-- select * from tb_courses;

-- SHOW CREATE VIEW view_tb_courses;