package handlers

import (
	"fmt"
	"net/http"

	"chitchat/models"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

// PostThread 在指定群组下创建新主题
// POST /thread/post
func PostThread(writer http.ResponseWriter, request *http.Request) {
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
		body := request.PostFormValue("body")
		uuid := request.PostFormValue("uuid")
		thread, err := models.ThreadByUUID(uuid)
		if err != nil {
			// log.Println("Cannot read thread")
			// danger(err, "Cannot read thread")
			// errorMessage(writer, request, "Cannot read thread")

			msg := localizer.MustLocalize(&i18n.LocalizeConfig{
				MessageID: "thread_not_found",
			})
			errorMessage(writer, request, msg)
		}
		if _, err := user.CreatePost(thread, body); err != nil {
			// log.Println("Cannot create post")
			danger(err, "Cannot create post")
		}
		url := fmt.Sprint("/thread/read?id=", uuid)
		http.Redirect(writer, request, url, 302)
	}
}
