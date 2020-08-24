package main

import "fmt"

func main() {
	numbers := [5]int{1, 1, 3, 5, 5}

	len := len(numbers)
	pre := numbers[0]
	count := 1
	for i:=1;i<len;i++ {
		fmt.Println(i)
		if numbers[i] != pre{
			count++
			pre = numbers[i]
		}
	}
	fmt.Println(count)
}
