package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	/*
		Q1.1(0)
	*/
	for i := 0; i < 10; i++ {
		fmt.Print(i)
	}
	fmt.Println()

	/*
		Q1.2
	*/
	var j int = 0
HERE1:
	if j < 10 {
		fmt.Print(j)
		j++
		goto HERE1
	}
	fmt.Println()

	/*
		Q1.3
	*/
	var a []int = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	//var a [10]int = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, k := range a {
		fmt.Print(k)
	}
	fmt.Println()

	/*
		Q2(0)
	*/
	for m := 1; m <= 100; m++ {
		if m%15 == 0 {
			fmt.Print("FizzBuzz ")
		} else if m%3 == 0 {
			fmt.Print("Fizz ")
		} else if m%5 == 0 {
			fmt.Print("Buzz ")
		} else {
			fmt.Print(m, " ")
		}
		if m%10 == 0 {
			fmt.Println()
		}
	}

	/*
		Q3.1(1)
	*/
	line, col := 1, 1
	total := 100

	for total > 0 {
		for col <= line {
			fmt.Print("A")
			total--
			col++
			if total == 0 {
				goto HERE2
			}

		}
		col = 1
		fmt.Println()
		line++
	}
HERE2:
	fmt.Println()

	/*
		Q3.2-3.4
	*/
	str := "asSASA abc你好"
	fmt.Println("count(str)= ", utf8.RuneCountInString(str))
	fmt.Println("bytes(str)= ", len(str))

	str_rune := []rune(str)
	//str_rune[4], str_rune[5], str_rune[6] = 'a', 'b', 'c'
	copy(str_rune[4:4+3], []rune("abc"))
	for p := len(str_rune) - 1; p >= 0; p-- {
		fmt.Print(string(str_rune[p]))
	}

}
