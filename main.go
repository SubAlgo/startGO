package main

import "net/http"
import "log"

//Mutiplexer (MUX)
func mux(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		indexHandler(w, r)
	case "/about":
		aboutHandler(w, r)
	case "/login":
		loginHandler(w, r)
	default:
		http.NotFound(w, r)
		//notFoudHandler(w, r)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Index Page"))
}
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("About Page"))
}
func loginHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Login Page"))
}
func notFoudHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("404 Page Not Foud1"))
}

func main() {
	h := http.HandlerFunc(mux)
	err := http.ListenAndServe(":8080", h)
	log.Println(err)
}
