/*问题1 求数组[1, 3, 5, 7, 8]所有元素的和
问题2 找出数组中和为指定值的两个元素的下标，比如从数组[1, 3, 5, 7, 8]中找出和为8的两个元素的下标分别为(0,3)和(1,2)。*/
package main

import "fmt"

func qustion1() {
	a :=[...]int{1,3,5,7,8}
	var sum int
	for i:=0;i<len(a);i++{
		sum+=a[i]
	}
	fmt.Printf("%d",sum)
}

func qustion2(){
	a :=[...]int{1,3,5,7,8}
	for i:=0;i<len(a);i++{
		for j:=0;j<len(a);j++{
			if a[i]+a[j]==8{
				if i<j{//这个地方可以不嵌套，与前条件相与
					fmt.Printf("(%d,%d)",i,j)
				}
			}
		}
	}
}

func main() {
	qustion1()
	qustion2()
}
