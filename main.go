package main

import (
	"log"
	"os"
)

func main() {
	f, err := os.Create("Hello.txt")
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close() //สั่งปิด f เมื่อ run main จบ
	f.WriteString("Hello")
}
