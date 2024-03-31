package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

type Respond struct {
	Success bool
	Status  string
	Data    []Proxy
}

type Proxy struct {
	Remark        string //描述
	Prefix        string //转发的前缀判断
	Upstream      string //后端 nginx 地址或者ip地址
	RewritePrefix string //重写
}

var (
	InfoLog  *log.Logger
	ErrorLog *log.Logger
	proxyMap = make(map[string]Proxy)
)

var adminUrl = flag.String("adminUrl", "", "admin的地址")
var profile = flag.String("profile", "", "环境")
var proxyFile = flag.String("proxyFile", "", "测试环境的数据")

// 日志初始化
func initLog() {
	errFile, err := os.OpenFile("errors.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	infoFile, err := os.OpenFile("info.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("打开日志文件失败：", err)
	}
	InfoLog = log.New(io.MultiWriter(os.Stderr, infoFile), "Info:", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
	ErrorLog = log.New(io.MultiWriter(os.Stderr, errFile), "Error:", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
}

func main() {
	router := gin.Default() //创建一个router
	flag.Parse()
	initLog()
	if *profile != "" {
		InfoLog.Printf("加载远端数据: %s ", *adminUrl)
		initProxyList()
	} else {
		InfoLog.Printf("加载本地配置数据: %s", *proxyFile)
		loadProxyListFromFile()
	}
	router.Any("/*action", Forward) //所有请求都会经过Forward函数转发

	router.Run(":8000")
}

//curl "http://127.0.0.1:8000/"

func initProxyList() {
	resp, _ := http.Get(*adminUrl)
	if resp != nil && resp.StatusCode == 200 {
		bytes, err := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		if err != nil {
			fmt.Println("ioutil.ReadAll err=", err)
			return
		}
		var respond Respond
		err = json.Unmarshal(bytes, &respond)
		if err != nil {
			fmt.Println("json.Unmarshal err=", err)
			return
		}
		proxyList := respond.Data
		for _, proxy := range proxyList {
			//追加 反斜杠，为了动态匹配的时候 防止 /proxy/test  /proxy/test1 无法正确转发
			proxyMap[proxy.Prefix+"/"] = proxy
		}
	}
}

func Forward(c *gin.Context) {
	HostReverseProxy(c.Writer, c.Request)
}

func HostReverseProxy(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HostReverseProxy")
	if r.RequestURI == "/favicon.ico" {
		io.WriteString(w, "Request path Error")
		return
	}
	//从内存里面获取转发的url
	var upstream = ""
	if value, ok := proxyMap[r.RequestURI]; ok {
		//如果转发的地址是 / 开头的,需要去掉
		if strings.HasSuffix(value.Upstream, "/") {
			upstream += strings.TrimRight(value.Upstream, "/")
		} else {
			upstream += value.Upstream
		}

		//如果首位不是/开头，则需要追加
		if !strings.HasPrefix(value.RewritePrefix, "/") {
			upstream += "/" + value.RewritePrefix
		} else {
			upstream += value.RewritePrefix
		}

		//去掉开头
		r.URL.Path = strings.ReplaceAll(r.URL.Path, r.RequestURI, "")
	}

	// parse the url
	remote, err := url.Parse(upstream)
	InfoLog.Printf("RequestURI %s upstream %s remote %s", r.RequestURI, upstream, remote)
	if err != nil {
		panic(err)
	}

	r.URL.Host = remote.Host
	r.URL.Scheme = remote.Scheme
	r.Header.Set("X-Forwarded-Host", r.Header.Get("Host"))
	r.Host = remote.Host

	httputil.NewSingleHostReverseProxy(remote).ServeHTTP(w, r)
}

func loadProxyListFromFile() {
	file, err := os.Open(*proxyFile)
	if err != nil {
		ErrorLog.Println("err:", err)
	}
	var respond Respond
	// 创建json解码器
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&respond)
	if err != nil {
		fmt.Println("LoadProxyListFromFile failed", err.Error())
	}
	proxyList := respond.Data
	for _, proxy := range proxyList {
		proxyMap[proxy.Prefix+"/"] = proxy
	}
	fmt.Printf("proxyMap %+v", proxyMap)
}

//加载本地配置文件数据
//go run proxy_agent.go -proxyFile ./proxy_data.json
//启动从配置中心获取数据
//go run proxy_agent.go -profile prod -adminUrl http://localhost:3000/proxy/findAll
