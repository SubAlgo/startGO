package main

//Basic Middleware
import (
	"log"
	"net/http"
	"time"
)

func main() {
	h := logger(http.HandlerFunc(indexHandler))
	err := http.ListenAndServe(":8080", h)
	log.Println(err)
}
func logger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//http://localhost:8080/about?hello=111
		log.Printf("r.URL.Path: %s", r.URL.Path)       //r.URL.Path: /about
		log.Printf("r.RequestURI: %s", r.RequestURI)   //r.RequestURI: /about?hello=111
		log.Printf("r.URL.Query(): %s", r.URL.Query()) //r.URL.Query(): map[hello:[111]]
		t := time.Now()
		h.ServeHTTP(w, r)
		diff := time.Now().Sub(t)
		log.Printf("path: %s, time: %d ms", r.URL.Path, diff/time.Microsecond)
	})
}
func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Index Page1"))
}
