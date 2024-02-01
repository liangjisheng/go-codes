package main

import (
	"fmt"
	"strings"

	jsoniter "github.com/json-iterator/go"
)

func demo1() {
	type ColorGroup struct {
		ID     int
		Name   string
		Colors []string
	}
	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}

	b, err := jsoniter.Marshal(group)
	if err != nil {
		return
	}
	fmt.Println(string(b))

	val := []byte(`{"ID":1,"Name":"Reds","Colors":["Crimson","Red","Ruby","Maroon"]}`)
	color := jsoniter.Get(val, "Colors", 0).ToString()
	fmt.Println(color)

	m := map[string]interface{}{
		"3": 3,
		"1": 1,
		"2": 2,
	}
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	b, err = json.Marshal(m)
	if err != nil {
		return
	}
	fmt.Println(string(b))
}

// Student 测试转换json的结构体
type Student struct {
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Address string   `json:"address"`
	Hobby   []string `json:"hobby"`
}

func marshalDemo() {
	stu := Student{
		Name:    "张三",
		Age:     18,
		Address: "chengdu",
		Hobby:   []string{"football", "swimming", "travel", "sing"},
	}

	// Marshal(v interface{}) ([]byte, error)
	//把结构体转化成json,兼容go标准库encoding/json的序列化方法,返回一个字节切片和错误
	b, err := jsoniter.Marshal(stu)
	if err != nil {
		fmt.Println("transformation error: ", err)
	}
	//输出转化后的字符串
	//{"name":"张三","age":18,"address":"chengdu","hobby":["football","swimming","travel","sing"]}
	fmt.Println(string(b))

	//直接把结构体转化成字符串MarshalToString,这里错误判断省略
	str, _ := jsoniter.MarshalToString(stu)
	//{"name":"张三","age":18,"address":"chengdu","hobby":["football","swimming","travel","sing"]}
	fmt.Println(str)

	//转化成字节切片,第一个参数是结构体对象，第二个参数是前缀字符串必须为"",第三个参数为缩进表示，只能是空格
	b, _ = jsoniter.MarshalIndent(stu, "", " ")
	//输出带格式的字符串
	fmt.Println(string(b))
}

func unmarshalDemo() {
	//反序列化给一个结构体
	var stu Student
	var jsonBlob = []byte(`
        {"name": "张三", "Age": 12,"Address":"chengdu","Hobby":["football", "swimming", "travel", "sing"]}
    `)
	//根据字节切片转换成结构体。注意这里传入的是结构体地址
	err := jsoniter.Unmarshal(jsonBlob, &stu)
	if err != nil {
		fmt.Println("unmarshal error: ", err)
	}
	fmt.Printf("根据字节切片转换的结构体数据：%+v\n", stu)

	//根据字节切片转换成结构体
	var students []Student
	var jsonSlice = []byte(`[
        {"name": "张三", "Age": 12,"Address":"chengdu","Hobby":["football", "swimming", "travel", "sing"]},
        {"name": "李四", "Age": 28,"Address":"sichuan","Hobby":["dance", "music"]}
   ]`)
	//根据字节切片转换成结构体。把多个结构体json数据反序列化给结构体切片
	err = jsoniter.Unmarshal(jsonSlice, &students)
	if err != nil {
		fmt.Println("unmarshal error: ", err)
	}
	fmt.Printf("根据字节切片转换的结构体数据：%+v\n", students)

	var stu1 Student
	stuStr := `{"name": "张三", "Age": 12,"Address":"chengdu","Hobby":["football", "swimming", "travel", "sing"]`
	//根据字符串转换成结构体。注意这里传入的是结构体地址
	jsoniter.UnmarshalFromString(stuStr, &stu1)
	fmt.Printf("根据字符串转换的结构体数据：%+v\n", stu1)
}

func getValue() {
	var jsonBlob = []byte(`
        {"name": "张三", "Age": 12,"Address":"chengdu","Hobby":["football", "swimming", "travel", "sing"]}
    `)
	//获取深层嵌套JSON结构的值的快速方法
	//返回一个数组指针
	hobby := jsoniter.Get(jsonBlob, "Hobby")
	fmt.Printf("%T\n", hobby) //*jsoniter.arrayLazyAny

	//获取这个结构体Hobby元素下面的第二个元素值
	str := jsoniter.Get(jsonBlob, "Hobby", 1).ToString()
	fmt.Println(str) //swimming
}

func decoderDemo() {
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	reader := strings.NewReader(`{"branch":"beta","change_log":"add the rows{10}","channel":"fros","create_time":"2017-06-13 16:39:08","firmware_list":"","md5":"80dee2bf7305bcf179582088e29fd7b9","note":{"CoreServices":{"md5":"d26975c0a8c7369f70ed699f2855cc2e","package_name":"CoreServices","version_code":"76","version_name":"1.0.76"},"FrDaemon":{"md5":"6b1f0626673200bc2157422cd2103f5d","package_name":"FrDaemon","version_code":"390","version_name":"1.0.390"},"FrGallery":{"md5":"90d767f0f31bcd3c1d27281ec979ba65","package_name":"FrGallery","version_code":"349","version_name":"1.0.349"},"FrLocal":{"md5":"f15a215b2c070a80a01f07bde4f219eb","package_name":"FrLocal","version_code":"791","version_name":"1.0.791"}},"pack_region_urls":{"CN":"https://s3.cn-north-1.amazonaws.com.cn/xxx-os/ttt_xxx_android_1.5.3.344.393.zip","default":"http://192.168.8.78/ttt_xxx_android_1.5.3.344.393.zip","local":"http://192.168.8.78/ttt_xxx_android_1.5.3.344.393.zip"},"pack_version":"1.5.3.344.393","pack_version_code":393,"region":"all","release_flag":0,"revision":62,"size":38966875,"status":3}`)
	decoder := json.NewDecoder(reader)
	params := make(map[string]interface{})
	err := decoder.Decode(&params)
	if err != nil {
		fmt.Println(err)
	} else {
		//map[firmware_list: note:map[CoreServices:map[package_name:CoreServices version_code:76 version_name:1.0.76 md5:d26975c0a8c7369f70ed699f2855cc2e] FrDaemon:map[md5:6b1f0626673200bc2157422cd2103f5d package_name:FrDaemon version_code:390 version_name:1.0.390] FrGallery:map[version_code:349 version_name:1.0.349 md5:90d767f0f31bcd3c1d27281ec979ba65 package_name:FrGallery] FrLocal:map[version_name:1.0.791 md5:f15a215b2c070a80a01f07bde4f219eb package_name:FrLocal version_code:791]] pack_version:1.5.3.344.393 pack_version_code:393 status:3 channel:fros pack_region_urls:map[CN:https://s3.cn-north-1.amazonaws.com.cn/xxx-os/ttt_xxx_android_1.5.3.344.393.zip default:http://192.168.8.78/ttt_xxx_android_1.5.3.344.393.zip local:http://192.168.8.78/ttt_xxx_android_1.5.3.344.393.zip] release_flag:0 size:3.8966875e+07 md5:80dee2bf7305bcf179582088e29fd7b9 region:all revision:62 change_log:add the rows{10} create_time:2017-06-13 16:39:08 branch:beta]
		fmt.Printf("%+v\n", params)
	}
}

func main() {
	//demo1()
	//marshalDemo()
	//unmarshalDemo()
	//getValue()
	decoderDemo()
}
