package main

import (
	"fmt"
)

func average(xs []float64) (avg float64) {
	sum := 0.0
	switch len(xs) {
	case 0:
		avg = 0.0
	default:
		for _, v := range xs {
			sum += v
		}
		avg = sum / (float64)(len(xs))
	}
	return
}

func order(a, b int) (x, y int) {
	if a < b {
		return a, b
	}

	return b, a
}

type stack struct {
	i    int
	data [3]int
}

func (s *stack) push(j int) {
	if s.i == 3 {
		fmt.Println("stack is full")
		return
	}
	s.data[s.i] = j
	s.i++

}

func (s *stack) pop() (j int) {
	if s.i == 0 {
		fmt.Println("stack is empty")
		return
	}
	s.i--
	j = s.data[s.i]
	s.data[s.i] = 0
	return j
}

func prt(numbers ...int) {
	for _, v := range numbers {
		fmt.Print(v)
	}
}

func fabi(i int) []int {
	if i <= 2 {
		fmt.Println("input shoub be more than 2")
		return []int{}
	}
	arr := make([]int, i)
	arr[0], arr[1] = 1, 1
	for j := 2; j < i; j++ {
		arr[j] = arr[j-1] + arr[j-2]
	}
	return arr
}

func mymap(f func(int) int, arr []int) []int {
	ret := make([]int, len(arr))
	for k, v := range arr {
		ret[k] = f(v)
	}
	return ret
}

func main() {
	xs := []float64{1.0, 2.3, 3.0}
	fmt.Println(average(xs))

	fmt.Println(order(6, 3))

	var s stack
	s.push(5)
	s.push(6)
	s.push(7)
	s.push(8)
	fmt.Printf("%v", s)
	fmt.Println()
	s.pop()
	fmt.Printf("%v", s)
	fmt.Println()
	s.pop()
	s.pop()
	s.pop()
	fmt.Printf("%v", s)
	fmt.Println()

	prt(1, 2, 3, 4, 5)
	fmt.Println()

	fmt.Printf("%v", fabi(2))
	fmt.Println()

	fmt.Printf("%v", mymap(func(i int) int { return i * 2 }, []int{1, 2, 3, 4}))

}
