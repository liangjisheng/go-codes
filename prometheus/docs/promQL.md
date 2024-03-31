# PromQL

|       函数       |                           含义                           |
|:--------------:|:------------------------------------------------------:|
|     sum()      |                         对样本值求和                         |
|     min()      |                       求取样本值中的最小者                       |
|     max()      |                      	求取样本值中的最大者                       |
|     avg()      |                       	对样本值求平均值                        |
|    count()     |                    	对分组内的时间序列进行数量统计                    |
|    stddev()    |           	对样本值求标准差，以帮助用户了解数据的波动大小（或称之为波动程度）           |
|    stdvar()    |                	对样本值求方差，它是求取标准差过程中的中间状态                |
|     topk()     |        	逆序返回分组内的样本值最大的前 k 个时间序列及其值，即最大的 k 个样本值         |
|   bottomk()    |        	顺序返回分组内的样本值最小的前 k 个时间序列及其值，即最小的 k 个样本值         |
|   quantile()   | 	分位数，用于评估数据的分布状态，该函数会返回分组内指定的分位数的值，即数值落在小于等于指定的分位区间的比例 |
| count_values() |            	对分组内的时间序列的样本值进行数量统计，即等于某值的样本个数             |

```shell
# 有没有 {} 的效果是一样的
prometheus_http_requests_total
prometheus_http_requests_total{}
prometheus_http_requests_total{code="200", handler="/api/v1/query_range"}
prometheus_http_requests_total{code="200", handler="/api/v1/query_range"}[5m]
rate(prometheus_http_requests_total{code="200", handler="/api/v1/query_range"}[5m])
# 查询系统所有 http 请求的总量, 样本求和
sum(prometheus_http_requests_total)
# 样本个数
count(prometheus_http_requests_total)
# 样本分组的个数
count_values("count", prometheus_http_requests_total)
# 使用内置的__name__标签来指定监控指标名称
# ~ 表示正则匹配
{__name__="prometheus_http_requests_total"}
{__name__ =~"prometheus_http_requests_total"}
# 最多和最少的 5 个指标
topk(5, prometheus_http_requests_total)
bottomk(5, prometheus_http_requests_total)
# 计算当前样本数据值的分布情况 当φ为0.5时，即表示找到当前样本数据中的中位数
quantile(0.5, prometheus_http_requests_total)

# summary
prometheus_tsdb_wal_fsync_duration_seconds
# histogram
# QPS = query per second
# 一般用counter定义指标收集，但是也可以直接对 Histogram 收集的count进行查询处理
rate(prometheus_http_request_duration_seconds_count{handler="/api/v1/format_query"}[1m])

# 查询指标
prometheus_http_request_duration_seconds_bucket{handler="/api/v1/format_query"}
# 从最后一分钟的指标中获取一系列值
prometheus_http_request_duration_seconds_bucket{handler="/api/v1/format_query"}[1m]
# 查看每个 bucket 的每秒变化率来了解这些 bucket 是如何随时间变化的
rate(prometheus_http_request_duration_seconds_bucket{handler="/api/v1/format_query"}[1m])
# 计算出哪个 bucket 标签包含给定的分位数（例如第 95 个百分位数）
histogram_quantile(0.95, rate(prometheus_http_request_duration_seconds_bucket{handler="/api/v1/format_query"}[1m]))
# 数据量少的话使用这个,上一条会导致数据为 0
histogram_quantile(0.95, prometheus_http_request_duration_seconds_bucket{handler="/api/v1/format_query"})
ceil(increase(prometheus_http_request_duration_seconds_bucket{le!="+Inf"}[1m]))

prometheus_http_request_duration_seconds_bucket{handler="/static/*filepath"}
prometheus_http_request_duration_seconds_bucket{handler="/static/*filepath", le="0.1"}
sum(prometheus_http_request_duration_seconds_bucket{handler="/static/*filepath", le="0.1"})
prometheus_http_request_duration_seconds_bucket{handler="/static/*filepath", le="0.4"}
sum(prometheus_http_request_duration_seconds_bucket{handler="/static/*filepath", le="0.4"}) - sum(prometheus_http_request_duration_seconds_bucket{handler="/static/*filepath", le="0.1"})

# 基于2小时的样本数据，来预测主机可用磁盘空间的是否在4个小时候被占满，可以使用如下表达式
predict_linear(node_filesystem_free_bytes{mountpoint="/data"}[2h], 4 * 3600) < 0

#summary
prometheus_tsdb_wal_fsync_duration_seconds
prometheus_tsdb_wal_fsync_duration_seconds{quantile="0.5"}
prometheus_tsdb_wal_fsync_duration_seconds{quantile="0.9"}
prometheus_tsdb_wal_fsync_duration_seconds{quantile="0.99"}
prometheus_tsdb_wal_fsync_duration_seconds_sum
prometheus_tsdb_wal_fsync_duration_seconds_count

prometheus_engine_query_duration_seconds

go_gc_duration_seconds{instance="localhost:9090"}
go_gc_duration_seconds_count{instance="localhost:9090"}
go_gc_duration_seconds_sum{instance="localhost:9090"}
#查询所有 instance 是 localhost 开头的指标
go_gc_duration_seconds_count{instance=~"localhost.*"}

go_gc_duration_seconds_count{instance="localhost:9090"}[5m]
#查询一天前当前 5 分钟前的时序数据集
go_gc_duration_seconds_count{instance="localhost:9090"}[5m] offset 1d

#每台主机 CPU 在最近 5 分钟内的平均使用率
node_cpu_seconds_total
node_cpu_seconds_total{mode="idle"}
node_cpu_seconds_total{mode="idle"}[5m]
rate(node_cpu_seconds_total{mode="idle"}[5m])
1-avg(rate(node_cpu_seconds_total{mode="idle"}[5m]))
(1-avg(rate(node_cpu_seconds_total{mode="idle"}[5m])) by (instance)) * 100

#查询 1 分钟的 load average 的时间序列是否超过主机 CPU 数量 2 倍
#计算每台主机 cpu 个数
count (node_cpu_seconds_total{mode="idle"}) by (instance)
#没有值说明没有超负荷运行
node_load1 > on(instance) 2 * count (node_cpu_seconds_total{mode="idle"}) by (instance)

# 机器内存使用率
#可用内存空间：空闲内存、buffer、cache 指标之和
node_memory_MemFree_bytes + node_memory_Buffers_bytes + node_memory_Cached_bytes
#已用内存空间：总内存空间减去可用空间
node_memory_MemTotal_bytes - (node_memory_MemFree_bytes + node_memory_Buffers_bytes + node_memory_Cached_bytes)
#使用率：已用空间除以总空间
(node_memory_MemTotal_bytes - (node_memory_MemFree_bytes + node_memory_Buffers_bytes + node_memory_Cached_bytes)) / node_memory_MemTotal_bytes * 100

(node_memory_MemTotal_bytes - node_memory_MemFree_bytes) / node_memory_MemTotal_bytes
(node_memory_MemTotal_bytes - node_memory_MemFree_bytes) / node_memory_MemTotal_bytes > 0.95
```
