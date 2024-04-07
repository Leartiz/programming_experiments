package n12

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func FuncTopic(w http.ResponseWriter, r *http.Request) {

	values := r.URL.Query()
	topicIdStr := values.Get("id")

	// ***

	if len(topicIdStr) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "bad request")
		return
	}

	topicId, err := strconv.Atoi(topicIdStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "topic id must be a number")
		return
	}

	// ***

	io.WriteString(
		w, fmt.Sprintf(
			"topic has id '%v'", topicId))
}

func FuncUser(w http.ResponseWriter, r *http.Request) {
	// TODO:

	w.Write([]byte(r.URL.Path))
}

func Exec() {
	http.HandleFunc("/test/", func(w http.ResponseWriter, r *http.Request) {
		uriParts := strings.Split(r.RequestURI, "/")
		if len(uriParts) != 3 {
			http.NotFound(w, r)
			return
		}

		// ***

		if strings.HasPrefix(uriParts[2], "topic") {
			FuncTopic(w, r)

		} else if strings.HasPrefix(uriParts[2], "user") {
			FuncUser(w, r)

		} else {
			http.NotFound(w, r)
		}
	})

	http.ListenAndServe(":42521", nil)
}
