package main

import "fmt"

type Do interface {
	Tname()
	Tvoice()
	Theight()
}

type Animal struct {
	name   string
	voice  string
	height string
}

func (animal Animal) Tname() {
	fmt.Printf("i am %s\n", animal.name)
}
func (animal Animal) Tvoice() {
	fmt.Printf("%s\n", animal.voice)
}

func (animal Animal) Theight() {
	fmt.Printf("my height is %s\n", animal.height)
}

func main() {

	dog := Animal{"dog", "wangwang", "0.5m"}

	var dogdo Do
	dogdo = dog

	dogdo.Tname()
	dogdo.Tvoice()
	dogdo.Theight()

}
