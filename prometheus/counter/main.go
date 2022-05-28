package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"time"
)

//curl "http://127.0.0.1:8080/metrics"
//curl "http://127.0.0.1:8080/ping"

var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "myapp_processed_ops_total",
		Help: "The total number of processed events",
	})

	pingCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "ping_request_count",
			Help: "No of request handled by Ping handler",
		},
	)

	totalCounterVec = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "processed_total",
			Help: "Total number of jobs processed by the workers",
		},
		[]string{"worker_id", "type"},
	)
)

func ping(w http.ResponseWriter, req *http.Request) {
	pingCounter.Inc()
	fmt.Fprintf(w, "pong")
}

func main() {
	prometheus.MustRegister(pingCounter)
	prometheus.MustRegister(totalCounterVec)

	go func() {
		for {
			opsProcessed.Inc()
			time.Sleep(time.Second)
		}
	}()

	go func() {
		for {
			totalCounterVec.WithLabelValues("worker1", "type1").Inc()
			time.Sleep(time.Second)
		}
	}()

	log.Println("server listen on :8080")

	//只暴露了默认指标
	http.Handle("/metrics", promhttp.Handler())

	http.HandleFunc("/ping", ping)
	http.ListenAndServe(":8080", nil)
}
