# sql 优化

[优化&explain](https://mp.weixin.qq.com/s/7GFyu2rpOviMF5yMC37TZg)
[查询执行计划](https://juejin.cn/post/7000319871896518663)

[explain](http://www.cnitblog.com/aliyiyi08/archive/2008/09/09/48878.html)
[explain](https://zhuanlan.zhihu.com/p/114182767)

优化：添加索引的时候，要根据MySQL解析顺序添加索引，又回到了MySQL的解析顺序，下面我们再来看看MySQL的解析顺序。
看mysql的编写过程和解析过程, 很显然，where后是group by，然后才是select

```sql
-- 编写过程
select dinstinct  ..from  ..join ..on ..where ..group by ..having ..order by ..limit ..
-- 解析过程
from .. on.. join ..where ..group by ..having ..select dinstinct ..order by ..limit ..
```

有几点可以注意下

复合索引，不要跨列或无序使用(最佳左前缀)，复合索引，尽量使用全索引匹配，也就是说，你建立几个索引，就使用几个索引
复合索引中如果有【>】，则自身和右侧索引全部失效。

不要在索引上进行任何操作(计算、函数、类型转换)，否则索引失效。

索引不能使用不等于（!= <>）或is null (is not null)，否则自身以及右侧所有全部失效(针对大多数情况)。复合索引中
如果有>，则自身和右侧索引全部失效。

like尽量以“常量”开头，不要以’%'开头，否则索引失效，like尽量不要使用类似"%x%"情况，但是可以使用"x%"情况。
如果非使用 "%x%"情况，需要使用索引覆盖。也就是 select 的字段只包含索引字段。

尽量不要使用类型转换（显示、隐式），否则索引失效。

对于主键连续时而且允许的情况下，我们甚至可以使用max(id)来代替count(*)来统计用户数。

用 in 代替 or，尽量不要使用or，否则索引失效。

## explain

通过 EXPLAIN [SQL] SHOW WARNINGS; 查看优化后的SQL
EXPLAIN EXTENDED [SQL;] SHOW WARNINGS; 就可以看到优化器重构后的查询语句

如果想更加精确地展示执行计划，可以在 EXPLAIN 后面加上 FORMAT=JSON 来将输出格式化为一颗树。
在这棵树中，可以看到每步执行的成本。
explain FORMAT=JSON select * from nft_token order by id desc limit 1;

```sql
-- 一个简单的查询
select * from nft_token order by id desc limit 1;
-- 查看上述语句的执行计划
explain select * from nft_token order by id desc limit 1;
```

### id

id: 表示查询中select操作表的顺序,按顺序从大到小依次执行

### select_type

select_type: 该表示选择的类型,可选值有: 

SIMPLE 就是非常简单的查询，意味着不包含 UNION 或子查询，连接查询也算是 SIMPLE 类型

PRIMARY 如果查询有任何复杂的子部分，如 UNION、UNION ALL、子查询，则最外层部分标记为 PRIMARY

UNION 在 UNION 或 UNION ALL 子句中的第二个之后的 SELECT 都会被标记为 UNION

UNION RESULT UNION 查询需要使用内存临时表来去重，针对该临时表的查询的 select_type 就是 UNION RESULT

SUBQUERY 包含在 SELECT 列表中的子查询中的 SELECT 会标记为 SUBQUERY，不包含 FROM 子查询

DERIVED 用来表示包含在 FROM 子句的子查询中的 SELECT，MySQL 会将子查询的结果物化，放到一个临时表中，
也称为 派生表，因为该临时表是从子查询中派生来的

DEPENDENT 意味着 SELECT 依赖于外层查询中发现的数据

MATERIALIZED 意味着SQL中包含子查询（非FROM子查询），且子查询是物化的方式和外层查询进行连接查询的

### table

table: 一般情况下它就是访问的那个表，或是该表的别名, 当有 FROM 子查询、UNION 查询时，这一列则不同，显示的是物化表

### partitions

partitions 这一列跟表分区相关，一般情况下查询语句的执行计划的 partitions 列的值都是 NULL

### type

type: 这一列表示这行的访问类型或者访问方法，就是MySQL决定如何查找表中的行。这个属性是分析执行计划时需要
重点关注的信息，它告诉我们SQL是如何查询这张表的数据，有没有用到索引等等。

下面是访问方法，依次从最优到最差
NULL 这种访问方式意味着MySQL能在优化阶段分解查询语句，在执行阶段甚至用不着再访问表或者索引

system 当表中只有一条记录并且该表使用的存储引擎的统计数据是精确的，比如 MyISAM、Memory，那么对该表的访问方法就是 system

const 当根据主键或者唯一二级索引列与常数进行等值匹配时，对单表的访问方法就是 const。出现 const 表明查询是极为快速的

eq_ref 在连接查询时，如果被驱动表是通过主键或者唯一二级索引列等值匹配的方式进行访问的，则对该被驱动表的访问方法
就是 eq_ref，被驱动表最多只会返回一条符合条件的记录

ref 当通过普通二级索引列与常量进行等值匹配时来查询某个表，那么对该表的访问方法就可能是 ref。这是一种索引访问方式，
相比 eq_ref，ref 可能返回多条符合条件的记录

ref_or_null 当对普通二级索引进行等值匹配查询，该索引列的值也可以是NULL值时，那么对该表的访问方法就可能是 ref_or_null

index_merge 一般情况下对于某个表的查询只会用到一个索引，但也可能同时使用多个索引，使用索引合并的方式来执行查询，
这时访问方法就可能是 index_merge

unique_subquery 在一些包含 IN 子查询的SQL语句中，如果查询优化器决定将 IN 子查询转换为 EXISTS
子查询，而且子查询可以使用到主键进行等值匹配的话，那么该子查询使用的方法就是 unique_subquery

index_subquery 类似于 unique_subquery，只不过访问子查询中的表时如果使用的是普通索引，则访问方法是 index_subquery

range 表示基于索引的范围扫描，返回索引某些范围区间的记录。这比 index 全索引扫描好一些，因为它不用遍历全部索引

index 类似于 ALL 全表扫描，不过 index 是顺序扫描索引的全部记录。它的主要优点是避免了排序；
但如果要回表的话，开销可能比 ALL 更大

ALL 就是全表扫描，通常意味着MySQL必须扫描整张表，应尽量避免这种方法的出现

### possible_keys

possible_keys: 顾名思义,该属性给出了,该查询语句,可能走的索引,(如某些字段上索引的名字)这里提供的只是参考,
而不是实际走的索引,也就导致会有possible_Keys不为null,key为空的现象

### key

key : 实际使用到的索引，也可能不在 possible_keys 中

### key_len

索引的长度，用于判断复合索引是否被完全使用(a,b,c)

### ref

这里的ref的作用，指明当前表所参照的字段。
注意与type中的ref值区分。在type中，ref只是type类型的一种选项值。
ref: 连接匹配条件,如果走主键索引的话,该值为: const, 全表扫描的话,为null值, 当使用索引列等值匹配
的条件去执行查询时，ref 列显示与索引列作等值匹配的对象，比如是常量、某个字段，或者是一个函数等

### rows

rows:扫描行数,也就是说，需要扫描多少行, 才能获取目标行数,一般情况下会大于返回行数。通常情况下,rows越小,效率越高, 
也就有大部分SQL优化，都是在减少这个值的大小，理想情况下扫描的行数与实际返回行数理论上是一致的,但这种情况及其少,
如关联查询,扫描的行数就会比返回行数大大增加。 表示使用指定的查询方式（type）会读出来的行数，注意它并不是最终要从
表里读取出来的行数，只是一个预估值

### filtered

filtered：表示针对读出来的行数（rows），再用上其它的查询条件（WHERE子句或连接条件），又可以过滤出百分之几的数据来

### Extra

Extra 这一列包含的是不适合在其他列显示的额外信息，这些额外信息可用来更准确的理解MySQL将如何执行给定的查询语句。
这个属性非常重要,该属性中包括执行SQL时的真实情况信息,如上面所属,使用到的是"using where"，表示使用where筛选
得到的值,常见的一些额外信息如下

Impossible WHERE 查询语句的 WHERE 子句永远为 FALSE 时将会提示该额外信息

```sql
-- 查看执行计划
explain select a1 from test02 where a1="a" and a1="b" ;
```

Using Index 表示使用了 覆盖索引，没有 回表 查询，就是仅仅涉及到一个二级索引。注意和 type=index 的区别，
覆盖索引是说查询只用到了二级索引，而没有回表查询。index 是全索引扫描。

```sql
-- 删除test02中的复合索引idx_a1_a2_a3
drop index idx_a1_a2_a3 on test02;
-- 重新创建一个复合索引idx_a1_a2
create index idx_a1_a2 on test02(a1,a2);
-- 查看执行计划
explain select a1,a3 from test02 where a1='' or a3= '' ;
explain select a1,a2 from test02 where a1='' and a2= '' ;
```

结果分析：我们创建的是a1和a2的复合索引，对于第一个执行计划，我们却出现了a3，该字段并没有创建索引，因此没有出现
using index，而是using where，表示我们需要回表查询。对于第二个执行计划，属于完全的索引覆盖，因此出现了using index

Using where 一般在没有使用到索引，执行全表扫描时，并且 WHERE 子句中有针对该表的搜索条件，就会提示
Using where。或者用到了索引，但还有其它 WHERE 条件，也可能会提示 Using where
表示需要【回表查询】，表示既在索引中进行了查询，又回到了源表进行了查询

```sql
-- 删除test02中的复合索引idx_a1_a2
drop index idx_a1_a2 on test02;
-- 将a1字段，新增为一个索引
create index a1_index on test02(a1);
-- 查看执行计划
explain select a1,a3 from test02 where a1="" and a3="" ;
```

结果分析：我们既使用了索引a1，表示我们使用了索引进行查询。但是又对于a3字段，我们并没有使用索引，
因此对于a3字段，需要回源表查询，这个时候出现了 using where。

Using temporary 出现 Using temporary 意味着MySQL内部使用了临时表来完成一些功能，比如去重、排序之类的，
前面演示 UNION 时，因为 UNION 要对结果集去重，所以就使用了临时表
当出现了这个词，也表示你当前的SQL性能消耗较大。这是由于当前SQL用到了临时表。一般出现在group by中。

```sql
explain select a1 from test02 where a1 in ('1','2','3') group by a1 ;
explain select a1 from test02 where a1 in ('1','2','3') group by a2 ; --using temporary
```

结果分析：当你查询哪个字段，就按照那个字段分组，否则就会出现using temporary

```sql
explain select * from test03 where a2=2 and a4=4 group by a2,a4;
explain select * from test03 where a2=2 and a4=4 group by a3;
```

分析如下：对于第一个执行计划，where后面是a2和a4，接着我们按照a2和a4分组，很明显这两张表已经有了，直接在
a2和a4上分组就行了。但是对于第二个执行计划，where后面是a2和a4，接着我们却按照a3分组，很明显我们没有a3
这张表，因此有需要再来一张临时表a3。因此就会出现using temporary。

Using filesort 有些排序可以直接使用索引的排序，但很多时候可能无法使用索引排序，这时就需要将数据查到内存中，
或者在数据了较大时放到磁盘中，再进行排序。这种在内存中或者磁盘上进行排序的方式统称为文件排序（filesort），
在 Extra 列就会显示 Using filesort，要注意文件排序这个过程是很耗费性能的，当出现 Using filesort 时，
最好能优化SQL或索引，将文件排序的执行方式改为使用索引进行排序

using filesort：针对单索引的情况
当出现了这个词，表示你当前的SQL性能消耗较大。表示进行了一次“额外”的排序。常见于order by语句中。

什么是“额外”的排序？为了讲清楚这个，我们首先要知道什么是排序。我们为了给某一个字段进行排序的时候，
首先你得先查询到这个字段，然后在将这个字段进行排序。
紧接着，我们查看如下两个SQL语句的执行计划。

```sql
-- 新建一张表，建表同时创建索引
create table test02
(
    a1 char(3),
    a2 char(3),
    a3 char(3),
    index idx_a1(a1),
    index idx_a2(a2),
    index idx_a3(a3)
);
-- 查看执行计划
explain select * from test02 where a1 ='' order by a1 ;
explain select * from test02 where a1 ='' order by a2 ; 
```

结果分析：对于第一个执行计划，where后面我们先查询了a1字段，然后再利用a1做了依次排序，这个很轻松。
但是对于第二个执行计划，where后面我们查询了a1字段，然而利用的却是a2字段进行排序，此时myql底层会
进行一次查询，进行“额外”的排序。

总结：对于单索引，如果排序和查找是同一个字段，则不会出现using filesort；如果排序和查找不是同一个
字段，则会出现using filesort；因此where哪些字段，就order by哪些些字段。

using filesort：针对复合索引的情况
不能跨列(官方术语：最佳左前缀)

```sql
-- 删除test02的索引
drop index idx_a1 on test02;
drop index idx_a2 on test02;
drop index idx_a3 on test02;
-- 创建一个复合索引
alter table test02 add index idx_a1_a2_a3 (a1,a2,a3) ;
-- 查看下面SQL语句的执行计划
explain select *from test02 where a1='' order by a3 ;  --using filesort
explain select *from test02 where a2='' order by a3 ; --using filesort
explain select *from test02 where a1='' order by a2 ;
```

结果分析：复合索引的顺序是(a1,a2,a3)，可以看到a1在最左边，因此a1就叫做“最佳左前缀”，如果要使用
后面的索引字段，必须先使用到这个a1字段。对于explain1，where后面我们使用a1字段，但是后面的排序
使用了a3，直接跳过了a2，属于跨列；对于explain2，where后面我们使用了a2字段，直接跳过了a1字段，
也属于跨列；对于explain3，where后面我们使用a1字段，后面使用的是a2字段，因此没有出现【using filesort】

Using join buffer (Block Nested Loop) 在连接查询执行过程中，当被驱动表没有索引或者不能有效的利用索引
加快访问速度时，MySQL一般会为其分配一块名叫 join buffer 的内存块，然后将数据读到这块内存块中，再基于块中
的数据做嵌套循环查询，以此来加快查询速度
