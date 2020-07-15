package handlers

import (
	"net/http"

	"chitchat/models"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

// NewThread 创建群组页面
// GET /threads/new
func NewThread(writer http.ResponseWriter, request *http.Request) {
	_, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		generateHTML(writer, nil, "layout", "auth.navbar", "new.thread")
	}
}

// CreateThread 执行群组创建逻辑
// POST /thread/create
func CreateThread(writer http.ResponseWriter, request *http.Request) {
	sess, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		err = request.ParseForm()
		if err != nil {
			// log.Println("Cannot parse form")
			danger(err, "Cannot parse form")
		}
		user, err := sess.User()
		if err != nil {
			// log.Println("Cannot get user from session")
			danger(err, "Cannot get user from session")
		}
		topic := request.PostFormValue("topic")
		if _, err := user.CreateThread(topic); err != nil {
			// log.Println("Cannot create thread")
			danger(err, "Cannot create thread")
		}
		http.Redirect(writer, request, "/", 302)
	}
}

// ReadThread 通过 ID 渲染指定群组页面
// GET /thread/read
func ReadThread(writer http.ResponseWriter, request *http.Request) {
	vals := request.URL.Query()
	uuid := vals.Get("id")
	thread, err := models.ThreadByUUID(uuid)
	if err != nil {
		// log.Println("Cannot read thread")
		// danger(err, "Cannot read thread")
		// errorMessage(writer, request, "Cannot read thread")

		msg := localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: "thread_not_found",
		})
		errorMessage(writer, request, msg)
	} else {
		_, err := session(writer, request)
		if err != nil {
			generateHTML(writer, &thread, "layout", "navbar", "thread")
		} else {
			generateHTML(writer, &thread, "layout", "auth.navbar", "auth.thread")
		}
	}
}
