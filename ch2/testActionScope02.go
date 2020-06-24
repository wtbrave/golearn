package main

import "fmt"

func main() {
	x := "hello!"
	for i := 0; i < len(x); i++ {
		fmt.Printf("%c", x)
		fmt.Println()
		x := x[i]
		fmt.Printf("%c",x)
		fmt.Println()
		if x != '!' {
			x := x + 'A' - 'a'
			fmt.Printf("%c", x)
			fmt.Println()
		}
	}
}
