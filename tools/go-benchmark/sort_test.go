package sort

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestInsertionSort(t *testing.T) {
	a := stuff(10)
	InsertionSort(a)
	t.Log(a)
}

func TestQuickSort(t *testing.T) {
	a := stuff(10)
	QuickSort(a)
	t.Log(a)
}

func BenchmarkInsertionSort10(b *testing.B) {
	benchmarkInsertionSort(b, 10)
}

func BenchmarkQuickSort10(b *testing.B) {
	benchmarkQuickSort(b, 10)
}

func BenchmarkInsertionSort100(b *testing.B) {
	benchmarkInsertionSort(b, 100)
}

func BenchmarkQuickSort100(b *testing.B) {
	benchmarkQuickSort(b, 100)
}

func BenchmarkInsertionSort1000(b *testing.B) {
	benchmarkInsertionSort(b, 1000)
}

func BenchmarkQuickSort1000(b *testing.B) {
	benchmarkQuickSort(b, 1000)
}

func BenchmarkInsertionSort10000(b *testing.B) {
	benchmarkInsertionSort(b, 10000)
}

func BenchmarkQuickSort10000(b *testing.B) {
	benchmarkQuickSort(b, 10000)
}

func BenchmarkInsertionSort100000(b *testing.B) {
	benchmarkInsertionSort(b, 100000)
}

func BenchmarkQuickSort100000(b *testing.B) {
	benchmarkQuickSort(b, 100000)
}

func BenchmarkInsertionSort1000000(b *testing.B) {
	benchmarkInsertionSort(b, 1000000)
}

func BenchmarkQuickSort1000000(b *testing.B) {
	benchmarkQuickSort(b, 1000000)
}

func stuff(count int) (a []int) {
	seed := time.Now().UTC().UnixNano()
	rng := rand.New(rand.NewSource(seed))
	a = make([]int, count, count)
	for i := 0; i < count; i++ {
		a[i] = rng.Intn(100000)
	}
	return
}

func benchmarkInsertionSort(b *testing.B, count int) {
	for i := 0; i < b.N; i++ {
		a := stuff(count)
		InsertionSort(a)
	}
}

func benchmarkQuickSort(b *testing.B, count int) {
	for i := 0; i < b.N; i++ {
		a := stuff(count)
		QuickSort(a)
	}
}

// Example函数没有参数，也没有返回值，主要作用是文档,go doc服务器将Example函数
// 嵌入到被测试函数的文档里
// 第二个作用是验证输出结果
func ExampleInsertionSort() {
	a := []int{3, 2, 4, 1, 5}
	QuickSort(a)
	fmt.Println(a)
	// Output:
	// [1 2 3 4 5]
}

func ExampleQuickSort() {
	a := []int{3, 2, 4, 1, 5}
	QuickSort(a)
	fmt.Println(a)
	// Output:
	// [1 2 3 4]
}
