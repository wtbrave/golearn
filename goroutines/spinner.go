package main

import (
	"time"
)

func main() {
	go spinner(100 * time.Millisecond)
	const n = 45
	fibN := fib(n)
	println("\r")
	println(fibN)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			time.Sleep(delay)
			print("\r%c ", r)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
