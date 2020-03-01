package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", hander)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("www.mygo.com:8081", nil))
}

func hander(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprint(w, "URL.Path = %q\n", r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprint(w, "Count %d\n", count)
	mu.Unlock()
}
