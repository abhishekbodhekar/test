package main

import (
	"net/http"

	mux "github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", Handler)
	r.HandleFunc("/suraj", Handler2)
	http.ListenAndServe(":7070", r)
}

func Handler(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("hi"))

}

func Handler2(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("suraj"))

}
