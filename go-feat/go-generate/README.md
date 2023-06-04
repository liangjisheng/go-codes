# go generate

[post](http://c.biancheng.net/view/4442.html)
[error-code](https://darjun.github.io/2019/08/21/golang-generate/)

go generate 命令是在Go语言 1.4 版本里面新添加的一个命令，当运行该命令时，它将扫描与当前包相关的源代码文件，找出所有包含//go:generate的特殊注释，提取并执行该特殊注释后面的命令

如果有多个目录且有多个 go generate 的话可以在最上层目录执行 go generate ./...
