# windows install

下载,解压

[redis](https://github.com/microsoftarchive/redis/releases/download/win-3.2.100/Redis-x64-3.2.100.zip)

配置成服务
./redis-server.exe --service-install redis.windows-service.conf --loglevel verbose

redis目录添加到环境变量

连接redis-server
redis-cli.exe
