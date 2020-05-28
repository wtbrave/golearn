package main

import (
	"fmt"
	"golearn/ch2/popcount"
)
var x uint64 = 1233334

func main() {
	population := popcount.PopCount(x)
	fmt.Printf("%d\n",population)
}
