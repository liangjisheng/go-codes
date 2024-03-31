package webservice

import (
	"encoding/json"
	"fmt"
	"go-demos/projects/user-register-login/database"
	"go-demos/projects/user-register-login/model"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"text/template"
)

// CustomMux ...
type CustomMux struct{}

func (mux *CustomMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		index(w, r)
		return
	case "/home":
		homePage(w, r)
		return
	case "/login":
		login(w, r)
		return
	case "/add":
		renderAdd(w, r)
		return
	case "/user/insert":
		insertUser(w, r)
		return
	case "/user/select":
		queryByID(w, r)
		return
	case "/upload/picture":
		uploadPicture(w, r)
		return
	case "/do/upload":
		doUploadAction(w, r)
		return
	case "/json/param":
		jsonHandler(w, r)
		return
	default:
		http.NotFound(w, r)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	log.Println("-----> index handler start <-----")
	t, err := template.ParseFiles("template/login.html")
	checkErr(err)
	t.Execute(w, nil)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	log.Println("-----> home page handler start <-----")
	t, err := template.ParseFiles("template/homePage.html")
	checkErr(err)
	t.Execute(w, "这是网站首页")
}

func login(w http.ResponseWriter, r *http.Request) {
	log.Println("-----> login handler start <-----")
	if r.Method == "Get" {
		t, err := template.ParseFiles("template/login.html")
		checkErr(err)
		t.Execute(w, r)
	} else {
		username := r.FormValue("username")
		password := r.Form["password"]
		fmt.Println("username: ", username)
		fmt.Println("password: ", password)
		http.Redirect(w, r, "/home", 302)
	}
}

func renderAdd(w http.ResponseWriter, r *http.Request) {
	// 跳转新加用户页面
	log.Println("-----> render to insert page handler start <-----")
	t, err := template.ParseFiles("template/user/addUser.html")
	checkErr(err)
	t.Execute(w, nil)
}

func insertUser(w http.ResponseWriter, r *http.Request) {
	log.Println("-----> insert new user handler start <-----")
	// 获取表单参数
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	age := r.PostFormValue("age")
	mobile := r.PostFormValue("mobile")
	address := r.PostFormValue("address")
	//省略校验逻辑....

	// 插入数据库
	db := database.GetConn()
	insert := "insert into t_user (username,password,address,age,mobile,status) values($1,$2,$3,$4,$5,$6)"
	rs, err := db.Exec(insert, username, password, address, age, mobile, 1)
	checkErr(err)
	row, err := rs.RowsAffected()
	checkErr(err)
	fmt.Println("----> insert account = " + strconv.FormatInt(row, 10) + " <----")
	if row != 1 {
		log.Fatal("----> error occurred <----")
	} else {
		fmt.Println("----> insert user to database succeed <----")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("insert succeed."))
	}

	checkErr(db.Close())
}

func queryByID(w http.ResponseWriter, r *http.Request) {
	log.Println("-----> query by id handler start <-----")

	fmt.Println(r.URL.RawQuery)
	params := r.URL.Query()
	id := params["id"][0]

	fmt.Println(id)
	db := database.GetConn()
	query := "select username,sex,age,address,mobile,role from t_user where id = $1"
	rows, err := db.Query(query, id)
	checkErr(err)

	var user model.User
	for rows.Next() {
		err = rows.Scan(&user.Username, &user.Sex, &user.Age, &user.Address, &user.Mobile, &user.Role)
		checkErr(err)

		w.Header().Set("content-type", "application/json")
		json.NewEncoder(w).Encode(user)
	}
	checkErr(db.Close())
}

func uploadPicture(w http.ResponseWriter, r *http.Request) {
	log.Println("-----> forward to picture file upload handler start <-----")
	t, err := template.ParseFiles("template/upload.html")
	checkErr(err)
	t.Execute(w, nil)
}

func doUploadAction(w http.ResponseWriter, r *http.Request) {
	log.Println("-----> upload picture handler start <-----")
	if r.Method == "GET" {
		t, _ := template.ParseFiles("template/error.html")
		t.Execute(w, r)
	} else {
		file1, head, err := r.FormFile("picture1")
		checkErr(err)
		fmt.Println("----> filename: " + head.Filename + " <----")
		defer file1.Close()

		fileBytes, err := ioutil.ReadAll(file1)
		checkErr(err)
		// 图片类型
		fileType := http.DetectContentType(fileBytes)
		fmt.Println(fileType)

		// 创建存放图片文件夹
		dest := "upload/"
		exist := dirExist(dest)
		if exist {
			// fmt.Println("----> directory has exist <----")
		} else {
			err := os.Mkdir(dest, os.ModePerm)
			checkErr(err)
		}

		newFile, err := os.Create(dest + head.Filename)
		checkErr(err)
		defer newFile.Close()
		len, err := newFile.Write(fileBytes)
		if err != nil {
			fmt.Println("----> error occurred while write file to disk <----")
		}

		fmt.Println(len)
		http.ServeFile(w, r, dest+head.Filename)
		//writer.Header().Set("content-type","application/json")
		//json.NewEncoder(writer).Encode("imageURL:" + dest + head.Filename)
	}
}

func jsonHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("-----> request json param handler start <-----")
	bytes, err := ioutil.ReadAll(r.Body)
	checkErr(err)
	var user model.User
	err = json.Unmarshal(bytes, &user)
	checkErr(err)

	// insert into database
	db := database.GetConn()
	insert := "insert into t_user (username,password,address,age,mobile,sex,status) values($1,$2,$3,$4,$5,$6,$7)"
	rs, _ := db.Exec(insert, user.Username, user.Password, user.Address, user.Age, user.Mobile, user.Sex, 1)
	count, _ := rs.RowsAffected()

	var res model.ComRes
	if count != 1 {
		res.Code = "0002"
		res.Success = false
		res.Message = "insert to database failed"
	} else {
		res.Code = "0001"
		res.Success = true
		res.Message = "insert to database success"
	}
	// 返回自定义的响应结果
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func dirExist(s string) bool {
	var exist = true
	if _, err := os.Stat(s); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func checkErr(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
