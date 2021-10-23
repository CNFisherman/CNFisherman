package main//数组初始化

import "fmt"
/*初始化数组时可以使用初始化列表来设置数组元素的值*/
func arrayDemo1() {
	var testArray [3]int                        //数组会初始化为int类型的零值
	var numArray = [3]int{1, 2}                 //使用指定的初始值完成初始化
	var cityArray = [3]string{"北京", "上海", "深圳"} //使用指定的初始值完成初始化
	fmt.Println(testArray)                      //[0 0 0]
	fmt.Println(numArray)                       //[1 2 0]
	fmt.Println(cityArray)                      //[北京 上海 深圳]
	fmt.Printf("type of cityArray:%T\n\n", cityArray) 
}
/*一般情况下我们可以让编译器根据初始值的个数自行推断数组的长度*/
func arrayDemo2() {
	var testArray [3]int
	var numArray = [...]int{1, 2}
	var cityArray = [...]string{"北京", "上海", "深圳"}
	fmt.Println(testArray)                          //[0 0 0]
	fmt.Println(numArray)                           //[1 2]
	fmt.Printf("type of numArray:%T\n", numArray)   //type of numArray:[2]int
	fmt.Println(cityArray)                          //[北京 上海 深圳]
	fmt.Printf("type of cityArray:%T\n\n", cityArray) //type of cityArray:[3]string
}
/*使用指定索引值的方式来初始化数组*/
func arrayDemo3() {
	a := [...]int{1: 1, 3: 5}
	fmt.Println(a)                  // [0 1 0 5]
	fmt.Printf("type of a:%T\n", a) //type of a:[4]int
}

func main() {
	arrayDemo1()
	arrayDemo2()
	arrayDemo3()
}
