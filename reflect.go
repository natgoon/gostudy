package main

import (
	"fmt"
	"reflect"
	"sync"
)

func main() {
	var once sync.Once

	t := reflect.TypeOf(once)

	fmt.Println(t.Name(), t.Kind(), t)

	v := reflect.ValueOf(once)
	fmt.Println(v, v.Kind(), v.Interface(), v.IsValid())
}
