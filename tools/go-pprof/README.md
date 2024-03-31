# README

[blog](https://mp.weixin.qq.com/s/pLp7DB8tBVrTgheZrPg5dg)

如果报没有找到 dot 错误, 则安装 graphviz

```shell
brew install graphviz
```

可以通过代码生成性能数据文件

```go
package main

import (
	"os"
	"runtime/pprof"
)

func main() {
	cpuOut, _ := os.Create("cpu.profile")
	defer cpuOut.Close()
	pprof.StartCPUProfile(cpuOut)
	defer pprof.StopCPUProfile()

	memOut, _ := os.Create("mem.profile")
	defer memOut.Close()
	defer pprof.WriteHeapProfile(memOut)

	Sum(3, 5)
}

func Sum(a, b int) int {
	return a + b
}
```
