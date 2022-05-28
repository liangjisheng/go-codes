package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"math/rand"
	"net/http"
	"time"
)

var (
	gaugeVec = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "inflight",
			Help: "Number of jobs inflight",
		},
		[]string{"type"},
	)
)

func main() {
	prometheus.MustRegister(gaugeVec)

	go func() {
		rand.Seed(time.Now().UnixNano())
		for {
			tmp := float64(rand.Intn(100))
			gaugeVec.WithLabelValues("type1").Add(tmp)
			time.Sleep(time.Millisecond)

			tmp = float64(rand.Intn(100))
			gaugeVec.WithLabelValues("type1").Sub(tmp)
			time.Sleep(time.Millisecond)
		}
	}()

	http.Handle("/metrics", promhttp.Handler())

	fmt.Println("server start 127.0.0.1:8080")
	err := http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
