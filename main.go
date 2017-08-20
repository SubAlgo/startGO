package main

import (
	"fmt"
	"log"
	"net/http"
)

/*
func main() {
	err := http.ListenAndServe(":8080", &indexHandler1{})
	log.Println(err)
}

//--รูปแบบ Handler function > ServeHTTP(w http.ResponseWriter, r *http.Request) {}--
type indexHandler1 struct{}
func (*indexHandler1) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, SubAlgo"))
}
*/

//วิธีนี้จะเขียน func ขึ้นมาแล้วโยนเข้า http.HandlerFunc() เพื่อแปลงเป็น HandlerFunc
func main() {
	h := http.HandlerFunc(indexHandler)
	err := http.ListenAndServe(":8080", h)
	log.Println(err)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	switch r.URL.Path {
	case "/":
		w.Write([]byte("Index Page"))
	case "/about":
		w.Write([]byte("about"))
	case "/login":
		w.Write([]byte("Login"))
	default:
		w.Write([]byte("404 Page Not Found"))
	}
}

//39
