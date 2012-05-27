package main

import (
	"code.google.com/p/gorilla/mux"
	"fmt"
	"github.com/hoisie/mustache"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", index)
	r.HandleFunc("/must", must)
	http.Handle("/", r)
	http.ListenAndServe("0.0.0.0:3000", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "hello, golang!")
}

func must(res http.ResponseWriter, req *http.Request) {
	data := mustache.Render("Hello, {{name}}!", map[string]string{"name": "mustache"})
	fmt.Fprintf(res, data)
}
