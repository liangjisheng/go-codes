package gotest

import (
	"fmt"
	"os"
	"testing"
)

// 如果在同一个测试文件中，每一个测试用例运行前后的逻辑是相同的
// 一般会写在 setup 和 teardown 函数中。例如执行前需要实例化待
// 测试的对象，如果这个对象比较复杂，很适合将这一部分逻辑提取出来
// 执行后可能会做一些资源回收类的工作 例如关闭网络连接 释放文件等

func setup() {
	fmt.Println("Before all tests")
}

func teardown() {
	fmt.Println("After all tests")
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}
