# sql 优化

有几点可以注意下

1. 通过explain 来查看执行计划。通过这一步骤，我们能够分析出，该语句有没有走索引，索引合不合理的重要依据
2. 缩小范围。例如使用 < > ，between …and。来缩小扫描范围
3. 减少连接数量 (对于连接查询，我们必须尽可能减少每个子连接的结果集数量，只包含有效数据
4. 避免类型转换， 隐式类型转换是最容易疏忽的慢SQL
5. 对于主键连续时而且允许的情况下，我们甚至可以使用max(id)来代替count(*)来统计用户数
6. 用 in 代替 or， 少用like，避免使用函数运算

## explain

```sql
-- 一个简单的查询
select * from nft_token order by id desc limit 1;
-- 查看上述语句的执行计划
explain select * from nft_token order by id desc limit 1;
```

id: 表示查询中select操作表的顺序,按顺序从大到小依次执行

select_type: 该表示选择的类型,可选值有: SIMPLE(简单的)

table: 表示该语句查询的表

type: 该属性表示访问类型,有很多种访问类型。最常见的其中包括以下几种: ALL(全表扫描), index(索引扫描),range(范围扫描), 
ref (非唯一索引扫描),eq_ref(唯一索引扫描,),(const)常数引用, 访问速度依次由慢到快  
range(范围)常见与 between and ..., 大于 and 小于这种情况

possible_keys: 顾名思义,该属性给出了,该查询语句,可能走的索引,(如某些字段上索引的名字)这里提供的只是参考,
而不是实际走的索引,也就导致会有possible_Keys不为null,key为空的现象

key : 显示MySQL实际使用的索引,其中就包括主键索引(PRIMARY),或者自建索引的名字

key_len : 表示索引所使用的字节数

ref: 连接匹配条件,如果走主键索引的话,该值为: const, 全表扫描的话,为null值

rows : 扫描行数,也就是说，需要扫描多少行,采能获取目标行数,一般情况下会大于返回行数。通常情况下,rows越小,效率越高, 
也就有大部分SQL优化，都是在减少这个值的大小，理想情况下扫描的行数与实际返回行数理论上是一致的,但这种情况及其少,
如关联查询,扫描的行数就会比返回行数大大增加

Extra 这个属性非常重要,该属性中包括执行SQL时的真实情况信息,如上面所属,使用到的是"using where"，表示使用where筛选得到的值,
常用的有: "Using temporary": 使用临时表 "using filesort": 使用文件排序
