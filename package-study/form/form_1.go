package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

func postFile(fileName string, targetUrl string) error {
	bodyBuf := &bytes.Buffer{}
	bodyWrite := multipart.NewWriter(bodyBuf)

	// 关键一步
	fileWriter, err := bodyWrite.CreateFormFile("uploadfile", fileName)
	if err != nil {
		fmt.Println("err writing to buffer")
		return err
	}

	// 打开文件句柄操作
	fh, err := os.Open(fileName)
	if err != nil {
		fmt.Println("error opening file")
		return err
	}
	defer fh.Close()

	// iocopy
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return err
	}

	contentType := bodyWrite.FormDataContentType()
	bodyWrite.Close()

	resp, err := http.Post(targetUrl, contentType, bodyBuf)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	resp_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println(resp.Status)
	fmt.Println(string(resp_body))
	return nil
}

// simple usage
func main() {
	targetUrl := "http://localhost:9090/upload"
	fileName := "./alice.txt"
	postFile(fileName, targetUrl)
}
