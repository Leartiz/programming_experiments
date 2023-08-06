package n2

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

type key int

const (
	userIdCtx key = 0
)

func Exec() {
	http.HandleFunc("/", handle)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func handle(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("User-Id")
	ctx := context.WithValue(r.Context(), userIdCtx, id)
	result := processLongTask(ctx)
	w.Write([]byte(result))
}

func processLongTask(ctx context.Context) string {
	id := ctx.Value(userIdCtx)

	// ***

	select {
	case <-time.After(2 * time.Second):
		return fmt.Sprintf("done processing id %v\n", id)
	case <-ctx.Done():
		log.Println("request canceled")
		return ""
	}
}
