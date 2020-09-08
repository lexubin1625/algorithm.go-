package main

import "fmt"

func main(){
	num := -121
	fmt.Println(isHuiwen(num))
}

func isHuiwen(x int) bool{
	tmp := x
	n := 0
	for x != 0 {
		fmt.Println(x % 10)
		n = n*10 + x % 10
		x = x / 10
	}
	return n == tmp
}