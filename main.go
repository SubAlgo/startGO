package main

import (
	"fmt"
)

func main() {
	fmt.Println("Start...")
	doSafeWork()
	fmt.Println("Done")
}

func doFailWork() {
	panic("fail")
}

func doSafeWork() {
	//defer func() {...} ()
	//defer เหมือนเป็นการบังคับว่าก่อนจบ function ให้มาทำงานในสั่ง defer ก่อนที่จะจบ
	//อย่างในกรณีนี้ ความจริงจะออกจาก function ตั้งแต่บรรทัด doFailWork()
	//แต่พอ doFailWork() ทำงานจบจะถูกบังคับให้ทำใน defer ก่อนจบ
	//เพื่อให้รู้ผลลัพธ์ที่แตกต่างให้ลอง commend doFailWork() ออก
	defer func() {
		r := recover()
		fmt.Println("r inside doSafeWork ", r)
	}()
	doFailWork()
	fmt.Println("work success")

}
