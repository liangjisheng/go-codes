package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
	"path/filepath"
	"strings"
)

//前面的客户端中存在一个问题，那就是我们在构建http body的时候，使用了一个bytes.Buffer
//加载了待上传文件的所有内容，这样一来，如果待上传的文件很大的话，内存空间消耗势必过大
//那么如何将每次上传内存文件时对内存的使用限制在一个适当的范围，或者说上传文件所消耗的
//内存空间不因待传文件的变大而变大呢？

var (
	filePath string
	addr     string
)

var quoteEscaper = strings.NewReplacer("\\", "\\\\", `"`, "\\\"")

func escapeQuotes(s string) string {
	return quoteEscaper.Replace(s)
}

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
	pr, pw := io.Pipe()
	bw := multipart.NewWriter(pw)

	f, err := os.Open(filePath)
	if err != nil {
		return "", nil, err
	}

	go func() {
		defer f.Close()

		// text part1
		p1w, _ := bw.CreateFormField("name")
		p1w.Write([]byte("alice"))

		// test part2
		p2w, _ := bw.CreateFormField("age")
		p2w.Write([]byte("15"))

		// file part1
		_, fileName := filepath.Split(filePath)
		h := make(textproto.MIMEHeader)
		h.Set("Content-Disposition",
			fmt.Sprintf(`form-data; name="%s"; filename="%s"`,
				escapeQuotes("file1"), escapeQuotes(fileName)))
		h.Set("Content-Type", "application/pdf")
		fw1, _ := bw.CreatePart(h)
		cnt, _ := io.Copy(fw1, f)
		log.Printf("copy %d bytes from file %s in total\n", cnt, fileName)

		bw.Close() // write the tail boundary
		pw.Close()
	}()

	return bw.FormDataContentType(), pr, nil
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
