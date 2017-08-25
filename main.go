package main

import (
	"log"
	"net/http"
)

//Config middleware 1
func main() {
	//http.DefaultServeMux()
	http.HandleFunc("/", indexHandler)                                    //ใช้ HandleFunc เพราะ indexHandler เป็นเหมือน func ธรรมดาที่ต้องแปลงเป็น HandleFunc
	http.Handle("/admin", allowRoleAdmin(http.HandlerFunc(adminHandler))) //ใช้ Handle เพราะ allowRoleAdmin retunr เป็น http.Handler
	// --/admin จะผ่าน  middleware [allowRoleAdmin]
	http.Handle("/stuff", allowRoleStuff(http.HandlerFunc(stuffHandler)))
	//ส่วนของ Handle กับ HandleFunc จะเรียก DefultServeMux อยู่แล้ว
	//เมื่อมีการใช้ DefuleServeMux http.ListenAndServe ตรง handler ก็ใส่ค่า nil ได้เลย เพราะมันจะไปเรียกจาก DefultServeMux ที่เราทำไว้ก่อนหน้านี้
	err := http.ListenAndServe(":8080", nil)
	log.Println(err)

}

/*********************************
*			MIDDLEWARE			 *
---------------------------------*/

type middleware func(http.Handler) http.Handler

//-------- Middleware จัดการ Admin --------
func allowRoleAdmin(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqRole := r.Header.Get("Role")
		if reqRole != "admin" {
			http.Error(w, "Forbidden Admin", http.StatusForbidden)
			return
		}
		h.ServeHTTP(w, r)
	})
}

//-------- Middleware จัดการ Stuff --------
func allowRoleStuff(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqRole := r.Header.Get("Role")
		if reqRole != "stuff" {
			http.Error(w, "Forbidden Stuff", http.StatusForbidden)
			return
		}
		h.ServeHTTP(w, r)
	})
}

/*********************************
*			HANDLER				 *
---------------------------------*/

//-------- INDEX-HANDLER --------
func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Index Page"))
}

//-------- ADMIN-HANDLER --------
func adminHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Admin"))
}

//-------- STUFF-HANDLER --------
func stuffHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Stuff"))
}
