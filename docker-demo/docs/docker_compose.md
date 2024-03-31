# docker compose

Docker Compose 是 Docker 官方的开源项目，能够实现对 Docker 容器集群的快速编排，
以保证快速部署分布式应用。

在实际的生产环境中，往往需要多个容器配合工作，才能称的上一个完整的应用服务。

举个栗子，你想搭建一个 Web 服务，除了 Web 服务容器本身，若涉及数据存储，就需要依赖数据库
容器（Mysql、Mongdb 等）；若想实现分词搜索，就需要依赖搜索引擎容器（Elasticsearch、Solor 等）
其他诸如缓存容器、负载均衡容器等。

Compose 是用于定义和运行多容器 Docker 应用程序的工具。通过 Compose，您可以使用
YML 文件来配置应用程序需要的所有服务。然后，使用一个命令，就可以从 YML 文件配置中
创建并启动所有服务。

Compose 中有两个重要的概念：
服务 (service)：一个应用的容器，实际上可以包括若干运行相同镜像的容器实例。
项目 (project)：由一组关联的应用容器组成的一个完整业务单元，在 docker-compose.yml 文件中定义

docker-compose --help 查看可用命令
也可以使用 docker 下的子命令 compose 查看， docker compose --help

```text
Usage:  docker compose [OPTIONS] COMMAND

Docker Compose

Options:
      --ansi string                Control when to print ANSI control
                                   characters
                                   ("never"|"always"|"auto") (default
                                   "auto")
      --compatibility              Run compose in backward
                                   compatibility mode
      --env-file string            Specify an alternate environment file.
  -f, --file stringArray           Compose configuration files
      --profile stringArray        Specify a profile to enable
      --project-directory string   Specify an alternate working
                                   directory
                                   (default: the path of the, first
                                   specified, Compose file)
  -p, --project-name string        Project name

Commands:
  build       Build or rebuild services
  convert     Converts the compose file to platform's canonical format
  cp          Copy files/folders between a service container and the local filesystem
  create      Creates containers for a service.
  down        Stop and remove containers, networks
  events      Receive real time events from containers.
  exec        Execute a command in a running container.
  images      List images used by the created containers
  kill        Force stop service containers.
  logs        View output from containers
  ls          List running compose projects
  pause       Pause services
  port        Print the public port for a port binding.
  ps          List containers
  pull        Pull service images
  push        Push service images
  restart     Restart service containers
  rm          Removes stopped service containers
  run         Run a one-off command on a service.
  start       Start services
  stop        Stop services
  top         Display the running processes
  unpause     Unpause services
  up          Create and start containers
  version     Show the Docker Compose version information
```