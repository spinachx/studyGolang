package main

import (
     "fmt"
     "time"
)

func main() {
     timeout := make(chan bool, 1)
     ch := make(chan int, 1)

     go func() {
          time.Sleep(2e9)
          ch <- 1
     }()

     go func() {
          time.Sleep(1e9)
          timeout <- true
     }()

     select {
     case <-ch:
          fmt.Println("get ch")
     case <-timeout:
          fmt.Println("get timeout")
     }
}
