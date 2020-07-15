package main

import (
	"archive/zip"
	"io"
	"log"
	"os"
	"path/filepath"
)

// this exmaple uses zip but standard library
// also support tar archive

// 打包(zip)文件
func Archive_zip() {
	// 创建一个打包文件
	outFile, err := os.Create("test.zip")
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()

	// 创建zip writer
	zipWriter := zip.NewWriter(outFile)
	// 写入内容
	var filesToArchive = []struct {
		Name, Body string
	}{
		{"test.txt", "String contents of file"},
		{"test2.txt", "\x61\x62\x63\n"},
	}

	// 将打包的内容依次写入打包文件
	for _, file := range filesToArchive {
		fileWriter, err := zipWriter.Create(file.Name)
		if err != nil {
			log.Fatal(err)
		}
		_, err = fileWriter.Write([]byte(file.Body))
		if err != nil {
			log.Fatal(err)
		}
	}

	// 清理
	err = zipWriter.Close()
	if err != nil {
		log.Fatal(err)
	}
}

// 抽取(unzip)文件
func Archive_unzip() {
	zipReader, err := zip.OpenReader("test.zip")
	if err != nil {
		log.Fatal(err)
	}
	defer zipReader.Close()

	// 遍历打包文件中的每一个文件/文件夹
	for _, file := range zipReader.Reader.File {
		// 打包文件中的文件就像普通的一个文件对象一样
		zippedFile, err := file.Open()
		if err != nil {
			log.Fatal(err)
		}
		defer zippedFile.Close()

		// 指定抽取的文件名
		// 你可以指定全路径名或者一个前缀，这样可以把它们放在不同的文件夹中
		// 我们这个例子使用打包文件中相同的文件名
		targetDir := "./"
		extractedFilePath := filepath.Join(
			targetDir,
			file.Name,
		)

		// 抽取项目或者创建文件夹
		if file.FileInfo().IsDir() {
			// 创建文件夹并设置同样的权限
			log.Println("Creating directory:", extractedFilePath)
			os.MkdirAll(extractedFilePath, file.Mode())
		} else {
			// 抽取正常的文件
			log.Println("Extracting file:", file.Name)

			outputFile, err := os.OpenFile(
				extractedFilePath,
				os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
				file.Mode(),
			)
			if err != nil {
				log.Fatal(err)
			}
			defer outputFile.Close()

			// 通过io.Copy()简洁的复制文件内容
			_, err = io.Copy(outputFile, zippedFile)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func main() {
	Archive_zip()
	Archive_unzip()
}
