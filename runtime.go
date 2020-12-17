package main

import (
	"fmt"
	"runtime"
	"time"
)

func init() {
	runtime.GOMAXPROCS(4)
}

func w(ch chan int) {
	// 函数退出时执行defer
	defer fmt.Println("write: write stop!")
	for i := 0; i < 10; i++ {
		ch <- 1
		// 让出当前cpu的时间片，其他进程执行完后本线程继续执行
		runtime.Gosched()
	}
	// 主动让goruntine线程退出
	runtime.Goexit()
}

func r(ch chan int) {
	for {
		sc := <-ch
		fmt.Println(sc)
		// fmt.Println(len(ch))
		time.Sleep(time.Second * time.Duration(1))
	}
}

func main() {
	var ch chan int
	ch = make(chan int, 10)
	go w(ch)
	go r(ch)
	go r(ch)
	fmt.Printf("操作系统：%v\n", runtime.GOOS)
	fmt.Printf("正在运行线程数：%v\n", runtime.NumGoroutine())
	runtime.GC()
	time.Sleep(time.Second * time.Duration(10))
	close(ch)
}
