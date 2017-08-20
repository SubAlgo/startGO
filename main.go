package main

import (
	"fmt"
	"time"
)

// GOROUTINES != THREADS
func main() {
	fmt.Println("Start...")
	go talk("Hello")
	go talk("hi")
	fmt.Println("Waiting")
	time.Sleep(5 * time.Second)
	fmt.Println("***END***")

}

func talk(prefix string) {
	for i := 0; i < 10; i++ {
		fmt.Println(prefix, i)
		time.Sleep(time.Second)
	}
}
