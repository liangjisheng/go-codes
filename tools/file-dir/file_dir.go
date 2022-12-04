package filedir

import (
	"fmt"
	"go.uber.org/zap"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)

// MkDirAll 创建多级目录
func MkDirAll(path string) bool {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

// Exist 检测文件或者文件夹是否存在
func Exist(file string) bool {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	}
	return true
}

// Ext 获取文件的类型，如: .jpg
// 如果获取不到，返回默认类型defaultExt
func Ext(fileName string, defaultExt string) string {
	t := path.Ext(fileName)
	if len(t) == 0 {
		return defaultExt
	}
	return t
}

// MakeDir 检测文件夹是否存在，不存在则创建
func MakeDir(filePath string) {
	if !Exist(filePath) {
		MkDirAll(filePath)
	}
}

// GetFilesAndDirs 获取指定目录下的所有文件和目录
func GetFilesAndDirs(dirPth string) (files []string, dirs []string, err error) {
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, nil, err
	}

	// PthSep := string(os.PathSeparator)
	PthSep := "/"
	// suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写

	for _, fi := range dir {
		if fi.IsDir() { // 目录, 递归遍历
			dirs = append(dirs, dirPth+PthSep+fi.Name())
			subFiles, subDirs, _ := GetFilesAndDirs(dirPth + PthSep + fi.Name())
			files = append(files, subFiles...)
			dirs = append(dirs, subDirs...)
		} else {
			// 过滤指定格式
			//ok := strings.HasSuffix(fi.Name(), ".go")
			//if ok {
			//	files = append(files, dirPth+PthSep+fi.Name())
			//}

			files = append(files, dirPth+PthSep+fi.Name())
		}
	}

	return files, dirs, nil
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func CreateDir(dirs ...string) error {
	for _, v := range dirs {
		exist, err := PathExists(v)
		if err != nil {
			return err
		}
		if !exist {
			fmt.Println("create directory" + v)
			err = os.MkdirAll(v, os.ModePerm)
			if err != nil {
				fmt.Println("create directory"+v, zap.Any(" error:", err))
			}
		}
	}
	return nil
}

// DeleteFile 删除文件或文件夹
func DeleteFile(absDir string) error {
	return os.RemoveAll(absDir)
}

// GetModelPath 获取程序运行目录
func GetModelPath() string {
	dir, _ := os.Getwd()
	return strings.Replace(dir, "\\", "/", -1)
}

// MoveFile 移动文件或文件夹(/结尾)
func MoveFile(from, to string) error {
	return os.Rename(from, to)
}

func CopyFile(src, des string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	desFile, err := os.Create(des)
	if err != nil {
		return err
	}
	defer desFile.Close()

	_, err = io.Copy(desFile, srcFile)
	return err
}
