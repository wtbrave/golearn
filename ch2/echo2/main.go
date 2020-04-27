package main

import "fmt"

func main() {
	var a int = 1;
	p := &a
	pp := &p
	fmt.Println(p)
	fmt.Println(pp)
	fmt.Println(*p)
	fmt.Println(*pp)
	fmt.Println(a)
	*p++
	fmt.Println(a)

}
