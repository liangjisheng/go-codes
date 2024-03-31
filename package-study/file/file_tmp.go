package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// 在系统临时文件夹中创建一个临时文件夹
	tempDirPath, err := ioutil.TempDir("", "myTempDir")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Tmpe dir created:", tempDirPath)

	// 在临时文件夹中创建临时文件
	tempFile, err := ioutil.TempFile(tempDirPath, "myTempFile.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Temp file created:", tempFile.Name())

	// do something

	// 关闭文件
	err = tempFile.Close()
	if err != nil {
		log.Fatal(err)
	}

	// 删除创建的资源
	err = os.Remove(tempFile.Name())
	if err != nil {
		log.Fatal(err)
	}

	err = os.Remove(tempDirPath)
	if err != nil {
		log.Fatal(err)
	}
}
