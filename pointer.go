package main

import (
     "fmt"
)

type Integer int

func (a Integer) Less(b Integer) bool {
     return a < b
}

func (a *Integer) Add(b Integer) {
     *a += b
}

func (a Integer) Add2(b Integer) Integer {
     return a + b
}

func main() {
     var i Integer = 1
     fmt.Println(i.Less(2))
     i.Add(3)
     fmt.Println(i)
     fmt.Println(i.Add2(5))
     fmt.Println(i)
}
