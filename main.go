package main

import (
	"log"
	"net/http"
)

//middleware: 1 page admin and stuff can sign in [struct Method]
func main() {
	//http.DefaultServeMux()
	http.HandleFunc("/", indexHandler)                                        //ใช้ HandleFunc เพราะ indexHandler เป็นเหมือน func ธรรมดาที่ต้องแปลงเป็น HandleFunc
	http.Handle("/admin", allowRole("admin")(http.HandlerFunc(adminHandler))) //ใช้ Handle เพราะ allowRoleAdmin retunr เป็น http.Handler
	// --/admin จะผ่าน  middleware [allowRoleAdmin]
	http.Handle("/stuff", allowRole("stuff")(http.HandlerFunc(stuffHandler)))
	http.Handle("/admin-stuff", allowRoles("admin", "stuff")(http.HandlerFunc(adminStuffHandler)))
	//ส่วนของ Handle กับ HandleFunc จะเรียก DefultServeMux อยู่แล้ว
	//เมื่อมีการใช้ DefuleServeMux http.ListenAndServe ตรง handler ก็ใส่ค่า nil ได้เลย เพราะมันจะไปเรียกจาก DefultServeMux ที่เราทำไว้ก่อนหน้านี้
	err := http.ListenAndServe(":8080", nil)
	log.Println(err)

}

/*********************************
*			MIDDLEWARE			 *
---------------------------------*/

type middleware func(http.Handler) http.Handler

/*
การทำงานของ allowRole คือ เมื่อมีการ request path (func main) จะเรียก func allowRole โดยมีการส่ง string เข้ามา
โดย จะถูกเก็บไว้ใน ตัวแปร role แล้ว Handler เช่น (http.HandlerFunc(adminHandler)) จะถูกโยนใส่ func ตัวใน
ต่อมาตัวแปร reqRole จะถูกสร้างมาเพื่อเก็บ Header สมมติ ที่ชื่ Role
ถ้า path ที่ถูกเรียก header ไม่ตรงกับ role ที่ส่งเข้ามา (role string) ก็จะเข้าไม่ได้
*/
func allowRole(role string) middleware { //รับ role return minddleware
	return func(h http.Handler) http.Handler { //Middleware
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			reqRole := r.Header.Get("Role")
			if reqRole != role {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}
			h.ServeHTTP(w, r)
		})
	}
}

func allowRoles(roles ...string) middleware { //สิ่งที่ roles ...string รับเข้ามาคือ ("admin", "stuff")
	//allow := make(map[string]bool)
	allow := make(map[string]struct{}) //เปลี่ยนจากการเก็บเป็น bool มาเป็น struct ทำให้ประหยัด memory มากกว่า
	for _, role := range roles {
		//allow[role] = true
		allow[role] = struct{}{} //struct{} = type struct | {} หมายถึง struct เปล่า
	}
	//ถ้าสมมติ reqRole มีค่าเป็น admin ดังนั้น allow[reqRole] = allow[admin] = true
	//(true เพราะ กำหนดไว้ตอน map ด้านบน)
	//ดังนั้น !true = false คำสั่งใน if ก็จะไม่ถูกทำงาน

	//***แต่ถ้า reqRole ไม่ใช่ admin หรือ stuff
	//allow[reqRole] = allow[xxx] = false
	//ดังนัั้น !false = true ทำให้คำสั่งใน if ถูกสั่งให้ทำงาน
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			reqRole := r.Header.Get("Role") //กำหนดให้ reqRole เก็บ Header ที่ Get มา
			//if !allow[reqRole] {
			//ให้ ok เช็คค่าว่ามีหรือไม่มีค่า
			//สมมติ allow[reqRole] มีค่าเป็น admin -> ok จะมีค่าเป็น true เมื่อ !trune = false จะออกจาก if
			if _, ok := allow[reqRole]; !ok {
				http.Error(w, "Forbidden /admin-stuff", http.StatusForbidden)
				return
			}
			h.ServeHTTP(w, r)
		})
	}
}

/*-----------------	-----------------------
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
----------------------------------------*/

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

//-------- ADMIN-STUFF-HANDLER --------
func adminStuffHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Admin Stuff"))
}
