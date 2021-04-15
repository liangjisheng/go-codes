package config

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

// Config ...
type Config struct {
	filename       string
	lastModifyTime int64
	data           map[string]string
	rwLock         sync.RWMutex
	// 用于将调用该包的程序追加到切片中，用于通知调用上面在
	// config_notify.go定义的callback回调函数
	notifyList []Notifyer
}

// NewConfig ...
func NewConfig(filename string) (conf *Config, err error) {
	conf = &Config{
		filename: filename,
		data:     make(map[string]string, 1024),
	}
	m, err := conf.parse()
	if err != nil {
		return
	}
	conf.rwLock.Lock()
	conf.data = m
	conf.rwLock.Unlock()
	go conf.reload()
	return
}

// AddNotifyer ...
func (c *Config) AddNotifyer(n Notifyer) {
	c.notifyList = append(c.notifyList, n)
}

func (c *Config) parse() (m map[string]string, err error) {
	m = make(map[string]string, 1024)
	file, err := os.Open(c.filename)
	if err != nil {
		log.Panicln("open", c.filename, "fail.")
		return
	}
	var lineNo int
	reader := bufio.NewReader(file)
	for {
		// 一行行的读文件
		line, errRet := reader.ReadString('\n')
		if errRet == io.EOF {
			break // 读到文件末尾
		}
		if errRet != nil {
			err = errRet
			return
		}
		lineNo++
		line = strings.TrimSpace(line) // 去除空格
		if len(line) == 0 || line[0] == '\n' || line[0] == '+' || line[0] == ';' {
			// 当前行为空行或者是注释行等
			continue
		}

		// 通过=进行切割取出k/v结构
		arr := strings.Split(line, "=")
		if len(arr) == 0 {
			log.Printf("invalid config, line:%d\n", lineNo)
			continue
		}
		key := strings.TrimSpace(arr[0])
		if len(key) == 0 {
			log.Printf("invalid config, line:%d\n", lineNo)
			continue
		}
		if len(arr) == 1 {
			m[key] = ""
			continue
		}
		value := strings.TrimSpace(arr[1])
		m[key] = value
	}
	return
}

func (c *Config) reload() {
	ticker := time.NewTicker(time.Second * 5)
	for range ticker.C {
		func() {
			file, err := os.Open(c.filename)
			if err != nil {
				log.Printf("open %s failed,err:%v\n", c.filename, err)
				return
			}
			defer file.Close()
			fileInfo, err := file.Stat()
			if err != nil {
				log.Printf("stat %s failed,err:%v\n", c.filename, err)
				return
			}
			curModifyTime := fileInfo.ModTime().Unix()
			log.Printf("%v --- %v\n", curModifyTime, c.lastModifyTime)
			if curModifyTime > c.lastModifyTime {
				m, err := c.parse()
				if err != nil {
					log.Println("parse failed,err:", err)
					return
				}
				c.rwLock.Lock()
				c.data = m
				c.rwLock.Unlock()
				for _, n := range c.notifyList {
					n.Callback(c)
				}
				c.lastModifyTime = curModifyTime
			}
		}()
	}
}

// GetInt ...
func (c *Config) GetInt(key string) (value int, err error) {
	// 根据int获取
	c.rwLock.RLock()
	defer c.rwLock.RUnlock()
	str, ok := c.data[key]
	if !ok {
		err = fmt.Errorf("key[%s] not found", key)
		return
	}
	value, err = strconv.Atoi(str)
	return
}

// GetIntDefault ...
func (c *Config) GetIntDefault(key string, defval int) (value int) {
	// 默认值
	c.rwLock.RLock()
	defer c.rwLock.RUnlock()
	str, ok := c.data[key]
	if !ok {
		value = defval
		return
	}
	value, err := strconv.Atoi(str)
	if err != nil {
		value = defval
		return
	}
	return
}

// GetString ...
func (c *Config) GetString(key string) (value string, err error) {
	// 根据字符串获取
	c.rwLock.RLock()
	defer c.rwLock.RUnlock()
	value, ok := c.data[key]
	if !ok {
		err = fmt.Errorf("key[%s] not found", key)
		return
	}
	return
}
