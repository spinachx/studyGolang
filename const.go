package main

import (
	"fmt"
)

func main() {
	const (
		a = 1
		b = 2
	)

	const (
		c = iota
		d = iota
	)

	const (
		e = iota
		f
		g
	)

	const (
		h string = "stop"
	)

	fmt.Println(a, b, c, d, e, f, g, h)
}
