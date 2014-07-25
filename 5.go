package main

import (
	"fmt"
)

type S struct{ i int }

func (p *S) Get() int  { return p.i }
func (p *S) Put(v int) { p.i = v }

//S没有显示声明实现了接口I,但的确实现了I,因为它定义了I的两个方法

type I interface {
	Get() int
	Put(int)
}

type R struct{ i int }

func (p *R) Get() int  { return p.i }
func (p *R) Put(v int) { p.i = v }

//R也实现了接口I

type T struct{ i int }

func (p T) Get() int  { return 1 }
func (p T) Put(v int) {}

func f(i I) {
	fmt.Println(i.Get())
}

func tp(i I) {
	switch i.(type) {
	case *S:
		fmt.Println("*S")
	case *R:
		fmt.Println("*R")
	// case S:
	// 	fmt.Println("S")
	// case R:
	// 	fmt.Println("R")
	case T:
		fmt.Println("T")
	default:
		fmt.Println("Other")
	}
}

func main() {
	var s S
	s.Put(5)
	fmt.Println(s.Get())

	s.Put(6)
	fmt.Println(s.Get())

	//f(s)
	//Error: cannot use s (type S) as type I in argument to f:
	//S does not implement I (Put method has pointer receiver)
	//Get和Put是在s的指针上定义的方法,所有应该要f传递S类型的值(s)的指针来获取s的地址,而不是传递s的值

	f(&s)

	var r R
	r.Put(50)
	fmt.Println(r.Get())

	r.Put(60)
	fmt.Println(r.Get())

	f(&r)

	var t T

	tp(&s)
	tp(&r)
	tp(t)

}
