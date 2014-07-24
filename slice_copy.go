package main

import (
     "fmt"
)

func main() {
     slice1 := []int{1, 2, 3, 4, 5}
     slice2 := []int{6, 7, 8}

     fmt.Println("init:")
     fmt.Println("slice1:", slice1)
     fmt.Println("slice2:", slice2)

     copy(slice1, slice2)
     fmt.Println("copy slice2 to slice1:", slice1)

     copy(slice2, slice1)
     fmt.Println("copy slice1 to slice2:", slice2)
}
