package utils

import (
	"fmt"
	"testing"

	"github.com/thoas/go-funk"
)

func TestDemo(t *testing.T) {
	m1 := map[int64]int64{
		1: 1,
		2: 2,
	}

	m2 := map[int64]int64{
		3: 3,
		4: 4,
	}

	m3 := m1
	for key, value := range m2 {
		m3[key] += value
	}

	t.Log(m3)
}

// 获取map中所有的key
func TestMapKeys(t *testing.T) {
	res := map[string]int{
		"张三": 18,
		"李四": 20,
		"赵武": 25,
	}
	keys := funk.Keys(res)
	fmt.Printf("keys: %#v %T \n", keys, keys)
}

// 获取map中的所有values
func TestMapValues(t *testing.T) {
	res := map[string]int{
		"张三": 18,
		"李四": 20,
		"赵武": 25,
	}
	values := funk.Values(res)
	fmt.Printf("values: %#v %T \n", values, values)
}

type User struct {
	Name string
	Home struct {
		City string
	}
}

// 取结构体某元素为切片
func TestGetToSlice(t *testing.T) {
	userList := []User{
		{
			Name: "张三",
			Home: struct {
				City string
			}{"北京"},
		},
		{
			Name: "小明",
			Home: struct {
				City string
			}{"南京"},
		},
	}
	// 取一层
	names := funk.Get(userList, "Name")
	fmt.Println("names:", names)
	// 取其他层
	homes := funk.Get(userList, "Home.City")
	fmt.Println("homes:", homes)
}

func TestIsEqual(t *testing.T) {
	// 对比字符串
	fmt.Println("对比字符串:", funk.IsEqual("a", "a"))
	// 对比int
	fmt.Println("对比int:", funk.IsEqual(1, 1))
	// 对比float64
	fmt.Println("对比float64:", funk.IsEqual(float64(1), float64(1)))
	// 对比结构体
	stu1 := Student{Name: "张三", Age: 18}
	stu2 := Student{Name: "张三", Age: 18}
	fmt.Println("对比结构体:", funk.IsEqual(stu1, stu2))
}

// 判断类型是否一样
func TestIsType(t *testing.T) {
	var a, b int8 = 1, 2
	fmt.Println("A: ", funk.IsType(a, b))
	c := 3
	d := "3"
	fmt.Println("B:", funk.IsType(c, d))
}

// 判断是否是array|slice
func TestCollect(t *testing.T) {
	a := []int{1, 2, 3}
	b := "str"
	c := []Student{
		{Name: "张三", Age: 18},
		{Name: "李四", Age: 18},
	}
	d := [2]string{"c", "go"}

	fmt.Println("a:", funk.IsCollection(a))
	fmt.Println("b:", funk.IsCollection(b))
	fmt.Println("c:", funk.IsCollection(c))
	fmt.Println("d:", funk.IsCollection(d))
}

// funk.IsEmpty(obj interface{}): 判断为空
// funk.NotEmpty(obj interface{}): 判断不为空
// 判断是否为空
func TestIsEmpty(t *testing.T) {
	// 空结构体
	fmt.Println("空结构体", funk.IsEmpty([]int{}))
	// 空字符串
	fmt.Println("空字符串:", funk.IsEmpty(""))
	// 判断数字0
	fmt.Println("0:", funk.IsEmpty(0))
	// 判断字符串'0'
	fmt.Println("'0':", funk.IsEmpty("0"))
	// nil
	fmt.Println("nil:", funk.IsEmpty(nil))
}

// 任意数字转float64
// 将任何数字类型，转成float64类型，@注:只能是数字类型: uint8、uint16、uint32、uint64、int、int8、int16、int32、int64、float32、float64
func TestToFloat64(t *testing.T) {
	// int to float64
	d1, _ := funk.ToFloat64(10)
	fmt.Printf("d1 = %v %T \n", d1, d1)
	//@会失败
	d2, err := funk.ToFloat64("10")
	fmt.Printf("d2 = %v %v \n", d2, err)
}

// 将X转成[]X
// 返回任一类型的类型切片
func TestSliceOf(t *testing.T) {
	a := []int{10, 20, 30}
	// []int 转成 [][]int
	fmt.Printf("%v %T \n", funk.SliceOf(a), funk.SliceOf(a))
	// string 转成 []string
	fmt.Printf("%v %T \n", funk.SliceOf("go"), funk.SliceOf("go"))
}

// 根据字符串生成切片
// func Shard(str string, width int, depth int, restOnly bool) []string
// width: 代表根据几个字节生成一个元素。
// depth: 将字符串前x个元素转成切片。
// restOnly: 当为false时，最后一个元素为原字符串,当为true时,最后一个元素为原字符串剩余元素
func TestShard(t *testing.T) {
	tokey := "Hello,Word"
	shard := funk.Shard(tokey, 1, 5, false)
	shard1 := funk.Shard(tokey, 1, 5, true)
	fmt.Println("shard: ", shard)
	fmt.Println("shard1: ", shard1)

	shard2 := funk.Shard(tokey, 2, 5, false)
	shard22 := funk.Shard(tokey, 2, 5, true)
	fmt.Println("shard2: ", shard2)
	fmt.Println("shard22: ", shard22)
}

func TestMax(t *testing.T) {
	// 求最大int
	fmt.Println("MaxInt:", funk.MaxInt([]int{30, 10, 8, 11}))
	// 求最大浮点数
	fmt.Println("MaxFloat64:", funk.MaxFloat64([]float64{10.2, 11.0, 8.03}))
	// 求最大字符串
	fmt.Println("MaxString:", funk.MaxString([]string{"a", "d", "c", "b"}))
}

func TestMin(t *testing.T) {
	// 求最小int
	fmt.Println("MinInt:", funk.MinInt([]int{30, 10, 8, 11}))
	// 求最小浮点数
	fmt.Println("MinFloat64:", funk.MinFloat64([]float64{10.2, 11.0, 8.03}))
	// 求最小字符串
	fmt.Println("MinString:", funk.MinString([]string{"a", "d", "c", "b"}))
}

func TestSum(t *testing.T) {
	// 整型
	a := []int{5, 10, 15, 20}
	fmt.Println("int sum:", funk.Sum(a))
	// 浮点型
	b := []float64{5.11, 2.23, 3.31, 0.32}
	fmt.Println("float64 sum:", funk.Sum(b))
}

func TestProduct(t *testing.T) {
	// 整型
	a := []int{2, 3, 4, 5}
	fmt.Println("int Product:", funk.Product(a))
	// 浮点型
	b := []float64{1.1, 1.2, 1.3, 1.4}
	fmt.Println("float64 Product:", funk.Product(b))
}

func TestRandom(t *testing.T) {
	for i := 1; i <= 10; i++ {
		// 生成任意数字类型
		fmt.Println(funk.RandomInt(1, 100))
	}
}

// 生成随机字符串
func TestRandomString(t *testing.T) {
	for i := 1; i <= 3; i++ {
		// 从默认字符串生成
		fmt.Println("从默认字符串生成:", funk.RandomString(i))
		// 从指的字符串生成
		fmt.Println("从指定字符串生成:", funk.RandomString(i, []rune{'您', '好', '北', '京'}))
	}
}

// 三元运算
func TestShortIf(t *testing.T) {
	fmt.Println("10 > 5 :", funk.ShortIf(10 > 5, 10, 5))
	fmt.Println("10.0 == 10 : ", funk.ShortIf(10.0 == 10, "yes", "no"))
	fmt.Println("'a' == 'b' : ", funk.ShortIf('a' == 'b', "equal chars", "unequal chars"))
}
