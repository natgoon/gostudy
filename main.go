package main

import "fmt"

// 该因式分解的写法一般用于声明全局变量
var (
	x int
	y bool
)

func main() {
	var a string = "runnodb\n"
	fmt.Printf(a)

	// 变量声明了必须要调用，且int类型不能直接printf打印，需要转换成string类型,println可以直接打印
	var b, c int = 1, 2
	fmt.Printf("b is %d,c is %d\n", b, c)

	var d bool
	fmt.Printf("d is %t\n", d)

	var i int
	var f float64
	var j bool
	var s string
	fmt.Println(i, f, j, s)

	var k = true
	fmt.Println(k)

	// 不带声明格式的变量申请只能在函数体中出现
	h := 1
	fmt.Println(h)

	fmt.Println(x, y)

	test := "asdsadsa"
	fmt.Println(test)

}
