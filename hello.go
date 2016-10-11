package main

import (
	"fmt"
	"net/http"
	"time"
)

func handler(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "%v", time.Now())
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
