package main

import "fmt"

// 遍历字符串
func traversalString() {
	s := "hello沙河"
	for i := 0; i < len(s); i++ { //byte
		fmt.Printf("%v(%c) ", s[i], s[i])
	}
	fmt.Println("\n")
	for _, r := range s { //rune
		fmt.Printf("%v(%c) ", r, r)
	}
	fmt.Println()
}

func main() {
	fmt.Printf("str := \"c:\\Code\\lesson1\\go.exe\"\n")
	traversalString()

}
