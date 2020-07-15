package main

import (
	"compress/gzip"
	"io"
	"log"
	"os"
)

// 使用gzip压缩文件

func Archive_gzip() {
	outputFile, err := os.Create("test.txt.gz")
	if err != nil {
		log.Fatal(err)
	}

	gzipWriter := gzip.NewWriter(outputFile)
	defer gzipWriter.Close()

	// 当我们写如到gizp writer数据时，它会依次压缩数据并写入到底层的文件中。
	// 我们不必关心它是如何压缩的，还是像普通的writer一样操作即可
	_, err = gzipWriter.Write([]byte("Gophers rule!\n"))
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Compressed data written to file.")
}

func Archive_ungzip() {
	gzipFile, err := os.Open("test.txt.gz")
	if err != nil {
		log.Fatal(err)
	}

	gzipReader, err := gzip.NewReader(gzipFile)
	if err != nil {
		log.Fatal(err)
	}
	defer gzipReader.Close()

	// 解压缩到一个writer,它是一个file writer
	outfileWriter, err := os.Create("unzipped.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer outfileWriter.Close()

	// copy
	_, err = io.Copy(outfileWriter, gzipReader)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	Archive_gzip()
	Archive_ungzip()
}
