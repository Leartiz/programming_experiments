package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func addRandomValueToHist(hist *prometheus.HistogramVec) {

}

func main() {

	// hist

	buckets := []float64{0.25, 1, 5, 10}

	hist := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "my_hist",
			Buckets: buckets,
		},
		[]string{
			"label_1",
		},
	)
	prometheus.MustRegister(hist)

	// counter

	counter := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "my_counter",
		},
		[]string{
			"label_1",
		},
	)
	prometheus.MustRegister(counter)
	counter.WithLabelValues()

	promHandler := promhttp.Handler()
	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(".")
		promHandler.ServeHTTP(w, r)
	})

	// ***

	go func() {
		step := float64(rand.Intn(50))

		for {
			select {
			case <-time.After(1 * time.Second):
				counter.WithLabelValues("11", "12").Add(step)
			}
		}
	}()

	// ***

	go func() {
		for {
			select {
			case <-time.After(1 * time.Second):
				addRandomValueToHist(hist)
			}
		}
	}()

	http.ListenAndServe(":8088", nil)
}
