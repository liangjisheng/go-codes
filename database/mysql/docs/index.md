# index

索引的几个建议

1. 在建表时就应该考虑添加索引,如: 外键字段,等等
2. 在写完SQL后,一定要查看执行计划。尽量避免全表扫描
3. 如果是已有表中添加索引,一定要先计算该字段的区分度
4. 联合索引,将区分度大放在前面
5. 遵从MySQL左列前缀优先原则

索引区分度: 指字段在数据库中的不重复比

```sql
select count(distinct(name))/count(*) from t_base_user;
```

其中区分度最大值为1.000,最小为0.0000,区分度的值越大,也就是数据不重复率越大，新建索引效果也越好,
在主键以及唯一键上面的区分度是最高的,为1.0000

单列索引: 可以查看该字段的区分度,根据区分度的大小,也能大概知道在该字段上的新建索引是否有效，以及效果如何。区分度越大,索引效果越明显

多列索引(联合索引) 多列索引中其实还有一个字段的先后顺序问题,一般是将区分度较高的放在前面,这样联合索引才更有效

最左前缀匹配原则 MySQL会一直向右匹配直到遇到范围查询(>、<、between、like)就停止匹配

```sql
select * from t_base_user where type="10" and created_at<"2017-11-03" and status=1;
```

在上述语句中,status就不会走索引,因为遇到<时,MySQL已经停止匹配,此时走的索引为:(type,created_at),其先后顺序是可以调整的,
而走不到status索引,此时需要修改语句为

```sql
select * from t_base_user where type=10 and status=1 and created_at<"2017-11-03";
```

尽量不要在索引列上,进行函数运算,否则索引会失效
扩展优先,不要新建索引,尽量在已有索引中修改

## 概念

聚簇索引, 其实就是我们常见的主键索引，是InnoDB存储引擎中对主键索引的别名而已。表中记录的存储都是按照主键列进行有序组织的，
主键索引的添加，删除操作成本是非常高的。要经过: 创建临时表，复制数据等一系列操作。

二级索引就是我们常见字段上的索引，由于二级索引的创建，删除，不涉及创建临时表，复制数据等操作，所以其修改成本相对低很多。效率也更高
二级索引只包含了索引键和主键值，所以在创建，删除时不需要进行数据复制，从而不会进行锁表。也就允许线上DML操作。不影响生产使用

主键索引和唯一索引非常相似。相同点：该列中的数据都不能有相同值；不同点：主键索引不能有null值，但是唯一索引可以有null值。

```sql
-- 创建单值索引
create index idx_dept on tb(dept);
alter table tb add index dept_index(dept);
-- 创建唯一索引
create unique index name_index on tb(name);
alter table tb add unique index name_index(name);
-- 创建复合索引
create index dept_name_index on tb(dept,name);
alter table tb add index dept_name_index(dept,name);
-- 索引删除
drop index name_index on tb;
-- 索引查询
show index from tb;
```
