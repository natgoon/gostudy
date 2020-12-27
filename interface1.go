package main

import (
	"fmt"
)

type DetailInfo struct {
	Email string
	Age   int
}

type Handler interface {
	parse(detail *DetailInfo)
}

type HandleFunc func(*DetailInfo)

func (f HandleFunc) parse(detail *DetailInfo) {
	f(detail)
}

// 针对不同默认参数设置闭包函数， 这里比较关键，闭包函数=函数+运行环境（可以引用外部函数的变量）
func WithEmail(email string) HandleFunc {
	return func(detail *DetailInfo) {
		detail.Email = email
	}
}

func WithAge(age int) HandleFunc {
	return func(detail *DetailInfo) {
		detail.Age = age
	}
}

type Persion struct {
	Name string
	DetailInfo
}

// 这里的接口类Handler并不是必须的，换成闭包函数HandleFunc类型一样可以
func newPersion(name string, infos ...Handler) Persion {
	detail := &DetailInfo{
		Email: "unkown",
		Age:   -1}
	for _, info := range infos {
		// 接口函数调用闭包函数
		info.parse(detail)
	}

	return Persion{Name: name, DetailInfo: DetailInfo{Email: detail.Email, Age: detail.Age}}
}

func main() {
	persion1 := newPersion("小明")
	fmt.Println("persion1:", persion1)
	persion2 := newPersion("小红", WithEmail("xiaohong@qq.com"))
	fmt.Println("persion2:", persion2)
	persion3 := newPersion("张三", WithAge(18))
	fmt.Println("persion3:", persion3)
	persion4 := newPersion("李四", WithEmail("lisi@qq.com"), WithAge(28))
	fmt.Println("persion3:", persion4)
}
