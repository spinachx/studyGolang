package main
import (
	"fmt"
)

var age int = 30

func sum(x int, y int) (int, int) {
	return x + y, 2
}

type User struct {
	Name string
	Age int
}

func(u *User) PrintName() {
	fmt.Println(u.Name)
}

func main() {
	p := fmt.Println
	fmt.Println("hello world")

	fmt.Println("################变量###############")
	var i int64 = 64
	fmt.Println(i)


	fmt.Println("################ array #############")
	
	arr := []int64{3,5,7}
	fmt.Println("arr=", arr)
	fmt.Println("sencond", arr[1])
	arr = append(arr, 4)
	fmt.Println("arr=", arr)

	for _, v := range arr {
		p( "=", v)
	}

	p("################ map ##################")
	m := map[string]string{
		"name": "dong",
	}

	if _, ok := m["age"]; !ok {
	//if !ok {
		p("不存在对应的值")
	}

	m["name"] = "sui"
	p(m)

	for k, v := range m {
		p(k, "=", v)
	}


	p("################# 控制语句 #################")
	b := false
	if b {
		p(" b is true")
	} else {
		p(" b is false")
	}

	for age:=0; age<10; age++ {
		p("age is ", age)
	}

	p("age is:", age)

	p("################# 函数 #################")
	su := func(x int, y int) {
		p("x+y", x + y)
	}
	su(99,1)

	total, coun := sum(8,12)
	p("total is:", total, "count is", coun)


	p("################# 结构体 #################")

	aUser := &User{
		Name: "dong",
		Age:  23,
	}
	aUser.Name = "sui chang ying"
	aUser.PrintName()

	p("################# interface #############")
	var n interface{}

	n = 3
	p("n is ", n)

	n = "pan "
	p("n is ", n)

	p(n.(string) + " li ")

	n = 55

	switch n.(type) {
		case string:
			p("n is string")
			break
		case int: 
			p("n is int")
			break
	}

	var nn string 
	nn = fmt.Sprintf("%d", n.(int))
	p( "nn is "+nn)

}
