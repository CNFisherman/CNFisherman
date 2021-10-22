//switch 使用一览
package main

import "fmt"

func switchDemo1() {
	finger := 3
	switch finger {
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("中指")
	case 4:
		fmt.Println("无名指")
	case 5:
		fmt.Println("小拇指")
	default://Go语言规定每个switch只能有一个default分支。
		fmt.Println("无效的输入！")
	}
}

func switchDemo2() {//一个分支可以有多个值，多个case值中间使用英文逗号分隔。
	switch n := 7; n {
	case 1, 3, 5, 7, 9:
		fmt.Println("奇数")
	case 2, 4, 6, 8:
		fmt.Println("偶数")
	default:
		fmt.Println(n)
	}
}

func switchDemo3() {//分支还可以使用表达式，这时候switch语句后面不需要再跟判断变量。例如：
	age := 30
	switch {
	case age < 25:
		fmt.Println("找个好活")
	case age > 25 && age < 35:
		fmt.Println("多挣挣钱")
	case age > 60:
		fmt.Println("早死早解脱")
	default:
		fmt.Println("hah")
	}
}

func switchDemo4() {//fallthrough语法可以执行满足条件的case的下一个case，是为了兼容C语言中的case设计的。
	s := "a"
	switch {
	case s == "a":
		fmt.Println("a")
		fallthrough
	case s == "b":
		fmt.Println("b")
	case s == "c":
		fmt.Println("c")
	default:
		fmt.Println("...")
	}
}

func main() {
	switchDemo1()
	switchDemo2() 
	switchDemo3() 
	switchDemo4()
}
