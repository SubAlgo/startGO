package main

import (
	"log"
	"net/http"
)

func main() {
	err := http.ListenAndServe(":8080", &indexHandler{})
	log.Println(err)
}

type indexHandler struct{}

func (*indexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, SubAlgo"))
}

//28
