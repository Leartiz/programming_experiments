package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func topicHandler(w http.ResponseWriter, r *http.Request) {
	topicIDStr := r.URL.Query().Get("id")
	if len(topicIDStr) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "bad request")
		return
	}
	topicID, err := strconv.Atoi(topicIDStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "topic id must be a number")
		return
	}
	io.WriteString(w, fmt.Sprintf("topic has id '%v'", topicID))
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(r.URL.Path))
}

func main() {
	http.HandleFunc("/test/", func(w http.ResponseWriter, r *http.Request) {
		uriParts := strings.Split(r.RequestURI, "/")
		if len(uriParts) != 3 {
			http.NotFound(w, r)
			return
		}
		switch {
		case strings.HasPrefix(uriParts[2], "topic"):
			topicHandler(w, r)
		case strings.HasPrefix(uriParts[2], "user"):
			userHandler(w, r)
		default:
			http.NotFound(w, r)
		}
	})
	http.ListenAndServe(":42521", nil)
}
