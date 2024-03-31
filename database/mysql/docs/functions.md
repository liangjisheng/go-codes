# functions

mysql 常用函数

```sql
-- 当前 db 的 collation
show variables like 'collation_database';
-- 所有可用 collation
show collation;

-- 找到 name 中包含 a 或者 A 字符的记录
select name from table_name where locate('A', name);
-- 只包含 A 或者 a
select name from table_name where locate('A', cast(name as binary));
select name from table_name where locate('a', cast(name as binary));

-- 找到 name 中包含 a 或者 A 字符的记录
select name from table_name where instr(name, 'A');
-- 只包含 A 或者 a
select name from table_name where instr(cast(name as binary), 'A');
select name from table_name where instr(cast(name as binary), 'a');

-- binary 返回二进制串
select binary('test');
-- 默认不区分大小写 返回 1
select 'TEST' = 'test';
-- Byte-by-Byte 使用 binary 比较 分别返回 0,1
select binary('TEST') = 'test';
select binary('TEST') = 'TEST';

-- cast() 函数用于将值从一种数据类型转换为另一种特定数据类型
select cast("2019-11-21" as date);
select cast(121 as char);
select cast(2-4 as signed);
select cast('3.12' as signed);
select cast('30a1.12' as signed);
select cast(2-4 as unsigned);
select cast(1231 as binary);
select cast(12.34 as decimal(3,1));

-- cast() 和 convert() 两者都是进行数据类型转换, 作用基本等同
select convert(121, char);
select convert(2-4, char);
select convert(2-4, signed);
select convert('123', signed);
select convert(1231, binary);

select md5('xyz');
select md5('xyz123');
select md5('test');
select md5('NULL');
select md5('');
-- 返回 null
select md5(null);

-- coalesce() 函数接受一个参数，该参数是可以包含各种值的列表。 coalesce()函数返回的值是表达式列表中的第一个非空值；如果列表中的所有值均为NULL，则为NULL
select coalesce(NULL, 'A', 'B', NULL);
select coalesce('A', NULL, 'B', NULL);
select coalesce(NULL, 1, 2, 3, NULL, 'B', NULL);
select coalesce(NULL, NULL, 'test', NULL);

-- if() 函数用于验证条件。如果条件为TRUE，则 if() 函数返回一个值，如果条件为FALSE，则返回另一个值
select if(5<12, 'true', 'false');
select if(strcmp('test', 'tes')=0, 'true', 'false');
select if(5<12, '1', '0');

-- LAST_DAY()函数返回有效日期参数的月份的最后一天。如果参数Date无效或为null，则该函数还将返回NULL
select last_day('2023-03-23');
select last_day('2023-03-23 07:22:20');
select last_day('2023-02-23');
select now();
select last_day(now());
select last_day(curdate());
-- 提取下个月的最后一天
select last_day(curdate() + interval 1 month);

select isnull(null);
select isnull('test');
select isnull(123);

-- 0x0D95FC9469EADFE7569D2FD35B8B76CC
select aes_encrypt('abc', 'key_string');
select aes_decrypt(aes_encrypt('abc', 'key_string'), 'key_string');
select convert(aes_decrypt(aes_encrypt('abc', 'key_string'), 'key_string'), char);

-- bin() 将十进制数转换为等效的二进制数
select bin(5);
select bin(12);
select bin(99);
-- 将数字 5 由 10 进制转换为 2 进制
select conv(5, 10, 2);
select conv(12, 10, 2);
select conv(99, 10, 2);
```

```sql
-- 显示指定数据库的创建语句
show create database dbname;
show create schema name;
-- 显示表中所有列信息
show full columns tables_name;
-- 查看MySQL版本
select version();
-- 查看当前用户
select current_user();
-- 显示单表信息
show table status like "table_name";
-- 显示正在操作数据库的进程数
show processlist;
-- 显示表中的所有索引
show index from t_base_data;
-- 查看查询语句的执行情况，常用于SQL优化
-- explain 查询语句
-- 显示当前时间
select now();
-- 显示指定字符长度
select char_length('test');
-- 格式化日期
select date_format(now(), '%y-%m-%d');
select DATE_FORMAT(now(), '%y-%m-%d %H:%i:%s');
-- 添加/减少日期时间 
-- DATE_ADD(date,interval expr unit) DATE_SUB(date,interval expr unit)
-- unit：表示单位，支持毫秒(microsecond)，秒(second)，小时(hour)，天(day)，周(week)，年(year)等
select date_add(now(), interval 1 day);
-- 类型转换
select cast(18700000000 as char);
-- md5(data)
select md5("test");
-- 字符串连接
select concat("test","big");
-- 获取系统当前时间的时间戳，类型: long 单位: s
select unix_timestamp(now()), unix_timestamp(current_timestamp()), unix_timestamp(sysdate());
-- 此时时间精度是s，也可以增加精度，给函数加上参数，表示s后面的小数位数，例如参数3，此时为ms
select unix_timestamp(now(3)), unix_timestamp(current_timestamp(3)), unix_timestamp(sysdate(3));
-- 如果直接输出毫秒单位的时间戳，就是去掉上面中间的小数点，可以借助 replace 函数
select replace(unix_timestamp(now(3)), '.', ''), replace(unix_timestamp(current_timestamp(3)), '.', ''),
       replace(unix_timestamp(sysdate(3)), '.', '');

-- 获取系统当前时间，类型：timestamp 格式yyyy-MM-dd HH:mm:ss
-- 2019-01-04 20:37:19
-- 三者基本没有区别，稍微一点的区别在于：NOW(),CURRENT_TIMESTAMP()都表示SQL开始执行的时间；SYSDATE()表示执行此SQL时的当前时间
select now(), current_timestamp(), sysdate();
select NOW(),CURRENT_TIMESTAMP(),SYSDATE(),SLEEP(2),NOW(),CURRENT_TIMESTAMP(),SYSDATE();
```

concat 将多个字符串连接成一个字符串, 返回结果为连接参数产生的字符串，如果有任何一个参数为null，则返回值为null

```sql
select concat(id, address, mission_name) as info from address_nft;
-- 可以加一个逗号作为分隔符
select concat(id, ',', address, ',', mission_name) as info from address_nft;
```

concat_ws 和 concat()一样，将多个字符串连接成一个字符串，但是可以一次性指定分隔符

```sql
select concat_ws(',', id, address, mission_name) as info from address_nft;
```

group_concat 将group by产生的同一个分组中的值连接起来，返回一个字符串结果
group_concat([distinct] 要连接的字段 [order by 排序字段 asc/desc  ] [separator '分隔符'])

```sql
select address, group_concat(mission_name) from address_nft group by address;
select address, group_concat(concat_ws(',', mission_name, status)) from address_nft group by address;
```

```sql
-- mysql time
select unix_timestamp('2022-03-23 18:00:00');
select unix_timestamp();
select from_unixtime(1648029600, '%Y-%m-%d %H:%i:%S');
select from_unixtime(1648029600);
```
