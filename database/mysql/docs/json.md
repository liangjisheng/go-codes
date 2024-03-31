# json

mysql json

json 列存储的必须是 json 格式数据，否则会报错
json 数据类型是没有默认值的

```sql
create table tab_base_info (
    id bigint not null primary key auto_increment,
    content json
)ENGINE=InnoDB CHARSET=utf8mb4;

insert into tab_base_info (content) value ('{"author": "Captain&D", "blog": "https://google.com"}');
```

```sql
-- 指定数据转换json字符串
select json_type('"hello"');
select json_type('["c++","Java","go"]');
select json_object("name","alice","database","MySQL");
select json_array("name","alice","database","MySQL");
select json_array(1, now(), "test");
-- 判断是否为json字符串的json_valid()
select json_valid('{"name": "alice", "database": "MySQL"}');
-- json_unquote 函数与 json_quote 函数作用正好相反
select json_quote('"test"');
select json_merge('{"blog":"www.baidu.com"}','{"account":"test"}');
```

```sql
-- 获取字符串中的某个 key
select content->"$.blog" from tab_base_info;
select content->>"$.blog" from tab_base_info;
-- 用户提取 json 属性 json_extract（json 串, 操作符）
select json_extract('{"blog":"www.baidu.com","name": "test"}' ,'$.blog');
-- json_insert
set @info = '{"blog":"www.baidu.com","name": "test"}';
select json_insert(@info,"$.email","andytohome@gmail.com");
-- json_replace
set @info = '{"blog":"www.baidu.com","name": "test"}';
select json_replace(@info, '$.name', "alice");
-- json_remove
select json_remove('{"blog":"www.baidu.com","name": "test"}' , '$.blog') as res;
```
