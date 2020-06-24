package main

import (
	"fmt"
	"os"
	"strconv"
	"golearn/ch2/popcount"
)

func main() {
	for _,x := range os.Args[1:] {
		fmt.Println(x);
		x,err := strconv.ParseUint(x, 10, 64)
		if err != nil {
			fmt.Printf("%v", err)
		}
		population := popcount.PopCount(x)
		fmt.Printf("%d\n",population)
	}
}
