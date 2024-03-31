package main

//https://cloud.tencent.com/developer/article/2211963

import (
	"encoding/json"
	"errors"
	"fmt"
	"unsafe"

	"github.com/mohae/deepcopy"
	"github.com/shopspring/decimal"
)

// DeepCopy 本来以为copy是深拷，想使用内置函数实现任意类型的深拷，可惜实现不了了！！！
func DeepCopy(a interface{}) (interface{}, error) {
	c := make([]interface{}, 1)
	b := make([]interface{}, 1)
	c[0] = a
	num := copy(b, c)
	if num != 1 {
		return nil, errors.New("copy failed")
	}
	return b[0], nil
}

type Email struct {
	Account  string
	Password string
}

type ComStruct struct {
	Name      string
	Age       int
	PhoneList []string
	foo       map[string]string
	email     Email
	Dec       decimal.Decimal
}

func copy1() {
	a := &ComStruct{
		Name:      "jam",
		Age:       23,
		PhoneList: []string{"a1", "a2", "a3"},
		foo:       map[string]string{"a1": "a1", "a2": "a2", "a3": "a3"},
		email:     Email{"jam", "123456"},
		Dec:       decimal.Zero,
	}
	//golang的内置copy函数是无法完成引用的深度拷贝，其内置函数copy是值拷贝，拷贝引用后，引用的地址相同
	b, _ := DeepCopy(a)
	b1, ok := b.(*ComStruct)
	if !ok {
		panic("not ComStruct")
	}
	b1.email.Account = "alice"
	b1.Dec = decimal.NewFromInt(1)
	fmt.Printf("%v\n", (*a).email)
	fmt.Printf("%v\n", (*b1).email)
	fmt.Printf("%v\n", (*a).Dec)
	fmt.Printf("%v\n", (*b1).Dec)

	fmt.Printf("%p\n", a)
	fmt.Printf("%p\n", b)
}

func copy2() {
	e1 := make([]*Email, 0)
	e1 = append(e1, &Email{
		Account:  "alice",
		Password: "123",
	})
	e1 = append(e1, &Email{
		Account:  "alice1",
		Password: "1234",
	})

	e2 := make([]*Email, 2)
	copy(e2, e1)

	fmt.Printf("%p\n", e1)
	fmt.Printf("%p\n", e2)

	e1[1].Account = "bob"
	fmt.Println("e1", e1[1].Account) // bob
	fmt.Println("e2", e2[1].Account) // bob
}

func copyIntPoint() {
	int1 := 1
	int2 := 2
	int3 := 3
	ptrSlice1 := []*int{&int1, &int2, &int3}
	ptrSlice2 := make([]*int, len(ptrSlice1))
	copy(ptrSlice2, ptrSlice1)
	fmt.Printf("2--%p---%p--不相同\n", ptrSlice1, ptrSlice2)
	fmt.Printf("3--%p---%p--相同\n", ptrSlice1[0], ptrSlice2[0])
	*(ptrSlice2[0]) = 4
	fmt.Printf("%v\n", *(ptrSlice1[0])) // 4
	fmt.Printf("%v\n", *(ptrSlice2[0])) // 4
}

func copyInt() {
	a := []int{1, 2, 3}
	b := make([]int, len(a), len(a))
	copy(b, a)
	fmt.Println(unsafe.Pointer(&a)) // 0xc00000c030
	fmt.Println(a, &a[0])           // [1 2 3] 0xc00001a078
	fmt.Println(unsafe.Pointer(&b)) // 0xc00000c048
	fmt.Println(b, &b[0])           // [1 2 3] 0xc00001a090

	a[0] = 4
	fmt.Println(a[0]) // 4
	fmt.Println(b[0]) // 1
}

func copyIntPoint1() {
	int1 := 1
	int2 := 2
	int3 := 3
	ptrSlice1 := []*int{&int1, &int2, &int3}

	ptrSliceInter := deepcopy.Copy(ptrSlice1)
	ptrSlice2, _ := ptrSliceInter.([]*int)
	fmt.Printf("4--%p---%p--不相同\n", ptrSlice1, ptrSlice2)
	fmt.Printf("5--%p---%p--不相同\n", ptrSlice1[0], ptrSlice2[0])

	*(ptrSlice2[0]) = 4
	fmt.Printf("%v\n", *(ptrSlice1[0])) // 1
	fmt.Printf("%v\n", *(ptrSlice2[0])) // 4
}

func copy5() {
	e1 := make([]*Email, 0)
	e1 = append(e1, &Email{
		Account:  "alice",
		Password: "123",
	})
	e1 = append(e1, &Email{
		Account:  "alice1",
		Password: "1234",
	})

	emailSlice := deepcopy.Copy(e1)
	e2, _ := emailSlice.([]*Email)
	fmt.Printf("%p\n", e1)
	fmt.Printf("%p\n", e2)

	e1[0].Account = "bob"
	fmt.Println("e1", e1[0].Account) // bob
	fmt.Println("e2", e2[0].Account) // alice
}

func copy6() {
	type S6 struct {
		F1 int
		F2 *int
		F3 Email
		F4 *Email
	}

	pi := new(int)
	*pi = 1

	d1 := S6{
		F1: 0,
		F2: pi,
		F3: Email{
			Account:  "alice",
			Password: "123",
		},
		F4: &Email{
			Account:  "alice",
			Password: "456",
		},
	}

	d1I := deepcopy.Copy(d1)
	d2, _ := d1I.(S6)
	fmt.Printf("%p\n", &d1)
	fmt.Printf("%p\n", &d2)

	d1.F3.Account = "bob"
	d1.F4.Account = "bob"
	b1, _ := json.Marshal(d1)
	b2, _ := json.Marshal(d2)
	fmt.Printf("d1 %s\n", string(b1))
	fmt.Printf("d2 %s\n", string(b2))
}

func main() {
	//copy1()
	//copy2()
	//copy3()
	//copyInt()
	//copyIntPoint()
	//copyIntPoint1()
	//copy5()
	copy6()
}
