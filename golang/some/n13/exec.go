package n13

import (
	"io"
	"net/http"
)

func Exec() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "main")
	})
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "pong")
	})

	http.ListenAndServe(
		":42521", nil)
}
