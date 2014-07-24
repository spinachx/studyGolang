package main

import (
     "errors"
     "fmt"
)

func Add(a, b int) (ret int, err error) {
     if a < 0 || b < 0 {
          err = errors.New("Should be non-negative numbers!")
          return
     }
     return a + b, nil
}

func main() {
     r, _ := Add(1, 2)
     fmt.Println(r)

     r1, e1 := Add(-1, 2)
     fmt.Println(r1, e1)
}
