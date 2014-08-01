package main

import (
	"fmt"
	"reflect"
)

type T struct {
	Name string
	F    func()
}

func main() {
	t := &T{"tt", func() { fmt.Println("ff") }}
	fmt.Println(t.Name)
	t.F()

	for i := 0; i < reflect.ValueOf(t).Elem().NumField(); i++ {
		fmt.Println(i, "--",
			reflect.ValueOf(t).Elem().Type().Field(i), "--",
			reflect.ValueOf(t).Elem().Type().Field(i).Name, "--",
			reflect.ValueOf(t).Elem().Field(i).Type(), "--",
			reflect.ValueOf(t).Elem().Field(i).Interface(), "--")
		if reflect.ValueOf(t).Elem().Field(i).Type().Kind() == reflect.Func {
			// fmt.Println("haha")
			reflect.ValueOf(t).Elem().Field(i).Call([]reflect.Value{})
		}
	}
}
