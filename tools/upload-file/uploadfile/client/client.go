package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

var (
	filePath string
	addr     string
)

//go run client.go -file /Users/liangjisheng/Downloads/2021926-12542.jpeg

func init() {
	flag.StringVar(&filePath, "file", "", "the file to upload")
	flag.StringVar(&addr, "addr", "localhost:8080", "the addr of file server")
	flag.Parse()
}

func main() {
	if filePath == "" {
		fmt.Println("file must not be empty")
		return
	}

	err := doUpload(addr, filePath)
	if err != nil {
		fmt.Printf("upload file [%s] error: %s", filePath, err)
		return
	}
	fmt.Printf("upload file [%s] ok\n", filePath)
}

func createReqBody(filePath string) (string, io.Reader, error) {
	var err error
	buf := new(bytes.Buffer)
	bw := multipart.NewWriter(buf)

	f, err := os.Open(filePath)
	if err != nil {
		return "", nil, err
	}
	defer f.Close()

	// text part1
	p1w, _ := bw.CreateFormField("name")
	p1w.Write([]byte("alice"))

	// test part2
	p2w, _ := bw.CreateFormField("age")
	p2w.Write([]byte("15"))

	// file part1
	_, fileName := filepath.Split(filePath)
	fw1, _ := bw.CreateFormFile("file1", fileName)
	io.Copy(fw1, f)

	bw.Close() // write the tail boundary
	return bw.FormDataContentType(), buf, nil
}

func doUpload(addr, filePath string) error {
	// create body
	contType, reader, err := createReqBody(filePath)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("http://%s/uploadMulti", addr)
	req, _ := http.NewRequest("POST", url, reader)

	// add headers
	req.Header.Add("Content-Type", contType)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("request send error:", err)
		return err
	}
	resp.Body.Close()
	return nil
}

// form-data 上传文件，并获取返回结果
func formDataUploadFile(url string, header, formParam map[string]string, fieldName string, filename string) ([]byte, error) {
	reqBody := &bytes.Buffer{}
	writer := multipart.NewWriter(reqBody)
	formFile, err := writer.CreateFormFile(fieldName, filepath.Base(filename))
	if err != nil {
		return nil, err
	}

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	// 从文件中读取数据，写入表单
	_, err = io.Copy(formFile, file)

	for key, val := range formParam {
		_ = writer.WriteField(key, val)
	}

	err = writer.Close() //write the tail boundary
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, reqBody)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	//设置http请求头
	for key, value := range header {
		req.Header.Add(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	respBody := &bytes.Buffer{}
	_, err = respBody.ReadFrom(resp.Body)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return respBody.Bytes(), nil
}
