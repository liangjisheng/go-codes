package handlers

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"chitchat/config"
	"chitchat/models"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

var logger *log.Logger
var localizer *i18n.Localizer

// var conf *config.Configuration

func init() {
	file, err := os.OpenFile("logs/chitchat.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file", err)
	}
	logger = log.New(file, "INFO ", log.Ldate|log.Ltime|log.Lshortfile)

	// conf = config.LoadConfig()                                          // 获取全局配置实例
	localizer = i18n.NewLocalizer(config.ViperConfig.LocaleBundle, config.ViperConfig.App.Language) // 获取本地化实例
}

// 通过 Cookie 判断用户是否已登录
func session(writer http.ResponseWriter, request *http.Request) (sess models.Session, err error) {
	cookie, err := request.Cookie("_cookie")
	if err == nil {
		sess = models.Session{UUID: cookie.Value}
		if ok, _ := sess.Check(); !ok {
			err = errors.New("Invalid session")
		}
	}
	return
}

// 解析 HTML 模板（应对需要传入多个模板文件的情况，避免重复编写模板代码）
func parseTemplateFiles(filenames ...string) (t *template.Template) {
	var files []string
	t = template.New("layout")
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("views/%s.html", file))
	}
	t = template.Must(t.ParseFiles(files...))
	return
}

// 生成响应 HTML
func generateHTML(writer http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		// files = append(files, fmt.Sprintf("views/%s.html", file))
		files = append(files, fmt.Sprintf("views/%s/%s.html", config.ViperConfig.App.Language, file))
	}

	// 新增一个格式化日期时间的函数 formatDate, 然后在 generateHTML 方法中将这个函数通过
	// template.FuncMap 组装后再通过 Funcs 方法应用到视图模板中, 这样就可以在所有视图模板
	// 中通过 fdate 别名来调用 formatDate 函数了
	funcMap := template.FuncMap{"fdate": formatDate}
	t := template.New("layout").Funcs(funcMap)
	templates := template.Must(t.ParseFiles(files...))
	templates.ExecuteTemplate(writer, "layout", data)
}

// Version 返回版本号
func Version() string {
	return "0.1"
}

func info(args ...interface{}) {
	logger.SetPrefix("INFO ")
	logger.Println(args...)
}

// 为什么不命名为 error？避免和 error 类型重名
func danger(args ...interface{}) {
	logger.SetPrefix("ERROR ")
	logger.Println(args...)
}

func warning(args ...interface{}) {
	logger.SetPrefix("WARNING ")
	logger.Println(args...)
}

// 异常处理统一重定向到错误页面
func errorMessage(writer http.ResponseWriter, request *http.Request, msg string) {
	url := []string{"/err?msg=", msg}
	http.Redirect(writer, request, strings.Join(url, ""), 302)
}

// 日期格式化辅助函数
func formatDate(t time.Time) string {
	datetime := "2006-01-02 15:04:05"
	return t.Format(datetime)
}
