#!/bin/bash

# https://prometheus.io/
# https://blog.csdn.net/csolo/article/details/82460539
# https://www.cnblogs.com/zqj-blog/p/10981350.html

# 安装 node_exporter
wget https://github.com/prometheus/node_exporter/releases/download/v0.18.1/node_exporter-0.18.1.linux-amd64.tar.gz
tar zxf node_exporter-0.18.1.linux-amd64.tar.gz
cd node_exporter-0.18.1.linux-amd64

# 修改端口启动 node_exporter
nohup /usr/local/prometheus/node_exporter-0.18.1.linux-amd64/node_exporter --web.listen-address=":9101" >nohup.out 2>&1 &

# 查看 ubuntu server 里面的 metrics
# curl http://localhost:9101/metrics


# 安装 prometheus
wget https://github.com/prometheus/prometheus/releases/download/v2.42.0/prometheus-2.42.0.linux-amd64.tar.gz
tar zxf prometheus-2.42.0.linux-amd64.tar.gz
cd prometheus-2.42.0.linux-amd64
nohup /usr/local/prometheus/prometheus-2.42.0.linux-amd64/prometheus >nohup.out 2>&1 &

# 查看状态
# http://117.51.148.112:9090/status


# 安装 grafana
# https://grafana.com/grafana/download
wget https://dl.grafana.com/enterprise/release/grafana-enterprise-9.4.3.linux-amd64.tar.gz
tar -zxvf grafana-enterprise-9.4.3.linux-amd64.tar.gz
cd grafana-9.4.3
nohup /usr/local/prometheus/grafana-9.4.3/bin/grafana-server >nohup.out 2>&1 &

# http://117.51.148.112:3000/


# 安装 alertmanager
wget https://github.com/prometheus/alertmanager/releases/download/v0.20.0/alertmanager-0.20.0.linux-amd64.tar.gz
tar zxf alertmanager-0.20.0.linux-amd64.tar.gz
cd alertmanager-0.20.0.linux-amd64
nohup /usr/local/prometheus/alertmanager-0.20.0.linux-amd64/alertmanager >nohup.out 2>&1 &

# http://117.51.148.112:9093/#/alerts
