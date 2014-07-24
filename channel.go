package main

import (
     "fmt"
)

func Count(ch chan int) {
     fmt.Println("Counting...")
     fmt.Println(ch)
     ch <- 1
}

func main() {
     chs := make([]chan int, 5)
     for i := 0; i < 5; i++ {
          chs[i] = make(chan int)
          go Count(chs[i])
     }

     for i, ch := range chs {
          fmt.Println("Read...", i)
          <-ch
     }
}
