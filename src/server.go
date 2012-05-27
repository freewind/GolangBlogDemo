package main

import (
	"code.google.com/p/gorilla/mux"
	"fmt"
	mus "github.com/hoisie/mustache"
	"html/template"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", index)
	r.HandleFunc("/mustache", mustache)
	r.HandleFunc("/tmpl", tmpl)
	http.Handle("/", r)
	http.ListenAndServe(":3000", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "hello, golang!")
}

func tmpl(res http.ResponseWriter, req *http.Request) {
	t := template.New("test")
	t, _ = t.Parse("hello, {{.Name}}!")
	t.Execute(res, map[string]string{"Name": "html/template"})
}

func mustache(res http.ResponseWriter, req *http.Request) {
	data := mus.Render("Hello, {{name}}!", map[string]string{"name": "mustache"})
	fmt.Fprintf(res, data)
}
