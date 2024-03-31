# varchar

[varchar](https://blog.csdn.net/liujinghu/article/details/121332879)
[varchar](https://blog.duhbb.com/2022/04/23775.html)

varchar(n) 指的字符数, 最大长度为 65535 个字节
lanti1 (定长, 1字节): 可存65535个字符
gbk(变长, 2字节): 可存32767个字符
utf8(变长, 3字节): 可存21845个字符
utf8mb4(变长, 4字节): 可存16383个字符
LENGTH: 字节长度, CHAR_LENGTH: 字符长度

mysql 默认定义 varchar(255) 而不是 varchar(256)  
如果varchar(20)，实际是需要1+20个字节的存储空间  
如果varchar(255)，实际是需要1+255个字节的存储空间  
如果varchar(256)，实际是需要2+256个字节的存储空间  
如果varchar(500)，实际是需要2+500个字节的存储空间  

utf8mb4 下，英文占用1个字节，一般汉字占3个字节，emoji表情占4个字节

当定义varchar长度小于等于255时，字符长度标识位需要一个字节, 当大于255时，字符长度标识位需要两个字节  
也就是说，当我们定义一个varchar(255)的字段时，其实它真实使用的空间是256(255+1)字节；(注意，字段非空，latin1编码)  
当我们定义一个一个varchar(256)字段时，它真实使用的空间是258(256+2)字节

```sql
mysql> create table test_varchar_length(v varchar(65536) not null) CHARSET=latin1;

ERROR 1074 (42000): Column length too big for column 'v' (max = 65535); use BLOB or TEXT instead

mysql> create table test_varchar_length(v varchar(65534) not null) CHARSET=latin1;

ERROR 1118 (42000): Row size too large. The maximum row size for the used table type, not counting BLOBs, is 65535. This includes storage overhead, check the manual. You have to change some columns to TEXT or BLOBs

create table test_varchar_length(v varchar(65533) not null) CHARSET=latin1;

Query OK, 0 rows affected (0.05 sec)

create table test_varchar_length(v varchar(65535) not null) CHARSET=utf8mb4;

1074 - Column length too big for column 'v' (max = 16383); use BLOB or TEXT instead, Time: 0.000000s

create table test_varchar_length(v varchar(16383) not null) CHARSET=utf8mb4;

Query OK, 0 rows affected (0.05 sec)
```


```sql
CREATE TABLE `encode`  (
  `id` bigint NOT NULL,
  `latin1` varchar(30) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT NULL COMMENT 'Latin1, 别名 ISO-8859-1 (定长, 1字节)',
  `gbk` varchar(30) CHARACTER SET gbk COLLATE gbk_chinese_ci NULL DEFAULT NULL COMMENT 'GBK(变长, 2字节)',
  `utf8` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'utf8(变长, 3字节)',
  `utf8mb4` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT 'utf8mb4(变长, 4字节)',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB;

INSERT INTO `encode` (`id`, `latin1`, `gbk`, `utf8`, `utf8mb4`) VALUES (1, 'abc123!@#', '中国人abc123!@#，。、', '中国人abc123!@#，。、', '中国人abc123!@#，。、❤️🍇💯');


select LENGTH(latin1), LENGTH(gbk), LENGTH(utf8), LENGTH(utf8mb4) from encode
union
select CHAR_LENGTH(latin1), CHAR_LENGTH(gbk), CHAR_LENGTH(utf8), CHAR_LENGTH(utf8mb4) from encode;
```
