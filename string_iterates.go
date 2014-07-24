package main

import (
     "fmt"
)

func f1() {
     //以字节数组遍历
     fmt.Println("This is f1 output.")
     str := "Hello, 世界"
     n := len(str)
     for i := 0; i < n; i++ {
          ch := str[i]
          fmt.Println(i, ch)
     }
}

func f2() {
     //以Unicode字符遍历
     fmt.Println("This is f2 output.")
     str := "Hello, 世界"
     for i, ch := range str {
          fmt.Println(i, ch)
     }
}

func main() {
     f1()
     f2()
}
