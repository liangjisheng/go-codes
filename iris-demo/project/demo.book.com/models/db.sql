create table `book_tb` (
    `ID` int(11) not null auto_increment,
    `BookName` varchar(100) not null default '' comment 'bookname',
    `State` int(11) not null default 0 comment '状态',
    `Author` varchar(100) not null default '' comment 'author',
    `Press` varchar(100) not null default '' comment '出版社',
    `PublishTime` timestamp not null default current_timestamp comment '出版时间',
    `BookImage` varchar(300) not null default '' comment '图书封面',
    `Price` decimal(10,2) not null default 0.00 comment '售价',
    `Introduction` varchar(300) not null default '' comment '简介',
    `UpdateTime` timestamp not null default current_timestamp comment '',
    `AddTime` timestamp not null default current_timestamp comment '',
    primary key (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='图书表'