package main//遍历数组

import "fmt"

func main() {
	var a = [...]string{"北京", "上海", "深圳"}
	b:= [3][2]string{
		{"北京", "上海"},
		{"广州", "深圳"},
		{"成都", "重庆"},
	}
	// 方法1：for循环遍历
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}//遍历一维数组
	fmt.Printf("\n")
	for i := 0; i < 3; i++{
		for j := 0; j < 2; j++{
			fmt.Println(b[i][j])
		}
	}

	// 方法2：for range遍历
	for index, value := range a {
		fmt.Println(index, value)
	}//遍历一维数组
	for _, v1 := range b {
		for _, v2 := range v1 {
			fmt.Printf("%s\t", v2)
		}
		fmt.Println()
	}//遍历二维数组
}
