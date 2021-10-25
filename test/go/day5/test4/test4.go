package main//请使用内置的sort包对数组var a = [...]int{3, 7, 8, 9, 1}进行排序

import(
	"fmt"
	"sort"
)

func main1() {
	var a = [...]int{3, 7, 8, 9, 1}
	b := a[:]
	sort.Ints(b)
	fmt.Println(a)//因为切片是引用类型，a和b指向同一内存地址
}

func main2() {
    // copy()复制切片
    a := []int{3, 7, 8, 9, 1}
	c := make([]int, 5, 5)
	copy(c, a)     //使用copy()函数将切片a中的元素复制到切片c
	sort.Ints(c)
	fmt.Println(a) 
	fmt.Println(c)
}

func main() {
	main1()
	main2()
}
