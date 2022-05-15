package main

import (
	"net/http"
)

func Decorate(h http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("this is "))
		h(writer, request)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

func main() {
	http.HandleFunc("/", Decorate(index))
	http.ListenAndServe(":8181", nil)
}
