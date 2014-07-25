package main

import (
	"fmt"
)

type S struct {
	i int
}

func (s S) Get() int { return s.i }

type I interface {
	Get() int
}

func g(sth interface{}) int {
	return sth.(I).Get()
}

func main() {
	var s S
	fmt.Println(g(s))

	v := 1
	fmt.Println(g(v))
}
