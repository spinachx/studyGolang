package main

import (
	"fmt"
	"reflect"
)

type T struct {
	Name string
	F    func(i int, j string)
}

func main() {
	t := &T{"tt", func(i int, j string) { fmt.Println(i); fmt.Println(j) }}
	fmt.Println(t.Name)
	t.F(5, "AA")

	for i := 0; i < reflect.ValueOf(t).Elem().NumField(); i++ {
		fmt.Println(i, "--",
			reflect.ValueOf(t).Elem().Type().Field(i), "--",
			reflect.ValueOf(t).Elem().Type().Field(i).Name, "--",
			reflect.ValueOf(t).Elem().Field(i).Type(), "--",
			reflect.ValueOf(t).Elem().Field(i).Interface(), "--")
		if reflect.ValueOf(t).Elem().Field(i).Type().Kind() == reflect.Func {
			// fmt.Println("haha")
			v := make([]reflect.Value, 2)
			v[0] = reflect.ValueOf(5)
			v[1] = reflect.ValueOf("AA")
			reflect.ValueOf(t).Elem().Field(i).Call(v)
		}
	}
}
