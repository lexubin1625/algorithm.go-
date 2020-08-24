package main

import "fmt"

func main(){
	a := -1432
	a = reverse(a)
	fmt.Println(a)
}

func reverse(x int) int{
	var nums,newnums int
	for  x != 0{//直到x等于0，跳出循环
		a:=x%10

		newnums=nums*10+a
		fmt.Println(newnums,nums,a)
		nums=newnums
		x=x/10

		//题目要求其数值范围是 [−2^31,  2^31 − 1]。如果反转后的整数溢出，则返回 0。
		MaxInt32 := 1<<31 - 1
		MinInt32 := -1 << 31
		if nums > MaxInt32 || nums < MinInt32 {
			return 0
		}
	}
	return nums
}
