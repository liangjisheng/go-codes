drop table if exists `doctor_tb`;
CREATE TABLE `doctor_tb` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT 'name',
  `age` int NOT NULL DEFAULT '0' COMMENT 'age',
  `sex` int NOT NULL DEFAULT '0' COMMENT 'sex',
  `ctime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `mtime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8 COMMENT='医生表'

insert into `doctor_tb` (`name`, `age`, `sex`) values('doctor-alice', 24, 0);
insert into `doctor_tb` (`name`, `age`, `sex`) values('doctor-alice1', 25, 0);
insert into `doctor_tb` (`name`, `age`, `sex`) values('doctor-alice2', 26, 0);

drop table if exists `user_tb`;
create table `user_tb` (    
    `id` int(11) not null auto_increment,
    `name` varchar(50) not null default '' comment 'name',
    `age` int(11) not null default 0 comment 'age',
    `sex` int(11) not null default 0 comment 'sex',
    `ctime` timestamp not null default current_timestamp comment '创建时间',
    `mtime` timestamp not null default current_timestamp comment '修改时间',
	primary key (`id`)
) engine=Innodb auto_increment=67 default charset=utf8 comment='用户表';

insert into `user_tb` (`name`, `age`, `sex`) values ("user-alice", 24, 0);
insert into `user_tb` (`name`, `age`, `sex`) values ("user-alice1", 25, 0);
insert into `user_tb` (`name`, `age`, `sex`) values ("user-alice2", 26, 0);
