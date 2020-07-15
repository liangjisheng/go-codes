#!/bin/bash

# 创建索引库
curl -X PUT http://117.51.148.112:9200/jw/
# {"acknowledged":true,"shards_acknowledged":true,"index":"jw"}

# 创建document
curl -X POST http://117.51.148.112:9200/jw/employee -d '{ "first_name":"bin", "last_name":"tang", "age":33, "about":"I love to go rock climbing", "interests":[ "sports", "music" ] }' -H "content-type: application/json"
# {"_index":"jw","_type":"employee","_id":"8M2vZXABL-v0t_yN2Cp5","_version":1,"result":"created","_shards":{"total":2,"successful":1,"failed":0},"_seq_no":0,"_primary_term":1}

# 对document扩展 比如新增爱好 年龄等属性
curl -X POST http://117.51.148.112:9200/jw/employee -d '{ "first_name":"bin", "last_name":"tang", "age":33, "about":"I love to go rock climbing", "interests":[ "sports", "music" ], "hobby":"play basketball" }' -H "content-type: application/json"
# {"_index":"jw","_type":"employee","_id":"8s2zZXABL-v0t_yNMirh","_version":1,"result":"created","_shards":{"total":2,"successful":1,"failed":0},"_seq_no":1,"_primary_term":1}

# 更新document
curl -X PUT http://117.51.148.112:9200/jw/employee/1 -d '{"first_name":"bin", "last_name":"pang", "age":30, "about":"I love to go rock climbing", "interests":["sports", "music"]}' -H "content-type: application/json"
# {"_index":"jw","_type":"employee","_id":"1","_version":1,"result":"created","_shards":{"total":2,"successful":1,"failed":0},"_seq_no":2,"_primary_term":1}

# 根据document的id来获取数据
curl -X GET http://117.51.148.112:9200/jw/employee/1?pretty

# 根据field来查询数据
curl -X GET http://117.51.148.112:9200/jw/employee/_search?q=first_name="bin"

# 根据field来查询数据: match
curl -X GET http://117.51.148.112:9200/jw/employee/_search?pretty -d '{ "query": {"match": {"first_name": "bin"}} }' -H "content-type: application/json"

# 对多个field发起查询：multi_match
curl -X GET http://117.51.148.112:9200/jw/employee/_search?pretty -d '{ "query":{"multi_match":{"query":"bin","fields":["last_name","first_name"],"operator":"and"}}}' -H "content-type: application/json"

# 多个term对多个field发起查询:bool（boolean）
# 组合查询: must,must_not,should
# must + must : 交集
# must +must_not ：差集
# should+should : 并集
curl -X GET http://117.51.148.112:9200/jw/employee/_search?pretty -d '{"query":{"bool" :{"must" : [{"match":{"first_name":"bin"} },{"match":{"age":33} }]}}}' -H "content-type: application/json"
