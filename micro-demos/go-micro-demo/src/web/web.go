package main

import (
	"encoding/json"
	"go-demos/micro-demos/go-micro-demo/src/share/config"
	"go-demos/micro-demos/go-micro-demo/src/share/pb"
	"log"
	"net/http"
	"strings"

	"github.com/micro/go-web"

	"golang.org/x/net/context"

	"github.com/micro/go-grpc/client"
)

var (
	CORS = map[string]bool{"*": true}
)

func handleRPC(w http.ResponseWriter, r *http.Request) {
	log.Println("handleRPC ... ")

	// 跨域处理
	if r.Method == "OPTIONS" {
		if origin := r.Header.Get("Origin"); CORS[origin] {
			w.Header().Set("Access-Control-Allow-Origin", origin)
		} else if len(origin) > 0 && CORS["*"] {
			w.Header().Set("Access-Control-Allow-Origin", origin)
		}

		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-Token, X-Client")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		return
	}
	defer r.Body.Close()

	ct := r.Header.Get("Content-Type")
	if idx := strings.IndexRune(ct, ';'); idx >= 0 {
		ct = ct[:idx]
	}

	switch ct {
	case "application/json":
		handleJSONRPC(w, r)
	default:
		if r.Method != "POST" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		handleProtobufRPC(w, r)
	}
}

// json请求的处理
func handleJSONRPC(w http.ResponseWriter, r *http.Request) {
	serviceName, methodName := urlToPath(r.URL.String())
	log.Println("现在进来的请求是 ： " + serviceName + "  " + methodName)
	if serviceName == "" || methodName == "" {
		w.Write([]byte(`<html><body><h1>` + "请求的url有异常..." + `</h1></body></html>`))
		return
	}

	decode := json.NewDecoder(r.Body)
	//根据serviceName请求分发，这里应该有更好的解决方案，但是目前还没有找到
	if serviceName == "user" {
		switch methodName {
		case "selectUser":
			cl := pb.NewUserService(config.Namespace+"user", client.DefaultClient)
			var p pb.SelectUserReq
			err := decode.Decode(&p)
			if err != nil {
				log.Fatalln(err)
			}
			rsp, err := cl.SelectUser(context.Background(), &p)
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
			w.Write([]byte(`<html><body><h1>` + rsp.String() + `</h1></body></html>`))
			return

		case "insertUser":
			cl := pb.NewUserService(config.Namespace+"user", client.DefaultClient)
			var p pb.InsertUserReq
			err := decode.Decode(&p)
			if err != nil {
				log.Fatalln(err)
			}
			rsp, err := cl.InsertUser(context.Background(), &p)
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
			w.Write([]byte(`<html><body><h1>` + rsp.String() + `</h1></body></html>`))
			return

		case "modifyUser":
			cl := pb.NewUserService(config.Namespace+"user", client.DefaultClient)
			var p pb.ModifyUserReq
			err := decode.Decode(&p)
			if err != nil {
				log.Fatalln(err)
			}
			rsp, err := cl.ModifyUser(context.Background(), &p)
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
			w.Write([]byte(`<html><body><h1>` + rsp.String() + `</h1></body></html>`))
			return

		case "deleteUser":
			cl := pb.NewUserService(config.Namespace+"user", client.DefaultClient)
			var p pb.DeletetUserReq
			err := decode.Decode(&p)
			if err != nil {
				log.Fatalln(err)
			}
			rsp, err := cl.DeletetUser(context.Background(), &p)
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
			w.Write([]byte(`<html><body><h1>` + rsp.String() + `</h1></body></html>`))

		default:
			w.Write([]byte(`<html><body><h1>` + "查询不到该method..." + `</h1></body></html>`))
			return
		}
	}
}

//这个待定
func handleProtobufRPC(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "handleProtobufRPC", http.StatusMethodNotAllowed)
	return
}

// 将/xxx/xxx的Url转为 xxx xxx返回
func urlToPath(url string) (string, string) {

	parts := strings.Split(url, "/")
	if len(parts) >= 3 {
		return parts[1], parts[2]
	}

	return "", ""
}

func main() {
	// mux := http.NewServeMux()
	// mux.HandleFunc("/", handleRPC)
	// log.Println("Listen on :8082")
	// http.ListenAndServe(":8082", mux)

	service := web.NewService(
		web.Name(config.ConfigPrefix + ".web"),
	)

	service.HandleFunc("/", handleRPC)

	if err := service.Init(); err != nil {
		log.Fatal(err)
	}

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
