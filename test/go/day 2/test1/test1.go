package main

import "fmt"

func main() {
	var a int = 10 //十进制
	fmt.Printf("a=%d\t", a)
	fmt.Printf("a=%b\n", a) //%b代表二进制

	var b int = 077 //八进制
	fmt.Printf("a=%o\n", b)

	var c int = 0xff //十六进制
	fmt.Printf("c=%x\t", c)
	fmt.Printf("c=%d\t", c)
	fmt.Printf("c=%o\t", c)
	fmt.Printf("c=%X\n", c)

}
