package main

import (
     "fmt"
)

func myfunc(args ...int) {
     for _, arg := range args {
          fmt.Println(arg)
     }
}

func main() {
     myfunc(1, 2, 3, 4, 5)
}
