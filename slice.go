package main

import (
     "fmt"
)

func main() {
     //mySlice1 := make([]int, 5)       //创建一个初始元素个数为5的数组切片,初始值为0
     //mySlice2 := make([]int, 5, 10)   //创建一个初始元素个数为5的数组切片,初始值为0,并预留10个元素的存储空间
     mySlice3 := []int{1, 2, 3, 4, 5} //直接创建并初始化包含5个元素的数组切片

     for i, v := range mySlice3 {
          fmt.Println("mySlice3[", i, "]=", v)
     }
}
