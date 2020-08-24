package main

import "fmt"

func main() {
	numbers := [5]int{1, 0, 4, 0, 2}
	j := 0
	for i := 0; i < 5; i++ {
		if numbers[i] != 0 {
			if numbers[i] != numbers[j] {
				numbers[j] = numbers[i]
				if i != j {
					numbers[i] = 0
				}
			}
			j++
		}
	}

	fmt.Println(numbers)
}
