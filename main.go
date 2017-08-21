package main

import "net/http"
import "log"

//Mutiplexer style2 (MUX)
//MUX วิธีนี้โดยค่าเริ่มต้นถ้าไม่มีurl จะไปที่ Index
//Code ที่ commend เป็นแบบวิธีที่1

/*
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
*/

/*
func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Index Page"))
}
*/

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
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
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/about", aboutHandler)
	mux.HandleFunc("/login", loginHandler)

	err := http.ListenAndServe(":8080", mux)
	//h := http.HandlerFunc(mux)
	//err := http.ListenAndServe(":8080", h)
	log.Println(err)
}
