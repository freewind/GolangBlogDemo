package main

import (
	"code.google.com/p/gorilla/mux"
	"fmt"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", index)
	http.Handle("/", r)
	http.ListenAndServe("0.0.0.0:3000", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Hello, golang!")
}
