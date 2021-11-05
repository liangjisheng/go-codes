package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

const uploadPath = "./upload"

func main() {
	http.HandleFunc("/uploadSingle", handleUnloadSingle)
	http.HandleFunc("/uploadMulti", handleUnloadMulti)
	log.Println("server listen on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleUnloadSingle(w http.ResponseWriter, r *http.Request) {
	// 根据字段名获取表单文件
	formFile, header, err := r.FormFile("uploadfile")
	if err != nil {
		log.Printf("Get form file failed: %s\n", err)
		return
	}
	defer formFile.Close()

	// 创建保存文件
	destFile, err := os.Create("." + r.URL.Path + "/" + header.Filename)
	if err != nil {
		log.Printf("Create failed: %s\n", err)
		return
	}
	defer destFile.Close()

	// 读取表单文件，写入保存文件
	_, err = io.Copy(destFile, formFile)
	if err != nil {
		log.Printf("Write file failed: %s\n", err)
		return
	}
}

func handleUnloadMulti(w http.ResponseWriter, r *http.Request) {
	// 设置缓冲区大小
	r.ParseMultipartForm(16 * 1024 * 1024)
	mForm := r.MultipartForm

	for k, _ := range mForm.File {
		// k is the key of file part
		file, fileHeader, err := r.FormFile(k)
		if err != nil {
			fmt.Println("inovke FormFile error:", err)
			return
		}
		defer file.Close()

		fmt.Printf("the uploaded file: name[%s], size[%d], header[%#v]\n",
			fileHeader.Filename, fileHeader.Size, fileHeader.Header)

		// store uploaded file into local path
		localFileName := uploadPath + "/" + fileHeader.Filename
		out, err := os.Create(localFileName)
		if err != nil {
			fmt.Printf("failed to open the file %s for writing", localFileName)
			return
		}
		defer out.Close()
		_, err = io.Copy(out, file)
		if err != nil {
			fmt.Printf("copy file err:%s\n", err)
			return
		}
		fmt.Printf("file %s uploaded ok\n", fileHeader.Filename)
	}
}
