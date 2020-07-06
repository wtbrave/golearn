package testint

import "fmt"

func main() {
	var a [8]byte
	var b uint8 = 255
	var c  int8 = 25

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

}

func Plus () {
	var a uint8 = 255;
	fmt.Println(a, a+1, a*a);
}

func BitOperate () {
	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2
	fmt.Println(x, y)

	fmt.Printf("%08b\n", x)
	fmt.Printf("%08b\b", y)

	fmt.Printf("%08b\n", x&y)
	fmt.Printf("%08b\n", x|y)
	fmt.Printf("%08b\n", x^y)
	fmt.Printf("%08b\n", x&^y)
	fmt.Println(x, y)


	var  a uint8 = 12
	var  b uint8 = 48
	fmt.Println((a<<2)==b)
}

