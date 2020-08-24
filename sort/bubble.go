package main

import "fmt"

func main(){
	numbers := [5]int{1,3,4,6,5}

	for i:=0;i<4;i++{
		for j:=i+1;j<5;j++{
			if numbers[i] > numbers[j] {
				numbers[i] , numbers[j] = numbers[j],numbers[i]
			}
		}
	}
	fmt.Println(numbers)
}
