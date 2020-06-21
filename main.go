package main

import (
	"fmt"
	"net/http"
)

type FooHanlder struct {
}

func (f FooHanlder) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello foo")
}
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello world!!")
	})
	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello bar")
	})
	http.Handle("/foo", FooHanlder{})
	http.ListenAndServe(":3000", nil)
}
