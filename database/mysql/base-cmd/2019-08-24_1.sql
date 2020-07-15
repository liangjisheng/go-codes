-- alter table teatalter_tbl modify c char(10); 
-- alter table teatalter_tbl change j i bigint;
-- alter table teatalter_tbl change j j int;
-- alter table teatalter_tbl modify i bigint not null default 100;
-- alter table teatalter_tbl modify c char(10) not null default 'test';
-- alter table teatalter_tbl alter j set default 1000;
-- show table status like 'teatalter_tbl';
-- alter table teatalter_tbl rename to testalter_tbl;
-- alter table testalter_tbl add index (c);
-- alter table testalter_tbl drop index c;
-- show index from testalter_tbl;
-- show tables;

-- show create table runoob_tbl;
-- CREATE TABLE `clone_tbl` (
--   `runoob_id` int(10) unsigned NOT NULL AUTO_INCREMENT,
--   `runoob_title` varchar(100) NOT NULL,
--   `runoob_author` varchar(40) NOT NULL,
--   `submission_date` date DEFAULT NULL,
--   PRIMARY KEY (`runoob_id`)
-- ) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8

-- insert into clone_tbl (runoob_id, runoob_title, runoob_author, submission_date) select runoob_id, runoob_title, runoob_author, submission_date from runoob_tbl;

