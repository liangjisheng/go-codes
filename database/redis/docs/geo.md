# geo

Redis的GEO(地理空间)数据类型是一种特殊的数据结构，从Redis 3.2版本开始引入，用于存储地理位置信息并执行地理位置相关的操作。

GEOADD:用于添加一个或多个地理位置到指定的键中。每个位置由其经度、纬度和一个关联的成员（如地点名称）定义。
GEOPOS:根据成员查询并返回一个或多个地理位置的坐标。
GEOHASH:将地理位置的坐标转化为一个唯一的字符串编码（GeoHash），这有助于快速索引和查询。
GEODIST:计算两个地理位置之间的距离，可选的距离单位有m（米）、km（千米）、mi（英里）、ft（英尺）。
GEORADIUS:查询指定位置周围的地点，基于给定的最大距离（半径）。
GEORADIUSBYMEMBER:类似于GEORADIUS，但不是基于坐标而是基于集合中的一个成员来执行查询。

因为GEO类型主要用于存储地理位置信息.因此在使用GEO数据类型之前,要获取一下某个位置的经纬度信息.
可以去以下几个网址进行查询:

https://lbs.qq.com/getPoint/  
https://lbs.amap.com/tools/picker  
https://api.map.baidu.com/lbsapi/getpoint/index.html  

GEOADD key longitude1 latitude1 member1 [longitude2 latitude2 member2 ...]:用于添加一个或多个地理位置到指定的键中。

GEO类型本质上是Zset类型

```redis
127.0.0.1:6379> geoadd city 121.47 31.23 "上海" 116.41 39.90 "北京"
(integer) 2
127.0.0.1:6379> type city
zset
127.0.0.1:6379> zrange city 0 -1
1) "\xe4\xb8\x8a\xe6\xb5\xb7"
2) "\xe5\x8c\x97\xe4\xba\xac"
127.0.0.1:6379>
```

注意上面使用zrange遍历city时,中文乱码了. 解决起来也很简单. 只需要在连接redis时加上 --raw

```redis
127.0.0.1:6379> quit
[root@localhost bin]# redis-cli -a 123123 --raw
Warning: Using a password with '-a' or '-u' option on the command line interface may not be safe.
127.0.0.1:6379> zrange city 0 -1
上海
北京
127.0.0.1:6379>
```

GEOPOS key member1 [member2 ...]:根据成员查询并返回一个或多个地理位置的坐标。

GEOHASH key member1 [member2 ...]:将地理位置的坐标转化为一个唯一的字符串编码（GeoHash），这有助于快速索引和查询。

```redis
127.0.0.1:6379> geopos city 北京
116.40999823808670044
39.90000009167092543
127.0.0.1:6379> 127.0.0.1:6379> geohash city 北京 上海
wx4fbzx4me0
wtw3sj5zbj0
127.0.0.1:6379>
```

GEODIST key member1 member2 [M|KM|FT|MI]:计算两个地理位置之间的距离，可选的距离单位有
m（米）、km（千米）、mi（英里）、ft（英尺）。

```redis
127.0.0.1:6379> geodist city 北京 上海 km
1066.9923
127.0.0.1:6379>
```

GEORADIUS key longitude latitude radius M|KM|FT|MI [WITHCOORD] [WITHDIST] [WITHHASH] [COUNT count [ANY]] [ASC|DESC] [STORE key] [STOREDIST key]
查询指定位置周围的地点，基于给定的最大距离（半径）。

命令参数含义：

key: 存储地理位置信息的键。  
longitude: 中心点的经度。  
latitude: 中心点的纬度。  
radius: 查询半径，可以是具体的数值加上单位（m 米, km 千米, mi 英里, ft 英尺）。  
[WITHCOORD]: 如果设置，返回的位置元素将包含它们的坐标。  
[WITHDIST]: 如果设置，返回的位置元素将包含它们与中心点的距离。  
[WITHHASH]: 如果设置，返回的位置元素将包含它们的GeoHash值。  
[COUNT count]: 限制返回的结果数量。  
[ASC|DESC]: 指定结果排序方式，ASC 表示按距离从近到远排序（默认），DESC 则相反。  

GEORADIUSBYMEMBER key member radius M|KM|FT|MI [WITHCOORD] [WITHDIST] [WITHHASH] [COUNT count [ANY]] [ASC|DESC] [STORE key] [STOREDIST key]
基于集合中的一个成员来执行查询。

```redis
127.0.0.1:6379> georadius city 121.48 31.23 10 km withdist withcoord withhash count 10 desc
上海
0.9509
4054803462927619
121.47000163793563843
31.22999903975783553
127.0.0.1:6379> georadiusbymember city "上海" 10 km withdist withcoord withhash count 10 desc
上海
0.0000
4054803462927619
121.47000163793563843
31.22999903975783553
127.0.0.1:6379>
```
