package n14

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
)

func Exec() {
	firstServeMux := http.NewServeMux()
	firstServeMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if rand.Int()%2 == 0 {
			io.WriteString(w, "pong")
		}
	})
	secondServeMux := http.NewServeMux()
	secondServeMux.HandleFunc("/api/general", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.RequestURI)
		io.WriteString(w, "general")
	})

	firstServeMux.Handle("/api/", secondServeMux)

	// ***

	http.ListenAndServe(
		":42521", firstServeMux,
	)
}
