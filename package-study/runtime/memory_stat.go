package runtime_demo

import (
	"fmt"
	"runtime"
	"time"
	"unsafe"
)

var memStats runtime.MemStats

func MemoryStat1() {
	data := make([]byte, 10)
	fmt.Printf("Allocated %d bytes\n", len(data))

	// 打印内存分配信息
	runtime.ReadMemStats(&memStats)
	fmt.Printf("Total allocated: %d bytes\n", memStats.TotalAlloc)
	fmt.Printf("Num alloc: %d\n", memStats.Alloc)
	fmt.Printf("Num sys: %d\n", memStats.Sys)
}

//Go 语言的垃圾收集器主要包括以下几个阶段：
//标记: 标记所有可达的对象。
//扫描: 回收未被标记的对象。
//复制: 将存活的对象复制到新的内存区域

func MemoryStat2() {
	// 设置垃圾收集间隔
	runtime.GC() // 手动触发一次 GC

	// 分配多个对象
	for i := 0; i < 1000; i++ {
		_ = make([]byte, 10)
	}

	// 触发垃圾收集
	runtime.GC()

	// 打印内存统计信息
	runtime.ReadMemStats(&memStats)
	fmt.Printf("Total allocated: %d bytes\n", memStats.TotalAlloc)
	fmt.Printf("Total sys: %d bytes\n", memStats.Sys)
	fmt.Printf("Num alloc: %d\n", memStats.Alloc)
	fmt.Printf("Num GC: %d\n", memStats.NumGC)

	// 等待一段时间，让 GC 运行
	time.Sleep(2 * time.Second)
}

func StackMemory() {
	var a [1000]int // 分配一个较大的数组
	fmt.Println("Array allocated on stack")
	fmt.Printf("Size of array: %d bytes\n", unsafe.Sizeof(a))
	runtime.ReadMemStats(&memStats)
	fmt.Printf("Total allocated: %d bytes\n", memStats.Alloc)
}

func PointMemory() {
	// 栈指针示例
	var a int = 10
	fmt.Printf("Stack pointer address: %p\n", &a)
	fmt.Printf("Size of int: %d bytes\n", unsafe.Sizeof(a))

	// 堆指针示例
	b := new(int)
	*b = 20
	fmt.Printf("Heap pointer address: %p\n", b)
	fmt.Printf("Size of int: %d bytes\n", unsafe.Sizeof(*b))

	// 打印内存统计信息
	runtime.ReadMemStats(&memStats)
	fmt.Printf("Total allocated: %d bytes\n", memStats.TotalAlloc)
	fmt.Printf("Total sys: %d bytes\n", memStats.Sys)
	fmt.Printf("Num alloc: %d\n", memStats.Alloc)
}
