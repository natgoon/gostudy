package main

import "fmt"

func main() {
	var number1 []int
	number1 = []int{2, 5, 6, 7, 8, 9, 1}

	number2 := []int{3, 4, 5}
	var number3 = []int{6, 7, 8}

	var array = [3]int{7, 8, 9}
	number4 := array[:]
	number5 := number4[:]

	number6 := make([]int, 5, 10)

	fmt.Println(number1)
	fmt.Println(number2)
	fmt.Println(number3)
	fmt.Println(number4)
	fmt.Println(number5)
	fmt.Println(len(number1))
	fmt.Println(cap(number1))
	fmt.Println(len(number6))
	fmt.Println(cap(number6))
	fmt.Println(number6)
	fmt.Println(append(number6, 9))
	fmt.Println(number6)
	number7 := make([]int, 5, cap(number6)*2)
	fmt.Println(copy(number7, number6))
	fmt.Println(number7)
	fmt.Println(cap(number7))

}
