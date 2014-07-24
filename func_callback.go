package main

import (
     "fmt"
)


func callback(y int, f func(int)) {
     f(y)
}

func f1(y int) {
     fmt.Println(y + 1)
}

func f2(y int) {
     fmt.Println(y + 2)
}

func main() {
     callback(2, f1)
     callback(3, f2)
}
