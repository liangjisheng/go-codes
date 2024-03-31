package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var httpRequestCount = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_request_count",
		Help: "http request count",
	},
	[]string{"endpoint"},
)

var orderNum = prometheus.NewGauge(
	prometheus.GaugeOpts{
		Name: "order_num",
		Help: "order num"},
)

var httpRequestDuration = prometheus.NewSummaryVec(
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

func init() {
	prometheus.MustRegister(httpRequestCount)
	prometheus.MustRegister(orderNum)
	prometheus.MustRegister(httpRequestDuration)
}

//curl "http://127.0.0.1:8080/index"
//curl "http://127.0.0.1:8080/metrics"

func main() {
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/index", index)
	fmt.Println("server start 127.0.0.1:8080")

	go func() {
		rand.Seed(time.Now().UnixNano())
		for {
			start := time.Now()
			tmp := 1 + rand.Intn(1000)
			time.Sleep(time.Duration(tmp) * time.Millisecond)
			elapsed := float64(time.Since(start).Milliseconds())
			log.Println("elapsed:", elapsed)
			httpRequestDuration.WithLabelValues("/index").Observe(elapsed)
		}
	}()

	err := http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	httpRequestCount.WithLabelValues(r.URL.Path).Inc()
	n := rand.Intn(100)
	if n >= 90 {
		orderNum.Dec()
	} else {
		orderNum.Inc()
	}

	w.Write([]byte("ok"))
}
