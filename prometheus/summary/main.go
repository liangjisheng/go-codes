package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"math/rand"
	"net/http"
	"time"
)

//summary 统计分位数, Objectives map 的 value 好像没啥作用, key 代表要统计的分位数

var (
	summary = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name: "http_request_duration",
			Help: "http request duration",
			Objectives: map[float64]float64{
				0.2: 0,
				0.4: 0,
				0.5: 0,
				0.8: 0,
				0.9: 0,
			},
		},
		[]string{"endpoint"},
	)
)

func main() {
	prometheus.MustRegister(summary)

	go func() {
		rand.Seed(time.Now().UnixNano())
		for {
			tmp := float64(rand.Intn(1000))
			time.Sleep(time.Millisecond)
			summary.WithLabelValues("/index").Observe(tmp)
		}
	}()

	http.Handle("/metrics", promhttp.Handler())

	fmt.Println("server start 127.0.0.1:8080")
	err := http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
