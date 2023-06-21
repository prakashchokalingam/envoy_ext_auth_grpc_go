package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("path:", r.URL.Path)
		w.Write([]byte(r.Header["X-Custom-Header"][0]))
	})

	fmt.Println("Go simple http server on port 3002")
	http.ListenAndServe(":3002", nil)
}
