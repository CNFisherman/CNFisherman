package main//切片

import "fmt"

func main() {
	// 声明切片类型
	var a []string              //声明一个字符串切片
	var b = []int{}             //声明一个整型切片并初始化
	var c = []bool{false, true} //声明一个布尔切片并初始化
	fmt.Println(a)              //[]
	fmt.Println(b)              //[]
	fmt.Println(c)              //[false true]
	fmt.Println(a == nil)       //true
	fmt.Println(b == nil)       //false
	fmt.Println(c == nil)       //false
	// fmt.Println(c == d)   //切片是引用类型，不支持直接比较，只能和nil比较

	/*切片的底层就是一个数组，所以我们可以基于数组通过切片表达式得到切片。 
	切片表达式中的low和high表示一个索引范围（左包含，右不包含），
	也就是下面代码中从数组a中选出1<=索引值<4的元素组成切片s，
	得到的切片长度=high-low，容量等于得到的切片的底层数组的容量。*/
	e := [5]int{1, 2, 3, 4, 5}
	s := e[1:3]  // s := a[low:high]
	fmt.Printf("s:%v len(s):%v cap(s):%v\n", s, len(s), cap(s))
	s2 := s[3:4]  // 索引的上限是cap(s)而不是len(s)
	fmt.Printf("s2:%v len(s2):%v cap(s2):%v\n", s2, len(s2), cap(s2))
    
	/*对于数组，指向数组的指针，或切片a(注意不能是字符串)支持完整切片表达式：
    a[low : high : max]
    上面的代码会构造与简单切片表达式a[low: high]相同类型、相同长度和元素的切片。
	另外，它会将得到的结果切片的容量设置为max-low。
	在完整切片表达式中只有第一个索引值（low）可以省略；它默认为0。*/
	t := e[1:3:5]
	fmt.Printf("t:%v len(t):%v cap(t):%v\n", t, len(t), cap(t))

	/*下面的代码中演示了拷贝前后两个变量共享底层数组，对一个切片的修改会影响另一个切片的内容，这点需要特别注意。*/
	s1 := make([]int, 3) //[0 0 0]
	s3 := s1             //将s1直接赋值给s2，s1和s2共用一个底层数组
	s3[0] = 100
	fmt.Println(s1) //[100 0 0]
	fmt.Println(s3) //[100 0 0]

}
