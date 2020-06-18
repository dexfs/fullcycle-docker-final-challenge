package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s", greeting(r.URL.Path[1:]))
	})

	http.ListenAndServe(":8123", nil)
}
