# mongo

## 数据库

有一些数据库名是保留的，可以直接访问这些有特殊作用的数据库。

admin： 从权限的角度来看，这是"root"数据库。要是将一个用户添加到这个数据库，这个用户自动继承所有数据库的权限。一些特定的服务器端命令也只能从这个数据库运行，比如列出所有的数据库或者关闭服务器。
local: 这个数据永远不会被复制，可以用来存储限于本地单台服务器的任意集合
config: 当Mongo用于分片设置时，config数据库在内部使用，用于保存分片的相关信息。

## 文档

文档是一组键值(key-value)对(即 BSON)。MongoDB 的文档不需要设置相同的字段，并且相同的字段不需要相同的数据类型，这与关系型数据库有很大的区别，也是 MongoDB 非常突出的特点。

## 集合

集合就是 MongoDB 文档组，类似于 RDBMS （关系数据库管理系统：Relational Database Management System)中的表格。

集合存在于数据库中，集合没有固定的结构，这意味着你在对集合可以插入不同格式和类型的数据，但通常情况下我们插入集合的数据都会有一定的关联性

## user

```conf
use <database_name>

db.createUser({
  user: "testuser",
  pwd: "password123",
  roles: [{ role: "readWrite", db: "testdb" }]
})
; 验证用户
db.auth("testuser", "password123")
db.dropUser("testuser")
```

## database

注意: 在 MongoDB 中，集合只有在内容插入后才会创建，就是说，创建集合(数据表)后要再插入一个文档(记录)，集合才会真正创建。

```conf
use DATABASE_NAME
; 查看所有数据库
show dbs

; use 后数据库并没有立马创建,要向数据库中插入一些数据
; 如果集合不存在，会自动创建
db.mycollection.insertOne({ name: "Alice", age: 30 })
; 手动创建一个集合
db.createCollection("myNewCollection")
; 查看当前正在使用的数据库
db
; 删除当前正在使用的 myDatabase 数据库及其所有集合
db.dropDatabase()
```

## collection

```conf
db.createCollection("myNewCollection")
; 集合删除语法
db.collection.drop()
; 查看已有集合
show collections

; 下面是带有几个关键参数的 createCollection() 的用法：
; 创建了一个固定大小的集合，最大大小为 5MB（5242880 字节），最多存储 5000 个文档
db.createCollection("myCappedCollection", { capped: true, size: 5242880, max: 5000 });

; 在 MongoDB 中，你不需要创建集合，当你插入一些文档时，MongoDB 会自动创建集合
db.mycol2.insert({"name" : "alice"})

; 重命名集合
db.adminCommand({
  renameCollection: "sourceDb.sourceCollection",
  to: "targetDb.targetCollection",
  dropTarget: <boolean>
})
```

## documents

```conf
db.mycollection.insertOne({
    name: "Alice",
    age: 25,
    city: "New York"
})

db.mycollection.insertMany([
    { name: "Bob", age: 30, city: "Los Angeles" },
    { name: "Charlie", age: 35, city: "Chicago" }
])

db.mycollection.find()

db.mycollection.updateOne(
    { name: "Alice" },
    { $set: { age: 31 } }
)

db.mycollection.updateMany(
    { age: { $lt: 30 } },
    { $set: { status: "active" } }
)

db.mycollection.replaceOne(
    { name: "Bob" },
    { name: "Bob", age: 31 }
)

db.mycollection.findOneAndUpdate(
    { name: "Charlie" },
    { $set: { age: 36 } },
    { returnDocument: "after" }
)

db.mycollection.deleteOne({ name: "Alice" })
db.mycollection.deleteOne({ _id: ObjectId("672c85ca0132d926c27975af") })

db.mycollection.deleteMany({ status: "inactive" })

db.mycollection.findOneAndDelete(
    { name: "Charlie" },
    { projection: { name: 1, age: 1 } }
)
```

## query

```conf
; db.collection.find(query, projection)
; query：用于查找文档的查询条件。默认为 {}，即匹配所有文档。
; projection（可选）：指定返回结果中包含或排除的字段。

db.mycollection.find({ age: { $gt: 30 } })
; 按条件查找文档，并只返回指定字段
db.mycollection.find(
    { age: { $gt: 30 } },
    { name: 1, age: 1, _id: 0 }
)
; 以易读的方式来读取数据，可以使用 pretty() 方法
db.mycollection.find().pretty()

; findOne()
db.mycollection.findOne({ name: "Alice" })
; 查找单个文档，并只返回指定字段
db.mycollection.findOne(
    { name: "Alice" },
    { name: 1, age: 1, _id: 0 }
)

; 比较操作符
; $gt、$lt、$gte、$lte、$eq、$ne

; 逻辑操作符
; $and、$or、$not、$nor
; 查找年龄大于 25 且城市为 "New York" 的文档
db.mycollection.find({
    $and: [
        { age: { $gt: 25 } },
        { city: "New York" }
    ]
})

; 查找名字以 "A" 开头的文档
db.mycollection.find({ name: /^A/ })

; 可以对查询结果进行排序，按年龄降序排序
db.mycollection.find().sort({ age: -1 })

; 可以对查询结果进行限制和跳过指定数量的文档。
; 返回前 10 个文档
db.mycollection.find().limit(10)
; 跳过前 5 个文档，返回接下来的 10 个文档
db.mycollection.find().skip(1).limit(10)

; 查找年龄大于 25 且城市为 "New York" 的文档，只返回名字和年龄字段，按年龄降序排序，并返回前 10 个文档
db.mycollection.find(
    {
        $and: [
            { age: { $gt: 25 } },
            { city: "New York" }
        ]
    },
    { name: 1, age: 1, _id: 0 }
).sort({ age: -1 }).limit(10)

db.mycollection.find(
    {
        $or: [
            { age: { $gt: 25 } },
            { city: "New York" }
        ]
    },
    { name: 1, age: 1, _id: 0 }
).sort({ age: -1 }).limit(10)

; MongoDB 的 find() 方法可以传入多个键(key)，每个键(key)以逗号隔开，即常规 SQL 的 AND 条件。
db.mycollection.find({"name":"Alice", "city":"New York"}).pretty()

; 查找年龄大于 25 且城市为 "New York"，名字以 "A" 开头的文档
db.mycollection.find({
    $and: [
        { age: { $gt: 25 } },
        { city: "New York" },
        { name: { $regex: /^A/ } }
    ]
})

; 查找字段类型为字符串的文档
db.mycollection.find({ city: { $type: "string" } })
; 或使用类型代码
db.mycollection.find({ city: { $type: 2 } })
; 查找 age 字段类型为整数的文档
db.mycollection.find({ age: { $type: "int" } })
; 或使用类型代码
db.mycollection.find({ age: { $type: 16 } })
; 查找字段类型为布尔值的文档
db.mycollection.find({ isActive: { $type: "bool" } })
; 或使用类型代码
db.mycollection.find({ isActive: { $type: 8 } })
; 查找字段类型为日期的文档
db.mycollection.find({ createdAt: { $type: "date" } })
; 或使用类型代码
db.mycollection.find({ createdAt: { $type: 9 } })
; 查找 value 字段类型为字符串或整数的文档
db.mycollection.find({ value: { $type: ["string", "int"] } })
; 或使用类型代码
db.mycollection.find({ value: { $type: [2, 16] } })

; db.collection.find().sort({ field1: 1, field2: -1 })
; { field1: 1, field2: -1 }：指定要排序的字段及排序顺序。1 表示升序，-1 表示降序
; 按 age 字段升序排序
db.mycollection.find().sort({ age: 1 })
; 按 createdAt 字段降序排序
db.mycollection.find().sort({ createdAt: -1 })
; 先按 age 字段升序排序，再按 createdAt 字段降序排序
db.mycollection.find().sort({ age: 1, createdAt: -1 })
```

## index

[index](https://www.runoob.com/mongodb/mongodb-indexing.html)

```conf
; db.collection.createIndex( keys, options )
; 创建 age 字段的升序索引
db.mycollection.createIndex({ age: 1 })
; 创建 name 字段的文本索引
db.mycollection.createIndex({ name: "text" })
; 创建唯一索引
db.mycollection.createIndex( { field: 1 }, { unique: true } )
; 创建后台运行的索引
db.mycollection.createIndex( { field: 1 }, { background: true } )
; 创建稀疏索引
db.mycollection.createIndex( { field: 1 }, { sparse: true } )
; 创建文本索引并指定权重
db.mycollection.createIndex( { field: "text" }, { weights: { field: 10 } } )
; 创建地理空间索引
; 对于存储地理位置数据的字段，可以使用 2dsphere 或 2d 索引类型来创建地理空间索引。
; 2dsphere 索引，适用于球形地理数据
db.mycollection.createIndex( { location: "2dsphere" } )
; 2d 索引，适用于平面地理数据
db.mycollection.createIndex( { location: "2d" } )
; 可以使用哈希索引对字段进行哈希，以支持大范围的数值查找
db.collection.createIndex( { field: "hashed" } )

db.collection.getIndexes()
; 删除指定的索引
db.collection.dropIndex( "indexName" )
; 删除所有索引
db.collection.dropIndexes()
```
