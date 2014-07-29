package main

import (
	"flag"
	"fmt"
)

func main() {
	username := flag.String("u", "root", "username")
	pwd := flag.String("p", "123456", "password")
	host := flag.String("h", "localhost", "host")
	sid := flag.String("sid", "tmp", "database")

	flag.Parse()

	fmt.Println("username: ", *username)
	fmt.Println("pwd: ", *pwd)
	fmt.Println("host: ", *host)
	fmt.Println("sid: ", *sid)

	//Non-flag
	num := flag.NArg()
	fmt.Printf("%d non-flag args:\n", num)
	for i, v := range flag.Args() {
		fmt.Println("\t", v, " ", flag.Arg(i))
	}

}
