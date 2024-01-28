package utils

import (
	"fmt"
	"testing"

	"github.com/thoas/go-funk"
)

func TestExist(t *testing.T) {
	// 判断任意类型
	fmt.Println("str->", funk.Contains([]string{"a", "b"}, "a"))
	// int 类型
	fmt.Println("int->", funk.ContainsInt([]int{1, 2}, 1))
}

func TestIndexOf(t *testing.T) {
	strArr := []string{"go", "java", "c", "java"}
	// 具体类型
	fmt.Println("c: ", funk.IndexOfString(strArr, "c"))
	// 验证第一次出现位置
	fmt.Println("java: ", funk.IndexOfString(strArr, "java"))
	// 任意类型
	fmt.Println("go: ", funk.IndexOf(strArr, "go"))
	// 不存在时返回-1
	fmt.Println("php: ", funk.IndexOfString(strArr, "php"))
}

func TestLastOf(t *testing.T) {
	strArr := []string{"go", "java", "c", "java"}
	// 具体类型
	fmt.Println("c: ", funk.LastIndexOfString(strArr, "c"))
	// 验证第一次出现位置
	fmt.Println("java: ", funk.LastIndexOfString(strArr, "java"))
	// 任意类型
	fmt.Println("go: ", funk.LastIndexOf(strArr, "go"))
	// 不存在时返回-1
	fmt.Println("php: ", funk.LastIndexOf(strArr, "php"))
}

func TestEvery(t *testing.T) {
	//当 elements 都在 in 中时，则返回 true; 否则为 false
	strArr := []string{"go", "java", "c", "python"}
	fmt.Println("都存在:", funk.Every(strArr, "go", "c"))
	fmt.Println("有一个不存在:", funk.Every(strArr, "php", "go"))
	fmt.Println("都不存在:", funk.Every(strArr, "php", "c++"))
}

func TestSome(t *testing.T) {
	//当elements至少有一个在in中时，则返回true; 否则为false
	strArr := []string{"go", "java", "c", "python"}
	fmt.Println("都存在:", funk.Some(strArr, "go", "c"))
	fmt.Println("至少一个存在:", funk.Some(strArr, "php", "go"))
	fmt.Println("都不存在:", funk.Some(strArr, "php", "c++"))
}

func TestLastOrFirst(t *testing.T) {
	number := []int{10, 30, 12, 23}
	// 获取第一个元素
	fmt.Println("Head: ", funk.Head(number))
	// 获取最后一个元素
	fmt.Println("Last: ", funk.Last(number))
}

type Student struct {
	Name string
	Age  int
}

func TestFill(t *testing.T) {
	//将in中的所有元素，设置成fillValue
	var data = make([]int, 3)
	fill, _ := funk.Fill(data, 100)
	fmt.Printf("fill: %v \n", fill)

	// 将所有值设置成2
	input := []int{1, 2, 3}
	result, _ := funk.Fill(input, 2)
	fmt.Printf("result: %v \n", result)

	var structData = make([]Student, 2)
	stuInfo, _ := funk.Fill(structData, Student{Name: "张三", Age: 18})
	fmt.Printf("stuInfo: %v \n", stuInfo)
}

//Join(larr, rarr interface{}, fnc JoinFnc):  当fnc=funk.InnerJoin,代表合并两个任意类型切片
//JoinXXX(larr, rarr interface{}, fnc JoinFnc): 指定类型合并，推荐使用

type cus struct {
	Name string
	Age  int
	Home string
}

// 取两个切片共同元素结果集
func TestJoin(t *testing.T) {
	a := []int64{1, 3, 5, 7}
	b := []int64{3, 7}
	// 任意类型切片交集
	join := funk.Join(a, b, funk.InnerJoin)
	fmt.Println("join = ", join)

	// 指定类型取交集
	joinInt64 := funk.JoinInt64(a, b, funk.InnerJoinInt64)
	fmt.Println("joinInt64 = ", joinInt64)

	// 自定义结构体交集
	sliceA := []cus{
		{"张三", 20, "北京"},
		{"李四", 22, "南京"},
	}
	sliceB := []cus{
		{"张三", 20, "北京"},
		{"李四", 22, "上海"},
	}
	res := funk.Join(sliceA, sliceB, funk.InnerJoin)
	fmt.Println("自定义结构体: ", res)
}

//获取去掉两切片共同元素结果集
//同样使用Join和JoinXXX方法，而fnc设置成funk.OuterJoin

func TestDiffSlice(t *testing.T) {
	a := []int64{1, 3, 5, 7}
	b := []int64{3, 7, 10}
	// 任意类型切片交集
	join := funk.Join(a, b, funk.OuterJoin)
	fmt.Println("OuterJoin = ", join)

	// 指定类型取交集
	joinInt64 := funk.JoinInt64(a, b, funk.OuterJoinInt64)
	fmt.Println("joinInt64 = ", joinInt64)

	// 自定义结构体交集
	sliceA := []cus{
		{"张三", 20, "北京"},
		{"李四", 22, "南京"},
	}
	sliceB := []cus{
		{"张三", 20, "北京"},
		{"李四", 22, "上海"},
	}
	res := funk.Join(sliceA, sliceB, funk.OuterJoin)
	fmt.Println("自定义结构体: ", res)
}

// 求只存在某切片的元素(除去共同元素)
func TestLeftAndRightJoin(t *testing.T) {
	a := []int64{10, 20, 30}
	b := []int64{30, 40}
	// 取出只在a，不在b的元素
	leftJoin := funk.Join(a, b, funk.LeftJoin)
	fmt.Println("只在a切片的元素: ", leftJoin)

	// 取出只在b，不在a的元素
	rightJoin := funk.Join(a, b, funk.RightJoin)
	fmt.Println("只在b切片的元素: ", rightJoin)
}

// 分别去掉两个切片共同元素(两结果集)
func TestDifferent(t *testing.T) {
	// 处理任意类型
	one := []interface{}{1, "go", 3.2, []int8{10}}
	two := []interface{}{2, "go", 3.2, []int{20}}
	oneRes, twoRes := funk.Difference(one, two)
	fmt.Println("oneRes: ", oneRes)
	fmt.Println("twoRes:  ", twoRes)

	// 只处理具体类型
	str1 := []string{"go", "php", "java"}
	str2 := []string{"c", "python", "java"}
	res, res1 := funk.DifferenceString(str1, str2)
	fmt.Println("res: ", res)
	fmt.Println("res1: ", res1)
}

//遍历切片
//ForEach: 从左边遍历切片
//ForEachRight: 从右边遍历切片

func TestForeachSlice(t *testing.T) {
	// 从左边遍历
	var leftRes []int
	funk.ForEach([]int{1, 2, 3, 4}, func(x int) {
		leftRes = append(leftRes, x*2)
	})
	fmt.Println("ForEach:", leftRes)

	// 从右边遍历
	var rightRes []int
	funk.ForEachRight([]int{1, 2, 3, 4}, func(x int) {
		rightRes = append(rightRes, x*2)
	})
	fmt.Println("ForEachRight:", rightRes)
}

// 除去第一个 或者 最后一个
func TestDelLastOrFirst(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	// 除去第一个，返回剩余元素
	fmt.Println(funk.Tail(a))
	// 除去最后一个，返回剩余元素
	fmt.Println(funk.Initial(a))
}

// 判断A切片是否属于B切片子集
// Subset(x interface{}, y interface{}) bool: 判断x是否属于y的切片
// 判断是否属于子集
func TestSubset(t *testing.T) {
	// 判断基础切片
	fmt.Println("是否属于子集:", funk.Subset([]int{1}, []int{1, 2}))
	// 判断自定义结构体切片
	subStu1 := []Student{{Name: "张三", Age: 18}}
	subStu2 := []Student{{Name: "张三", Age: 22}}
	allStu := []Student{
		{Name: "张三", Age: 18},
		{Name: "李四", Age: 22},
	}
	fmt.Println("subStu1: ", funk.Subset(subStu1, allStu))
	fmt.Println("subStu2: ", funk.Subset(subStu2, allStu))
	// 判断空切片是否属于另一切片子集
	fmt.Println("判断空集：", funk.Subset([]int{}, []int{1, 2}))
}

// 分组 Chunk(arr interface{}, size int) interface{}: 把arr按照每组size个进行分组
func TestChunk(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	// 分组(每组2个)
	fmt.Println(funk.Chunk(a, 2))

	// 分组自定义结构体切片 (每组2个)
	stuList := []Student{
		{Name: "张三", Age: 18},
		{Name: "小明", Age: 20},
		{Name: "李四", Age: 22},
		{Name: "赵武", Age: 32},
		{Name: "小英", Age: 19},
	}
	fmt.Println(funk.Chunk(stuList, 2))
}

type Book struct {
	Id   int
	Name string
}

// 把结构体切片转成map
// ToMap(in interface{}, pivot string) interface{}: 把切片in,转成以pivot为key的map
func TestToMap(t *testing.T) {
	bookList := []Book{
		{Id: 1, Name: "西游记"},
		{Id: 2, Name: "水浒传"},
		{Id: 3, Name: "三国演义"},
	}
	// 转成以Id为Key的Map
	fmt.Println("结果:", funk.ToMap(bookList, "Id"))
}

// 把切片值转成Map中的Key
func TestMap(t *testing.T) {
	// 将切片转为map的key
	r := funk.Map([]int{1, 2, 3, 4}, func(x int) (int, string) {
		return x, "go"
	})
	fmt.Println("r=", r)
}

// 把二维切片转成一维切片
func TestFlatMap(t *testing.T) {
	r := funk.FlatMap([][]int{{1}, {2}, {3}, {4}}, func(x []int) []int {
		return x
	})
	fmt.Printf("%#v\n", r)
}

// 打乱切片
func TestShuffle(t *testing.T) {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7}
	// 打乱多次
	for i := 1; i <= 3; i++ {
		fmt.Println(fmt.Sprintf("第%v次打乱a", i), funk.Shuffle(a))
	}
}

// 反转切片
func TestReverse(t *testing.T) {
	fmt.Println("ReverseInt:", funk.ReverseInt([]int{1, 2, 3, 4, 5, 6}))
	fmt.Println("Reverse,任意类型:", funk.Reverse([]interface{}{1, "2", 3.02, 4, "5", 6}))
	fmt.Println("Reverse,Str:", funk.ReverseStrings([]string{"a", "b", "c", "d"}))
}

// 元素去重
func TestUniq(t *testing.T) {
	a := []int64{1, 2, 3, 4, 3, 2, 1}
	// 过滤整型类型
	fmt.Println("Uniq:", funk.Uniq(a))
	fmt.Println("UniqInt64:", funk.UniqInt64(a))

	b := []string{"php", "go", "c", "go"}
	// 过滤字符串类型
	fmt.Println("UniqString:", funk.UniqString(b))
}

// 删除指定元素
func TestWithOut(t *testing.T) {
	// 删除具体元素
	b := []int{10, 20, 30, 40}
	without := funk.Without(b, 30)
	fmt.Println("删除30:", without)

	// 删除自定义结构体元素
	stuList := []Student{
		{Name: "张三", Age: 18},
		{Name: "李四", Age: 22},
		{Name: "范五", Age: 24},
	}
	res := funk.Without(stuList, Student{Name: "李四", Age: 22})
	fmt.Println("删除李四:", res)
}
