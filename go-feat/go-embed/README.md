# go embed

[embed](https://zhuanlan.zhihu.com/p/351931501)
[embed](https://www.cnblogs.com/apocelipes/p/13907858.html)

embed是在Go 1.16中新加包。它通过//go:embed指令，可以在编译阶段将静态资源文件打包进编译好的程序中，并提供访问这些文件的能力。

在embed中，可以将静态资源文件嵌入到三种类型的变量，分别为：字符串、字节数组、embed.FS文件类型
embed.FS自身是只读的，所以我们不能在运行时添加或删除嵌入的文件，fs.File也是只读的，所以我们不能修改嵌入资源的内容

注意事项

go:embed指令只能用在包一级的变量中，不能用在函数或方法级别
