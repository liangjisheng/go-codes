package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

// 通过http下载文件

func main() {
	newFile, err := os.Create("devdungeon.html")
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	url := "http://www.baidu.com"
	response, err := http.Get(url)
	defer response.Body.Close()

	// 将HTTP response Body中的内容写入到文件
	// Body满足reader接口，因此我们可以使用ioutil.Copy
	numBytesWritten, err := io.Copy(newFile, response.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Downloaded %d byte file.\n", numBytesWritten)
}
