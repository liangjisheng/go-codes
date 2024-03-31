package main

// replace除了可以将远程的包进行替换外 还可以将本地存在的modules替换成任意指定的名字
// my/example/pkg显然是个不存在的包 我们将用本地目录的pkg包替换它
// 本地替换的价值在于它提供了一种使自动生成的代码进入go modules系统的途径
// 毕竟不管是go tools还是rpc工具，这些自动生成代码也是项目的一部分
// 如果不能纳入包管理器的管理范围想必会带来很大的麻烦
// replace唯一的限制是它只能处理顶层依赖

import "my/example/pkg"

func main() {
	pkg.Hello()
}
