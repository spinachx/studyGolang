package main

import (
     "fmt"
)

func f() (ret int) {
     defer func() {
          ret++
     }()
     return 5
}

func main() {
     fmt.Println(f())
}

/*
    output:
    6
*/
