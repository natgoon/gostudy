package main

import (
	"fmt"
)

func fibo(n int) int {
	if n < 2 {
		return n
	} else {
		return fibo(n-2) + fibo(n-1)
	}

}

func fact(n int) int {
	if n < 1 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fact(n-1) * n
	}
}

func main() {
	for i := 0; i < 20; i++ {
		fmt.Printf("%d ", fibo(i))
	}

	fmt.Println(fact(15))

}
