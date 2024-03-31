package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

// NlpResp 返回信息
type NlpResp struct {
	// bool state = 1; // 状态
	MatchID    int      `json:"matchId"`
	Audio      string   `json:"audio"`
	Hai        string   `json:"hai"`
	Text       string   `json:"text"`
	RemainText []string `json:"remainText"`
}

func httpServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		flusher, ok := w.(http.Flusher)
		if !ok {
			panic("expected http.ResponseWriter to be an http.Flusher")
		}
		w.Header().Set("X-Content-Type-Options", "nosniff")
		for i := 1; i <= 20; i++ {
			tmp := NlpResp{
				MatchID: i,
				Audio:   "aaaaaaaaaaaaaaaaaaaaa",
				Hai:     "hhhhhhhhhhhhhhhhhhhh",
				Text:    "test",
			}
			fmt.Fprintf(w, GetJSONStr(tmp, false)+"\n")
			flusher.Flush() // Trigger "chunked" encoding and send a chunk...
			time.Sleep(1 * time.Second)
		}
	})

	log.Print("Listening on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func httpClient() {
	resp, err := http.Get("http://localhost:8080")
	if err != nil {
		// handle error
	}

	defer resp.Body.Close()

	// chunkedReader := httputil.NewChunkedReader(resp.Body)

	buf := make([]byte, 40960)
	for {
		n, err := resp.Body.Read(buf)
		fmt.Println(n, err)
		if n != 0 || err != io.EOF { // simplified
			fmt.Println(string(buf[:n]))
			time.Sleep(1 * time.Second)
		} else {
			break
		}
	}
	fmt.Println("client exit")
}

func GetJSONStr(obj interface{}, isFormat bool) string {
	var b []byte
	if isFormat {
		b, _ = json.MarshalIndent(obj, "", "     ")
	} else {
		b, _ = json.Marshal(obj)
	}
	return string(b)
}

func main() {
	go httpServer()
	go httpClient()

	time.Sleep(25 * time.Second)
}
