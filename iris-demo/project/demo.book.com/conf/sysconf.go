package conf

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

// SysTimeform ...
const SysTimeform string = "2006-01-02 15:04:05"

// SysTimeformShort ...
const SysTimeformShort string = "2006-01-02"

// SysTimeLocation 中国时区
var SysTimeLocation, _ = time.LoadLocation("Asia/Chongqing")

// SysWebconfigPath 自定义配置文件地址
const SysWebconfigPath string = "./conf/web.config"

// SysConfMap 全局配置文件map
var SysConfMap map[string]string

func init() {
	ReLoad()
}

// ReLoad ...
func ReLoad() {
	config := make(map[string]string)

	f, err := os.Open(SysWebconfigPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		b, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		s := strings.TrimSpace(string(b))
		// 如果前两位是// 则视为注释
		if len(s) >= 2 && s[0:2] == "//" {
			continue
		}
		index := strings.Index(s, "=")
		if index < 0 {
			continue
		}
		key := strings.TrimSpace(s[:index])
		if len(key) == 0 {
			continue
		}
		value := strings.TrimSpace(s[index+1:])
		if len(key) == 0 {
			continue
		}
		config[key] = value
	}
	fmt.Printf("config: %+v\n", config)
	SysConfMap = config
}
