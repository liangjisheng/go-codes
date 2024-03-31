package main

import (
	"container/list"
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

/*
session 管理设计
·全局session 管理器
·保证sessionid 的全局唯一性
·为每个客户关联一个session
·session 的存储（可以存储到内存、文件、数据库等）
·session 过期处理
*/

var globalSessions *Manager

func init() {
	provider.sessions = make(map[string]*list.Element, 0)
	Register("memory", provider)

	var err error
	globalSessions, err = NewManager("memory", "gosessionid", 8)
	if err != nil {
		log.Println("new manager err:", err)
		os.Exit(-1)
		return
	}
	if nil == globalSessions {
		log.Println("global session is nil")
		os.Exit(-1)
		return
	}

	fmt.Printf("globalSessions:%+v \n", globalSessions)

	// 异步使用Gc 销毁session
	go globalSessions.GC()
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello testsession.go")
}

func testSession(w http.ResponseWriter, r *http.Request) {
	sess := globalSessions.SessionStart(w, r)
	r.ParseForm()

	if r.Method == "GET" {
		t, _ := template.New("foo").Parse(`{{define "T"}}hello,{{.}}!{{end}}
username: {{.UserName}}
token: {{.Token}}
`)
		err := t.ExecuteTemplate(w, "T", template.HTML("<script>alert('you have been testSession'</script>"))
		if err != nil {
			log.Println("exec template err:", err)
			fmt.Fprintf(w, "exec template err: %+v\n", err)
			return
		}

		w.Header().Set("Content-Type", "text/html")
		//创建唯一的token ,并写入session 里
		token := uniqueToken()
		sess.Set("token", token)

		// 创建一个数据对象
		type Result struct {
			UserName interface{}
			Token    interface{}
		}
		res := Result{
			UserName: sess.Get("username"),
			Token:    token,
		}
		//数据对象输出到模板中
		t.Execute(w, res)
	} else {
		sessToken := sess.Get("token")
		// token := r.Form["token"] // 通常token 是放到表单中提交过来的
		token := r.Form.Get("token")
		// token := uniqueToken()
		fmt.Println("session token:", sessToken)
		fmt.Println("token:", token)
		if sessToken != token {
			//提示登录
			fmt.Fprintf(w, "token无效\n")
			return
		}
		sess.Set("username", "testsession")
		http.Redirect(w, r, "/", 302)
	}
}

func countSession(w http.ResponseWriter, r *http.Request) {
	sess := globalSessions.SessionStart(w, r)
	createtime := sess.Get("createtime")
	if createtime == nil {
		sess.Set("createtime", time.Now().Unix())
	} else if (createtime.(int64) + 3600) < (time.Now().Unix()) {
		//间隔生成新的SID
		//每3600秒就刷新一次session ,用户需重新登录
		globalSessions.SessionDestroy(w, r)
		sess = globalSessions.SessionStart(w, r)
	}
	ct := sess.Get("countnum")
	if ct == nil {
		sess.Set("countnum", 1)
	} else {
		sess.Set("countnum", (ct.(int) + 1))
	}
	t, _ := template.New("foo").Parse(`{{define "T"}}hello,{{.}}!{{end}}
refresh count：{{.}}
`)
	err := t.ExecuteTemplate(w, "T", template.HTML("<script>alert('you have been countSession')</script>"))
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "text/html")
	t.Execute(w, sess.Get("contnum"))
}

func uniqueToken() string {
	h := md5.New()
	salt := "sessionss%^7&888"
	io.WriteString(h, salt+time.Now().String())
	token := fmt.Sprintf("%x", h.Sum(nil))
	return token
}

func main() {
	http.HandleFunc("/", sayhelloName)
	http.HandleFunc("/test", testSession)
	http.HandleFunc("/count", countSession)

	err := http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe：", err)
	}
}
