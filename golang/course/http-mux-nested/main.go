package main

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
)

func main() {
	first := http.NewServeMux()
	first.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if rand.Int()%2 == 0 {
			io.WriteString(w, "pong")
		}
	})

	second := http.NewServeMux()
	second.HandleFunc("/api/general", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.RequestURI)
		io.WriteString(w, "general")
	})

	first.Handle("/api/", second)
	http.ListenAndServe(":42521", first)
}
