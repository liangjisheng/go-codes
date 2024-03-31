# lua

[post](https://shidawuhen.github.io/2021/10/29/Redis%E7%9A%84%E4%BA%8B%E5%8A%A1%EF%BC%8CGo-lua%E7%94%A8%E8%B5%B7%E6%9D%A5%E7%9C%9F%E9%A6%99/)
[post](https://studygolang.com/articles/19703)

```lua
-- this is a comment

-- 全局变量
x = 123

本地变量
local y = 456

function hello_world()
 return "Hello World"
end

for i = 1, 10 do
   print(i)
end

if x == 123 then
    print("x is the magic number")
else
    print("I have no idea what x is")
end

-- 字符串连接
print("Hello" .. " World")

-- 数组的索引从1开始
data_types = {1.0, 123, "redis", true, false, hello_world}
print(data_types[3]) -- the output is "redis"

languages = {lua = 1993, javascript = 1995, python = 1991, ruby =1995}
print("Lua was created in " .. languages["lua"])
print("JavaScript was created in " .. languages.javascript)
```

Redis 提供了简单的事务功能，将一组需要一起执行的命令放到 multi 和 exec 两个命令之间
multi 命令代表事务开始， exec 命令代表事务结束，它们之间的命令是原子顺序执行的
如果要停止事务的执行，可以使用 discard 命令代替 exec 命令即可

Redis事务作为一个整体被执行，执行期间不会被其它客户端打断
Redis事务在执行的中途遇到错误，不会回滚，而是继续执行后续命令
Redis里遇到有查询的情况穿插在事务中间，不会返回结果

命令错误
该种类型属于语法错误，会造成整个事务无法执行，key 和 value 的值未发生变化

运行时错误
如命令写错了，误把 sadd 写为 zadd, 该种类型会执行成功，且不支持回滚。需要开发人员自己进行修复

Redis2.6 之后新增的功能，我们可以在 Redis 中通过 lua 脚本操作 Redis
脚本会将多个命令和操作当成一个命令在 Redis 中执行，也就是说该脚本在执行的过程中，不会被任何其他脚本或命令打断干扰
正是因此这种原子性，lua 脚本才可以代替 multi 和 exec 的事务功能。同时也是因此，在lua脚本中不宜进行过大的开销操作，避免影响后续的其他请求的正常执行

使用lua脚本的好处
lua脚本是作为一个整体执行的，所以中间不会被其他命令插入
可以把多条命令一次性打包，所以可以有效减少网络开销
lua脚本可以常驻在redis内存中，所以在使用的时候,可以直接拿来复用，也减少了代码量

```shell
127.0.0.1:6379> eval 'return "hello " .. KEYS[1] .. ARGV[1]' 1 redis world
```

当脚本会被多次执行时，可考虑使用ScriptLoad和EvalSha代替RUN节省带宽。
先用命令ScriptLoad将脚本缓存到Redis，Redis返回一个sha1的标识符
命令EvalSha基于sha1执行脚本
这种方案只有标识符sha1通过网络传输，而不需传输lua代码块，节省流量s

除了使用eval， Redis还提供了evalsha命令来执行Lua脚本。 首先要将Lua脚本加载到Redis服务端， 得到该脚本的SHA1校验和，
evalsha命令使用SHA1作为参数可以直接执行对应Lua脚本， 避免每次发送Lua脚本的开销。 这样客户端就不需要每次执行脚本内容
而脚本也会常驻在服务端， 脚本功能得到了复用

lua_get.lua

```lua
return "hello " .. KEYS[1] .. ARGV[1]
```

```shell
#将脚本内容加载到 Redis 内存中，得到 SHA1
redis-cli script load "$(cat lua_get.lua)"
#evalsha 脚本 SHA1 值 key个数 key列表 参数列表
127.0.0.1:6379> evalsha 7413dc2440db1fea7c0a0bde841fa68eefaf149c 1 redis world

#此命令用于将 Lua 脚本加载到 Redis 内存中
127.0.0.1:6379> script load script
#判断 sha1 是否已经加载到 Redis 内存中
127.0.0.1:6379> script exists
清除Redis内存已经加载的所有Lua脚本 
127.0.0.1:6379> script flush
#杀掉正在执行的Lua脚本。如果Lua脚本比较耗时， 甚至Lua脚本存在问题， 那么此时Lua脚本的执行会阻塞Redis，
#直到脚本执行完毕或者外部进行干预将其结束。 (如果当前Lua脚本正在执行写操作， 那么script kill将不会生效)
127.0.0.1:6379> script kill
```

lua-time-limit 默认5s，当Lua脚本时间超过lua-time-limit后， 向其他命令调用发送BUSY的信号， 但是并不会停止掉服务端和客户端的脚本执行，
所以当达到lua-time-limit值之后， 其他客户端在执行正常的命令时， 将会收到“Busy Redis is busy running a script”错误， 并且提示使用script kill或者shutdown nosave命令来杀掉这个busy的脚本

