package utils

import (
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
	"os"
	"strconv"
)

// 前端传来文件片与当前片为什么文件的第几片
// 后端拿到以后比较次分片是否上传 或者是否为不完全片
// 前端发送每片多大
// 前端告知是否为最后一片且是否完成

const breakpointDir = "./breakpointDir/"
const finishDir = "./fileDir/"

// BreakPointContinue ...
//@function: BreakPointContinue
//@description: 断点续传
//@param: content []byte, fileName string, contentNumber int, contentTotal int, fileMd5 string
//@return: error, string
func BreakPointContinue(content []byte, fileName string, contentNumber int, contentTotal int, fileMd5 string) (string, error) {
	path := breakpointDir + fileMd5 + "/"
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return path, err
	}
	pathc, err := makeFileContent(content, fileName, path, contentNumber)
	return pathc, err

}

// CheckMd5 ...
//@description: 检查Md5
//@param: content []byte, chunkMd5 string
//@return: CanUpload bool
func CheckMd5(content []byte, chunkMd5 string) (CanUpload bool) {
	fileMd5 := MD5V(content)
	if fileMd5 == chunkMd5 {
		return true // "可以继续上传"
	}
	return false // "切片不完整，废弃"
}

//@description: 创建切片内容
//@param: content []byte, fileName string, FileDir string, contentNumber int
//@return: error, string
func makeFileContent(content []byte, fileName string, FileDir string, contentNumber int) (string, error) {
	path := FileDir + fileName + "_" + strconv.Itoa(contentNumber)
	f, err := os.Create(path)
	if err != nil {
		return path, err
	} else {
		_, err = f.Write(content)
		if err != nil {
			return path, err
		}
	}
	defer f.Close()
	return path, nil
}

// MakeFile ...
//@description: 创建切片文件
//@param: fileName string, FileMd5 string
//@return: error, string
func MakeFile(fileName string, FileMd5 string) (string, error) {
	rd, err := ioutil.ReadDir(breakpointDir + FileMd5)
	if err != nil {
		return finishDir + fileName, err
	}
	_ = os.MkdirAll(finishDir, os.ModePerm)
	fd, err := os.OpenFile(finishDir+fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return finishDir + fileName, err
	}
	defer fd.Close()
	for k := range rd {
		content, _ := ioutil.ReadFile(breakpointDir + FileMd5 + "/" + fileName + "_" + strconv.Itoa(k))
		_, err = fd.Write(content)
		if err != nil {
			_ = os.Remove(finishDir + fileName)
			return finishDir + fileName, err
		}
	}

	return finishDir + fileName, nil
}

// RemoveChunk ...
//@description: 移除切片
//@param: FileMd5 string
func RemoveChunk(FileMd5 string) error {
	err := os.RemoveAll(breakpointDir + FileMd5)
	return err
}

// MD5V ...
func MD5V(str []byte) string {
	h := md5.New()
	h.Write(str)
	return hex.EncodeToString(h.Sum(nil))
}
