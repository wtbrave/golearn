package main

import (
	"time"
)

func main() {
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			time.Sleep(delay)
			print("\r%c", r)
		}
	}
}
