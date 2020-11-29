package main

import (
	"fmt"
)

func main() {
	// var city map[string]string
	// city = make(map[string]string)
	city1 := make(map[string]string)

	city1["italy"] = "rome"

	capi, ok := city1["italy"]

	fmt.Println(city1["italy"])
	fmt.Println(capi, ok)

	city1["japan"] = "tokyo"
	fmt.Println(city1)
	delete(city1, "italy")
	fmt.Println(city1)
}
