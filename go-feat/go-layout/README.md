# go layout

[marmotedu/go-standards](https://github.com/marmotedu/go-standards)
[go-engineering](https://www.yipwinghong.com/2021/08/04/Go_engineering-standard/#/cmd)
[go-engineering](https://www.yipwinghong.com/2021/12/10/Go_engineering-specification-design)
[go layout](https://github.com/marmotedu/iam)

概述 go 项目布局

/cmd 项目主干, 不要在这个目录中放置太多代码。如果你认为代码可以导入并在其他项目中使用，那么它应该位于 /pkg 目录中。
如果代码不是可重用的，或者你不希望其他人重用它，请将该代码放到 /internal 目录中
通常有一个小的 main 函数，从 /internal 和 /pkg 目录导入和调用代码，除此之外没有别的东西

/internal 私有应用程序和库代码。这是你不希望其他人在其应用程序或库中导入代码。请注意，这个布局模式是由 Go 编译器本身执行的。

/pkg 外部应用程序可以使用的库代码

/api OpenAPI/Swagger 规范，JSON 模式文件，协议定义文件

/web 特定于 Web 应用程序的组件:静态 Web 资产、服务器端模板和 SPAs

/config 配置文件模板或者默认配置

/init System init（systemd，upstart，sysv）和 process manager/supervisor（runit，supervisor）配置

/scripts 执行各种构建、安装、分析等操作的脚本 保持了根级别的 Makefile 变得小而简单

/build 打包和持续集成 CI (travis、circle、drone)配置

/deploy IaaS、PaaS、系统和容器编排部署配置和模板(docker-compose、kubernetes/helm、mesos、terraform、bosh)

/test 外部测试应用程序和测试数据 可以使用 /test/data

/docs 文档

/tools 这个项目的支持工具

/examples 应用程序和/或公共库的示例

/third_party 外部辅助工具，分叉代码和其他第三方工具(例如 Swagger UI)。

/githooks Git hooks

/assets 与存储库一起使用的其他资产(图像、徽标等)

/website 如果你不使用 Github 页面，则在这里放置项目的网站数据
