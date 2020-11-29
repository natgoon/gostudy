package main

import "fmt"

func main() {
	str := "hello everyone!"
	arr := [5]int{1, 2, 3, 4, 5}
	sli := []int{7, 8, 9}
	cha := make(chan int, 100)
	mp := make(map[string]string)
	mp["italy"] = "roma"
	mp["japan"] = "tokyo"

	fmt.Println(str, arr, sli, cha)

	for str := range str {
		cha <- str
		fmt.Println("cha is", <-cha)
	}

	var sum int
	sum = 0
	for _, num := range arr {
		sum = sum + num
	}
	fmt.Println("sum is", sum)

	for country, city := range mp {
		fmt.Println(country, "capital is", city)
	}

}
