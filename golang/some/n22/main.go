package main

<<<<<<< HEAD
import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	// Serves unicode entities
	r.GET("/json", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"html": "<b>Hello, world!</b>",
		})
	})

	// Serves literal characters
	r.GET("/purejson", func(c *gin.Context) {
		c.PureJSON(200, gin.H{
			"html": "<b>Hello, world!</b>",
		})
	})

	// listen and serve on 0.0.0.0:8080
	r.Run(":8079")
=======
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
>>>>>>> ee62c8224e235c260521116d0b8d874a16de0dc5
}
