package main

import (
	"sync"
	"net/http"
	"fmt"
	"log"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler1)
	//http.HandleFunc("/image", handler3)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("0.0.0.0:8000", nil))
}

func handler1(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request ) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

func handler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %sn", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}

//func handler3(w http.ResponseWriter, r *http.Request) {
//	Lissajous(w)
//}
