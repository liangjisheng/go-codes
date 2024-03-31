package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

// GetPosts ...
func GetPosts(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "All posts")
}

// Post ...
type Post struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

// AddPost ...
func AddPost(w http.ResponseWriter, r *http.Request) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	// io.WriteString(w, string(body))
	post := Post{}
	json.Unmarshal(body, &post)
	fmt.Fprintf(w, "%#v\n", post)
}

// EditPost ...
func EditPost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	// request 对象上的 Form 可以获取所有请求参数, 包括查询字符串和请求实体
	// log.Fprintln(w, r.Form)

	id := r.Form.Get("id")
	log.Println("post id:", id)
	// 如果只想获取请求实体(即 POST 表单中的数据), 可以通过 PostForm 实现
	log.Println("form data:", r.PostForm)

	// 这两个方法的时候只能获取特定请求数据，不能一次获取所有请求数据
	log.Println("post id:", r.FormValue("id"))
	log.Println("post title:", r.PostFormValue("title"))
	log.Println("post title:", r.PostFormValue("content"))

	// application/x-www-form-urlencoded 仅限于文本字符类数据编码, 不能用于二进制数据编码
	// 而通过表单上传的文件是以二进制流的方式提交到服务器的, 因此不能通过默认的编码格式进行
	// 进行编码, 需要通过专门的 multipart/form-data 编码类型. 这种编码类型同时支持文本字符
	// 和二进制文件, 在具体编码时, 会将表单数据分成多个部分, 每个文件单独占用一个部分, 表单正
	// 文中包含的文本数据占用一个部分

	// 需要在调用 ParseMultipartForm 时传入存储解析后文件的最大内存值(单位是字节).
	// MultipartForm 包含了所有 POST 表单请求字段, 即 PostForm 中的所有内容,
	// 但不包含 URL 查询字符串中的请求参数. MultipartForm 返回的值包含两个部分
	// 一部分是单纯的 POST 请求字段, 我们可以通过 Value 字段来访问它, 另一部分就
	// 是包含文件信息的字典，我么可以通过 File 字段来访问它

	r.ParseMultipartForm(1024 * 1024)
	log.Println("post file:", r.MultipartForm)

	io.WriteString(w, "表单提交成功")
}

// UploadImage ...
func UploadImage(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(1024 * 1024)      // 最大支持 1024 KB, 即 1 M
	name := r.MultipartForm.Value["name"]  // 文件名
	image := r.MultipartForm.File["image"] // 图片文件

	log.Println("图片上传成功: ", name[0])

	file, err := image[0].Open()
	if err == nil {
		data, err := ioutil.ReadAll(file) // 读取二进制文件字节流
		if err == nil {
			fmt.Fprintln(w, string(data)) // 将读取的字节信息输出
		}
	}
}

// UploadImage1 将上传的文件保存到服务器指定目录
func UploadImage1(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(1024 * 1024)
	name := r.MultipartForm.Value["name"][0]
	image := r.MultipartForm.File["image"][0]

	file, err := image.Open()
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	// 读取二进制文件字节流
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	// fmt.Fprintln(w, string(data))   // 将读取的字节信息输出
	// 将文件存储到项目根目录下的 images 子目录
	// 从上传文件中读取文件名并获取文件后缀
	names := strings.Split(image.Filename, ".")
	suffix := names[len(names)-1]
	// 将上传文件名字段值和源文件后缀拼接出新的文件名
	filename := name + "." + suffix
	// 创建这个文件
	newFile, _ := os.Create("images/" + filename)
	defer newFile.Close()
	// 将上传文件的二进制字节信息写入新建的文件
	size, err := newFile.Write(data)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	fmt.Fprintf(w, "图片上传成功, 图片大小: %d 字节\n", size)
}
