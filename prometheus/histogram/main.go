package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

//这里的 histogram 是累积直方图, 而不是非累积直方图
//仅仅 histogram 的指标名有后缀 Name_bucket
//_sum、_count 也有这2个总数指标

var (
	histogramVec = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "process_time_seconds",
			Help:    "Amount of time spent processing jobs",
			Buckets: prometheus.LinearBuckets(0, 10, 10),
		},
		[]string{"worker_id", "type"},
	)
)

func main() {
	prometheus.MustRegister(histogramVec)

	go func() {
		rand.Seed(time.Now().UnixNano())
		for i := 0; i < 1000; i++ {
			tmp := float64(rand.Intn(100))
			histogramVec.WithLabelValues("worker1", "type1").Observe(tmp)
		}
	}()

	http.Handle("/metrics", promhttp.Handler())

	fmt.Println("server start 127.0.0.1:8080")
	err := http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
