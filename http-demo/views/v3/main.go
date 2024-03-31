package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func checkErr(str string, err error) {
	if err != nil {
		log.Println(str, err)
	}
}

// 存放用户数据
type UserData struct {
	Name string
	Text string
}

// 渲染页面并输出
func renderHTML(w http.ResponseWriter, file string, data interface{}) {
	// 获取页面内容
	t, err := template.New(file).ParseFiles("views/" + file)
	checkErr("template.New", err)
	// 将页面渲染后反馈给客户端
	t.Execute(w, data)
}

// 写入数据库（返回写入后的数据）
func writeData(userData *UserData) string {
	// 打开数据库
	db, err := sql.Open("sqlite3", "./data.db")
	checkErr("sql.Open", err)
	defer db.Close()

	// 如果数据表不存在则创建（如果存在则跳过）
	db.Exec(`create table data (id integre not null primary key, name text, data string);`)

	var olddata string // 数据库中已存在的数据
	var sqlStmt string // sql 内容

	// 查询用户是否存在，同时读取用户数据
	err = db.QueryRow("select data from data where name = ?", userData.Name).Scan(&olddata)
	if err != nil { // 用户不存在
		sqlStmt = "insert into data(id, data, name) values(1,?,?)" // 添加数据
	} else { // 用户存在
		sqlStmt = "update data set data = ? where name = ?" // 更新数据
		// 如果 data 为空，则删除用户
		if len(userData.Text) == 0 {
			sqlStmt = "delete from data where data >= ? and name = ?" // 删除字段
		} else {
			// 否则将 data 追加到数据库
			userData.Text = olddata + "\n" + userData.Text
		}
	}

	// 准备 SQL
	stmt, err := db.Prepare(sqlStmt)
	checkErr("db.Prepare", err)
	defer stmt.Close()

	// 执行 SQL
	_, err = stmt.Exec(userData.Text, userData.Name)
	checkErr("stmt.Exec", err)
	return userData.Text
}

// 处理主页请求
func index(w http.ResponseWriter, r *http.Request) {
	// 渲染页面并输出
	renderHTML(w, "index.html", "no data")
}

// 处理用户提交的数据
func page(w http.ResponseWriter, r *http.Request) {
	// 我们规定必须通过 POST 提交数据
	if r.Method == "POST" {
		// 解析客户端请求的信息
		if err := r.ParseForm(); err != nil {
			log.Println("Handler:page:ParseForm: ", err)
		}

		// 获取客户端输入的内容
		u := UserData{}
		u.Name = r.Form.Get("username")
		u.Text = r.Form.Get("usertext")

		// 写入数据库，同时获取处理后的数据
		u.Text = writeData(&u)

		// 渲染页面并输出
		renderHTML(w, "page.html", u)
	} else {
		// 如果不是通过 POST 提交的数据，则将页面重定向到主页
		renderHTML(w, "redirect.html", "/")
	}
}

func main() {
	http.HandleFunc("/", index)              // 设置访问的路由
	http.HandleFunc("/page", page)           // 设置访问的路由
	err := http.ListenAndServe(":9090", nil) // 设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
