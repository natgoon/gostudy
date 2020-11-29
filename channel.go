package main

import (
	"fmt"
	"time"
)

func read(ch chan int) {
	var num int
	for {
		num = <-ch
		fmt.Printf("read: %d\n", num)
	}
}

func write(num int, ch chan int) {
	ch <- num
	fmt.Printf("write: %d\n", num)
}

func main() {
	ch := make(chan int, 10)
	array := []int{2, 5, 6, 7, 8, 9, 1}

	for _, num := range array {
		go write(num, ch)
		go read(ch)
		// fmt.Printf("%d\n", num)
	}
	time.Sleep(time.Duration(1) * time.Second)
	close(ch)
}
