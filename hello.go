package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func Hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Fully-integrated-Hello world! time:", time.Now())
}

func main() {
	// http.HandleFunc("/", handler)
	// http.ListenAndServe(":8080", nil)

	r := mux.NewRouter()
	r.HandleFunc("/", Hello)
	http.Handle("/", r)
	fmt.Println("Starting up on 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
