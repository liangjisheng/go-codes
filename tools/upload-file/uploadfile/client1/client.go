package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
	"path/filepath"
	"strings"
)

var (
	filePath string
	addr     string
)

func init() {
	flag.StringVar(&filePath, "file", "", "the file to upload")
	flag.StringVar(&addr, "addr", "localhost:8080", "the addr of file server")
	flag.Parse()
}

var quoteEscaper = strings.NewReplacer("\\", "\\\\", `"`, "\\\"")

func escapeQuotes(s string) string {
	return quoteEscaper.Replace(s)
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

	// 自定义file分段中的header
	// file part1
	_, fileName := filepath.Split(filePath)
	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition",
		fmt.Sprintf(`form-data; name="%s"; filename="%s"`,
			escapeQuotes("file1"), escapeQuotes(fileName)))
	h.Set("Content-Type", "text/plain")

	//fw1, _ := bw.CreateFormFile("file1", fileName)
	fw1, _ := bw.CreatePart(h)
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
