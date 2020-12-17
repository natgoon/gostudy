package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	ready := make(chan bool, 1)
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		s := <-sig
		fmt.Println(s)
		fmt.Println("exit..")
		ready <- true
	}()

	r, ok := <-ready
	if ok {
		fmt.Println(r)
	}
}
