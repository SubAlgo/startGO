package main

// GoRutine and Chanel CheckWorkWithTimeLimit
import (
	"fmt"
	"time"
)

func main() {
	//res := doVerryLongWork()
	res, err := doWorkWithLimitTime(2 * time.Second)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}

func doVerryLongWork() int {
	time.Sleep(4 * time.Second)
	return 1
}

func doWorkWithLimitTime(d time.Duration) (int, error) {
	ch := make(chan int)
	go func() {
		ch <- doVerryLongWork()
	}()

	select { //ถ้า chanel ไหนเสร็จก่อนก็ให้ทำ chanel นั้น
	case r := <-ch: //ch ทำงานเสร็จก็ให้เอามาใส่ในตัวแปร r
		return r, nil
	case <-time.After(d): //หลังเวลาที่กำหนด (d) ให้ทำ statement นี้ เช่นค่าที่ส่งมา 5 * time.Second คือ หลังจาก 5วิ
		return 0, fmt.Errorf("timeout")
	}
}
