# README

[blog](https://mp.weixin.qq.com/s?__biz=MzIzMDU0MTA3Nw==&mid=2247483947&idx=1&sn=0e91d102be608ef78793e118b942c548&chksm=e8b09051dfc719471764a541a48285b699f25427017dd98113359a75e0609224d3b4082c53cb&scene=178&cur_album_id=1504658477523206145#rd)

使用go tool pprof命令获取指定的profile文件，采集60s的CPU使用情况
会将采集的数据下载到本地，之后进入交互模式，可以使用命令行查看运行信息
go tool pprof http://127.0.0.1:6060/debug/pprof/profile -seconds 60

使用命令行进入交互式模式查看
go tool pprof pprof.samples.cpu.001.pb.gz
也可以打开浏览器查看 cpu 使用火焰图
go tool pprof -http=:8081 ~/pprof/pprof.samples.cpu.001.pb.gz

如果报没有找到 dot 错误, 则安装 graphviz
brew install graphviz
