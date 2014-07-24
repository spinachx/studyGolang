package main

import (
	"fmt"
)

func main() {
	mySlice := make([]int, 5, 10)
	fmt.Println("init:")
	fmt.Println("len of mySlice is", len(mySlice))
	fmt.Println("cap of mySlice is", cap(mySlice))

	mySlice = append(mySlice, 1, 2, 3)
	fmt.Println("after append three elem:")
	fmt.Println("len of mySlice is", len(mySlice))
	fmt.Println("cap of mySlice is", cap(mySlice))

	mySlice2 := []int{8, 9, 10}
	mySlice = append(mySlice, mySlice2...)
	fmt.Println("after append mySlice2:")
	fmt.Println("len of mySlice is", len(mySlice))
	fmt.Println("cap of mySlice is", cap(mySlice))

}
