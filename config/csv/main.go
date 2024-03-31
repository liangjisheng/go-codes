package main

import (
	"encoding/csv"
	"log"
	"os"
)

func main() {
	filename := "test.csv"
	ReadCsv(filename)
	//WriterCSV(filename)
}

func ReadCsv(filepath string) [][]string {
	file, err := os.Open(filepath)
	if err != nil {
		log.Println("csv文件打开失败")
		return nil
	}
	defer file.Close()

	//创建csv读取接口实例
	reader := csv.NewReader(file)

	//获取一行内容，一般为第一行内容标题
	read, _ := reader.Read() //返回切片类型：[chen  hai wei]
	log.Println(read)

	//读取所有内容
	contents, err := reader.ReadAll() //返回切片类型：[[s s ds] [a a a]]
	log.Println("len", len(contents))
	log.Println(contents[0])
	return contents

	//1、读取csv文件返回的内容为切片类型，可以通过遍历的方式使用或Slicer[0]方式获取具体的值。
	//2、同一个函数或线程内，两次调用Read()方法时，第二次调用时得到的值为每二行数据，依此类推。
	//3、大文件时使用逐行读取，小文件直接读取所有然后遍历，两者应用场景不一样，需要注意。
}

func WriterCSV(path string) {
	//OpenFile读取文件，不存在时则创建，使用追加模式
	file, err := os.OpenFile(path, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		log.Println("文件打开失败！")
		return
	}
	defer file.Close()

	//创建写入接口
	writer := csv.NewWriter(file)
	str := []string{"chen1", "hai1", "wei1"} //需要写入csv的数据，切片类型

	//写入一条数据，传入数据为切片(追加模式)
	err1 := writer.Write(str)
	if err1 != nil {
		log.Println("WriterCsv写入文件失败")
		return
	}
	writer.Flush() //刷新，不刷新是无法写入的
	log.Println("数据写入成功...")
}
