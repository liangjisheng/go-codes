package main

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

func main() {
	// 创建表单文件，CreateFormFile用来创建表单文件，第一个参数是字段名，
	// 第二个参数是文件名
	buf := new(bytes.Buffer)
	writer := multipart.NewWriter(buf)
	filename := "test.jpg"
	formFile, err := writer.CreateFormFile("uploadfile", filename)
	if err != nil {
		log.Fatalf("Create form file failed: %s\n", err)
	}

	// 从文件中读取数据，写入表单
	srcFile, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Open source file failed: %s\n", err)
	}
	defer srcFile.Close()

	_, err = io.Copy(formFile, srcFile)
	if err != nil {
		log.Fatalf("Write to form file failed: %s\n", err)
	}

	// 发送表单
	contentType := writer.FormDataContentType()
	writer.Close() // 发送前必须调用Close()写入结尾行
	_, err = http.Post("http://localhost:9090/upload", contentType, buf)
	if err != nil {
		log.Fatalf("Post failed: %s\n", err)
	}
}
