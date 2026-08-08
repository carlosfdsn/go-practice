package main

import "fmt"

func frequency(frequency_list []int) map[int]int {
	frequencies := make(map[int]int)
	
	for _, num := range frequency_list {
		frequencies[num]++
	}
	
	return frequencies
}


func main() {
	numbers := []int{2, 2, 3, 3, 5, 1, 4}
	
	frequencies := frequency(numbers)

	for number, count := range frequencies {
		fmt.Printf("Number %d appears %d time(s)\n", number, count)
	}
}